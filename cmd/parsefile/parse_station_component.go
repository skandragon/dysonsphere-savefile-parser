package main

type StationComponent struct {
	ID                    int32           `json:"id"`
	GID                   int32           `json:"gid"`
	PlanetID              int32           `json:"planet_id"`
	IsStellar             bool            `json:"is_stellar"`
	Name                  string          `json:"name"`
	Energy                int64           `json:"energy"`
	WarperCount           int32           `json:"warper_count"`
	WarperMaxCount        int32           `json:"warper_max_count"`
	IdleDroneCount        int32           `json:"idle_drone_count"`
	WorkDroneCount        int32           `json:"work_drone_count"`
	IdleShipCount         int32           `json:"idle_ship_count"`
	WorkShipCount         int32           `json:"work_ship_count"`
	IsCollector           bool            `json:"is_collector"`
	CollectionIDs         []int32         `json:"collection_i_ds"`
	CollectionPerTick     []float32       `json:"collection_per_tick"`
	CurrentCollections    []float32       `json:"current_collections"`
	CollectSpeed          int32           `json:"collect_speed"`
	StationStores         []*StationStore `json:"station_stores"`
	TripRangeDrones       float64         `json:"trip_range_drones"`
	TripRangeShips        float64         `json:"trip_range_ships"`
	IncludeOrbitCollector bool            `json:"include_orbit_collector"`
	WarpEnableDist        float64         `json:"warp_enable_dist"`
	WarperNecessary       bool            `json:"warper_necessary"`
	DeliveryShips         int32           `json:"delivery_ships"`
	DeliveryDrones        int32           `json:"delivery_drones"`
}

func parseStationComponent(b *Buffer) *StationComponent {
	checkVers(b, 2, "StationComponent")
	station := &StationComponent{}

	station.ID = b.GetInt32le()        // id = r.ReadInt32();
	station.GID = b.GetInt32le()       // gid = r.ReadInt32();
	b.GetInt32le()                     // entityId = r.ReadInt32();
	station.PlanetID = b.GetInt32le()  // planetId = r.ReadInt32();
	b.GetInt32le()                     // pcId = r.ReadInt32();
	b.GetInt32le()                     // gene = r.ReadInt32();
	b.GetFloat32()                     // droneDock.x = r.ReadSingle();
	b.GetFloat32()                     // droneDock.y = r.ReadSingle();
	b.GetFloat32()                     // droneDock.z = r.ReadSingle();
	b.GetFloat32()                     // shipDockPos.x = r.ReadSingle();
	b.GetFloat32()                     // shipDockPos.y = r.ReadSingle();
	b.GetFloat32()                     // shipDockPos.z = r.ReadSingle();
	b.GetFloat32()                     // shipDockRot.x = r.ReadSingle();
	b.GetFloat32()                     // shipDockRot.y = r.ReadSingle();
	b.GetFloat32()                     // shipDockRot.z = r.ReadSingle();
	b.GetFloat32()                     // shipDockRot.w = r.ReadSingle();
	station.IsStellar = b.GetBoolean() // isStellar = r.ReadBoolean();
	//fmt.Printf("Station (%d,%d): isStellar=%v\n", id, gid, isStellar)
	hasName := b.GetInt32le()
	if hasName > 0 {
		station.Name = b.GetString() // name = r.ReadString();
		//fmt.Printf("Station (%d,%d) name: %s\n", id, gid, name)
	}
	station.Energy = b.GetInt64le()         // energy = r.ReadInt64();
	b.GetInt64le()                          // energyPerTick = r.ReadInt64();
	b.GetInt64le()                          // energyMax = r.ReadInt64();
	station.WarperCount = b.GetInt32le()    // warperCount = r.ReadInt32();
	station.WarperMaxCount = b.GetInt32le() // warperMaxCount = r.ReadInt32();
	station.IdleDroneCount = b.GetInt32le() // idleDroneCount = r.ReadInt32();
	station.WorkDroneCount = b.GetInt32le()
	//fmt.Printf("Drones: %d idle, %d working\n", idleDroneCount, workDroneCount)
	b.GetInt32le() // workDroneDatas.Length
	for i := int32(0); i < station.WorkDroneCount; i++ {
		parseDroneData(b)
	}
	for i := int32(0); i < station.WorkDroneCount; i++ {
		parseLocalLogisticOrder(b)
	}
	station.IdleShipCount = b.GetInt32le() // idleShipCount = r.ReadInt32();
	station.WorkShipCount = b.GetInt32le() // workShipCount = r.ReadInt32();
	//fmt.Printf("Ships: %d idle, %d working\n", idleShipCount, workShipCount)
	b.GetInt64le() // idleShipIndices = r.ReadUInt64();
	b.GetInt64le() // workShipIndices = r.ReadUInt64();
	b.GetInt32le() // workShipDatas.Length
	for i := int32(0); i < station.WorkShipCount; i++ {
		parseShipData(b)
	}
	for i := int32(0); i < station.WorkShipCount; i++ {
		parseRemoteLogisticOrder(b)
	}
	stationStoreCount := b.GetInt32le()
	station.StationStores = make([]*StationStore, 0)
	for i := int32(0); i < stationStoreCount; i++ {
		station.StationStores = append(station.StationStores, parseStationStore(b))
	}
	slotCount := b.GetInt32le()
	for i := int32(0); i < slotCount; i++ {
		b.GetInt32le() // slots[n].dir = (IODir)r.ReadInt32();
		b.GetInt32le() // slots[n].beltId = r.ReadInt32();
		b.GetInt32le() // slots[n].storageIdx = r.ReadInt32();
		b.GetInt32le() // slots[n].counter = r.ReadInt32();
	}
	b.GetInt32le() // localPairProcess = r.ReadInt32();
	b.GetInt32le() // remotePairProcess = r.ReadInt32();
	b.GetInt32le() // nextShipIndex = r.ReadInt32();
	station.IsCollector = b.GetBoolean()

	collectionIDsLength := b.GetInt32le()
	station.CollectionIDs = make([]int32, int(collectionIDsLength))
	for i := int32(0); i < collectionIDsLength; i++ {
		station.CollectionIDs[i] = b.GetInt32le()
	}

	collectionPerTickLength := b.GetInt32le()
	station.CollectionPerTick = make([]float32, int(collectionPerTickLength))
	for i := int32(0); i < collectionPerTickLength; i++ {
		station.CollectionPerTick[i] = b.GetFloat32()
	}

	currentCollectionsLength := b.GetInt32le()
	station.CurrentCollections = make([]float32, int(currentCollectionsLength))
	for i := int32(0); i < currentCollectionsLength; i++ {
		station.CurrentCollections[i] = b.GetFloat32()
	}
	station.CollectSpeed = b.GetInt32le()
	station.TripRangeDrones = b.GetFloat64()
	station.TripRangeShips = b.GetFloat64()
	station.IncludeOrbitCollector = b.GetBoolean()
	station.WarpEnableDist = b.GetFloat64()
	station.WarperNecessary = b.GetBoolean()
	station.DeliveryDrones = b.GetInt32le()
	station.DeliveryShips = b.GetInt32le()

	return station
}
