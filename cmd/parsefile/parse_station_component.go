package main

// StationCollection describes an item this station is sucking down.
type StationCollection struct {
	ItemID  int32   `json:"item_id"`
	PerTick float32 `json:"per_tick"`
	Current float32 `json:"current"`
}

// StationComponent describes a single planet/interplanetary/collector config and status
type StationComponent struct {
	ID                    int32                `json:"id"`
	GID                   int32                `json:"gid"`
	EntityID              int32                `json:"entity_id"`
	PlanetID              int32                `json:"planet_id"`
	IsStellar             bool                 `json:"is_stellar"`
	Name                  string               `json:"name"`
	Energy                int64                `json:"energy"`
	EnergyPerTick         int64                `json:"energy_per_tick"`
	EnergyMax             int64                `json:"energy_max"`
	WarperCount           int32                `json:"warper_count"`
	WarperMaxCount        int32                `json:"warper_max_count"`
	IdleDroneCount        int32                `json:"idle_drone_count"`
	WorkDroneCount        int32                `json:"work_drone_count"`
	IdleShipCount         int32                `json:"idle_ship_count"`
	WorkShipCount         int32                `json:"work_ship_count"`
	IsCollector           bool                 `json:"is_collector"`
	Collections           []*StationCollection `json:"collections,omitempty"`
	CollectSpeed          int32                `json:"collect_speed"`
	StationStores         []*StationStore      `json:"station_stores"`
	TripRangeDrones       float64              `json:"trip_range_drones"`
	TripRangeShips        float64              `json:"trip_range_ships"`
	IncludeOrbitCollector bool                 `json:"include_orbit_collector"`
	WarpEnableDist        float64              `json:"warp_enable_dist"`
	WarperNecessary       bool                 `json:"warper_necessary"`
	DeliveryShips         int32                `json:"delivery_ships"`
	DeliveryDrones        int32                `json:"delivery_drones"`
}

func parseStationComponent(b *Buffer) *StationComponent {
	checkVers(b, 2, "StationComponent")
	station := &StationComponent{}

	station.ID = b.GetInt32le()
	station.GID = b.GetInt32le()
	station.EntityID = b.GetInt32le()
	station.PlanetID = b.GetInt32le()
	b.GetInt32le() // pcId = r.ReadInt32();
	b.GetInt32le() // gene = r.ReadInt32();
	b.GetFloat32() // droneDock.x = r.ReadSingle();
	b.GetFloat32() // droneDock.y = r.ReadSingle();
	b.GetFloat32() // droneDock.z = r.ReadSingle();
	b.GetFloat32() // shipDockPos.x = r.ReadSingle();
	b.GetFloat32() // shipDockPos.y = r.ReadSingle();
	b.GetFloat32() // shipDockPos.z = r.ReadSingle();
	b.GetFloat32() // shipDockRot.x = r.ReadSingle();
	b.GetFloat32() // shipDockRot.y = r.ReadSingle();
	b.GetFloat32() // shipDockRot.z = r.ReadSingle();
	b.GetFloat32() // shipDockRot.w = r.ReadSingle();
	station.IsStellar = b.GetBoolean()
	hasName := b.GetInt32le()
	if hasName > 0 {
		station.Name = b.GetString()
	}
	station.Energy = b.GetInt64le()
	station.EnergyPerTick = b.GetInt64le()
	station.EnergyMax = b.GetInt64le()
	station.WarperCount = b.GetInt32le()
	station.WarperMaxCount = b.GetInt32le()
	station.IdleDroneCount = b.GetInt32le()
	station.WorkDroneCount = b.GetInt32le()
	b.GetInt32le() // workDroneDatas.Length
	for i := int32(0); i < station.WorkDroneCount; i++ {
		parseDroneData(b)
	}
	for i := int32(0); i < station.WorkDroneCount; i++ {
		parseLocalLogisticOrder(b)
	}
	station.IdleShipCount = b.GetInt32le()
	station.WorkShipCount = b.GetInt32le()
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
		store := parseStationStore(b, i)
		if store.ItemID > 0 {
			station.StationStores = append(station.StationStores, store)
		}
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

	station.Collections = parseStationCollections(b)

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

func parseStationCollections(b *Buffer) []*StationCollection {
	collectionIDsLength := b.GetInt32le()
	ids := make([]int32, int(collectionIDsLength))
	for i := int32(0); i < collectionIDsLength; i++ {
		ids[i] = b.GetInt32le()
	}

	collectionPerTickLength := b.GetInt32le()
	perTick := make([]float32, int(collectionPerTickLength))
	for i := int32(0); i < collectionPerTickLength; i++ {
		perTick[i] = b.GetFloat32()
	}

	currentCollectionsLength := b.GetInt32le()
	current := make([]float32, int(currentCollectionsLength))
	for i := int32(0); i < currentCollectionsLength; i++ {
		current[i] = b.GetFloat32()
	}

	// Not sure what to do if these are not the same length.  For now, ignore the undefined ones...
	count := collectionIDsLength
	if collectionPerTickLength < count {
		count = collectionPerTickLength
	}
	if currentCollectionsLength < count {
		count = currentCollectionsLength
	}

	if count == 0 {
		return nil
	}

	ret := make([]*StationCollection, count)
	for i := int32(0); i < count; i++ {
		ret[i] = &StationCollection{
			ItemID:  ids[i],
			PerTick: perTick[i],
			Current: current[i],
		}
	}
	return ret
}
