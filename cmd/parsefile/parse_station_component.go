package main

func parseStationComponent(b *Buffer) {
	checkVers(b, 2, "StationComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // gid = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // planetId = r.ReadInt32();
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
	/*isStellar :=*/ b.GetBoolean() // isStellar = r.ReadBoolean();
	//fmt.Printf("Station (%d,%d): isStellar=%v\n", id, gid, isStellar)
	hasName := b.GetInt32le()
	if hasName > 0 {
		/*name := */ b.GetString() // name = r.ReadString();
		//fmt.Printf("Station (%d,%d) name: %s\n", id, gid, name)
	}
	b.GetInt64le() // energy = r.ReadInt64();
	b.GetInt64le() // energyPerTick = r.ReadInt64();
	b.GetInt64le() // energyMax = r.ReadInt64();
	b.GetInt32le() // warperCount = r.ReadInt32();
	b.GetInt32le() // warperMaxCount = r.ReadInt32();
	b.GetInt32le() // idleDroneCount = r.ReadInt32();
	workDroneCount := b.GetInt32le()
	//fmt.Printf("Drones: %d idle, %d working\n", idleDroneCount, workDroneCount)
	b.GetInt32le() // workDroneDatas.Length
	for i := int32(0); i < workDroneCount; i++ {
		parseDroneData(b)
	}
	for i := int32(0); i < workDroneCount; i++ {
		parseLocalLogisticOrder(b)
	}
	b.GetInt32le()                  // idleShipCount = r.ReadInt32();
	workShipCount := b.GetInt32le() // workShipCount = r.ReadInt32();
	//fmt.Printf("Ships: %d idle, %d working\n", idleShipCount, workShipCount)
	b.GetInt64le() // idleShipIndices = r.ReadUInt64();
	b.GetInt64le() // workShipIndices = r.ReadUInt64();
	b.GetInt32le() // workShipDatas.Length
	for i := int32(0); i < workShipCount; i++ {
		parseShipData(b)
	}
	for i := int32(0); i < workShipCount; i++ {
		parseRemoteLogisticOrder(b)
	}
	stationStoreCount := b.GetInt32le()
	for i := int32(0); i < stationStoreCount; i++ {
		parseStationStore(b)
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
	b.GetBoolean() // isCollector = r.ReadBoolean();
	collectionIDsLength := b.GetInt32le()
	for i := int32(0); i < collectionIDsLength; i++ {
		b.GetInt32le() // collectionIds[num7] = r.ReadInt32();
	}
	collectionPerTickLength := b.GetInt32le()
	for i := int32(0); i < collectionPerTickLength; i++ {
		b.GetFloat32() // collectionPerTick[num8] = r.ReadSingle();
	}
	currentCollectionsLength := b.GetInt32le()
	for i := int32(0); i < currentCollectionsLength; i++ {
		b.GetFloat32() // currentCollections[num9] = r.ReadSingle();
	}
	b.GetInt32le() // collectSpeed = r.ReadInt32();
	b.GetFloat64() // tripRangeDrones = r.ReadDouble();
	b.GetFloat64() // tripRangeShips = r.ReadDouble();
	b.GetBoolean() // includeOrbitCollector = r.ReadBoolean();
	b.GetFloat64() // warpEnableDist = r.ReadDouble();
	b.GetBoolean() // warperNecessary = r.ReadBoolean();
	b.GetInt32le() // deliveryDrones = r.ReadInt32();
	b.GetInt32le() // deliveryShips = r.ReadInt32();
}
