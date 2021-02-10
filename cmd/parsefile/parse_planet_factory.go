package main

import (
	"fmt"
	"os"
)

func parsePlanetFactory(b *Buffer, i int) {
	checkVers(b, 1, "PlanetFactory")

	planetID := b.GetInt32le()
	fmt.Printf("Planet Factory %d, planetID: %d\n", i, planetID)

	parsePlanetDataRuntime(b)

	entityCapacity := b.GetInt32le() // entityCapacity
	entityCursor := b.GetInt32le()
	entityRecycleCursor := b.GetInt32le()
	fmt.Printf(" entityCapacity %d, cursor %d, recycleCursor %d\n",
		entityCapacity, entityCursor, entityRecycleCursor)
	for i := 1; int32(i) < entityCursor; i++ {
		parseEntityData(b)
	}
	for i := 1; int32(i) < entityCursor; i++ {
		b.GetFloat32()  // entityAnimPool.time
		b.GetFloat32()  // entityAnimPool.prepare_length
		b.GetFloat32()  // entityAnimPool.working_length
		b.GetUInt32le() // entityAnimPool.state
		b.GetFloat32()  // entityAnimPool.power
	}
	for i := 1; int32(i) < entityCursor; i++ {
		b.GetUInt32le() // entitySignPool[k].signType)
		b.GetUInt32le() // entitySignPool[k].iconType)
		b.GetUInt32le() // entitySignPool[k].iconId0)
		b.GetUInt32le() // entitySignPool[k].iconId1)
		b.GetUInt32le() // entitySignPool[k].iconId2)
		b.GetUInt32le() // entitySignPool[k].iconId3)
		b.GetFloat32()  // entitySignPool[k].count0)
		b.GetFloat32()  // entitySignPool[k].count1)
		b.GetFloat32()  // entitySignPool[k].count2)
		b.GetFloat32()  // entitySignPool[k].count3)
		b.GetFloat32()  // entitySignPool[k].x)
		b.GetFloat32()  // entitySignPool[k].y)
		b.GetFloat32()  // entitySignPool[k].z)
		b.GetFloat32()  // entitySignPool[k].w)
	}
	connPoolCount := entityCursor * 16
	for i := 16; int32(i) < connPoolCount; i++ {
		b.GetInt32le() // entityConnPool[l])
	}
	for i := 0; int32(i) < entityRecycleCursor; i++ {
		b.GetInt32le() // entityRecycle[m])
	}

	prebuildCapacity := b.GetInt32le() // prebuildCapacity
	prebuildCursor := b.GetInt32le()
	prebuildRecycleCursor := b.GetInt32le()
	fmt.Printf(" prebuildCapacity %d, cursor %d, recycleCursor %d\n",
		prebuildCapacity, prebuildCursor, prebuildRecycleCursor)
	for i := 1; int32(i) < prebuildCursor; i++ {
		parsePrebuildData(b)
	}
	prebuildConnPoolCount := prebuildCursor * 16
	for i := 16; int32(i) < prebuildConnPoolCount; i++ {
		b.GetInt32le() // prebuildConnPool[l])
	}
	for i := 0; int32(i) < prebuildRecycleCursor; i++ {
		b.GetInt32le() // prebuildRecycle[m])
	}

	vegeCapacity := b.GetInt32le() // vegeCapacity);
	vegeCursor := b.GetInt32le()
	vegeRecycleCursor := b.GetInt32le()
	fmt.Printf(" vegeCapacity %d, cursor %d, recycleCursor %d\n",
		vegeCapacity, vegeCursor, vegeRecycleCursor)
	for i := 1; int32(i) < vegeCursor; i++ {
		parseVegeData(b)
	}
	for i := 0; int32(i) < vegeRecycleCursor; i++ {
		b.GetInt32le() // vegeRecycle[num6]);
	}

	veinCapacity := b.GetInt32le() // veinCapacity
	veinCursor := b.GetInt32le()   // veinCursor
	veinRecycleCursor := b.GetInt32le()
	fmt.Printf(" veinCapacity %d, cursor %d, recycleCursor %d\n",
		veinCapacity, veinCursor, veinRecycleCursor)
	for i := 1; int32(i) < veinCursor; i++ {
		parseVeinData(b)
	}
	for i := 0; int32(i) < veinRecycleCursor; i++ {
		b.GetInt32le() // veinRecycle
	}
	for i := 1; int32(i) < veinCursor; i++ {
		b.GetFloat32()  // veinAnimPool.time
		b.GetFloat32()  // veinAnimPool.prepare_length
		b.GetFloat32()  // veinAnimPool.working_length
		b.GetUInt32le() // veinAnimPool.state
		b.GetFloat32()  // veinAnimPool.power
	}

	parseCargoContainer(b)
	parseCargoTraffic(b)
	parseFactoryStorage(b)
	parsePowerSystem(b)
	os.Exit(1)
	parseFactorySystem(b)
	parsePlanetTransport(b)
	parseMonsterSystem(b)
	parsePlatformSystem(b)
}

func parsePlanetDataRuntime(b *Buffer) {
	fmt.Println("Start: PlanetData.Runtime")
	count := b.GetInt32le()
	b.GetBytes(int(count)) // modData

	veinCount := b.GetInt32le()
	for i := 0; int32(i) < veinCount; i++ {
		veinAmount := b.GetInt64le()
		if veinAmount > 0 {
			fmt.Printf("  vein type %d, amount: %d\n", i, veinAmount)
		}
	}

	veinGroupCount := b.GetInt32le()
	for i := 0; int32(i) < veinGroupCount; i++ {
		ty := b.GetInt32le() // vein type
		b.GetFloat32()       // pos.x
		b.GetFloat32()       // pos.y
		b.GetFloat32()       // pos.z
		c := b.GetInt32le()  // count
		amount := b.GetInt64le()
		if amount > 0 {
			fmt.Printf("  vein type %d (%s), amount %d, count %d\n", ty, VeinType(ty), amount, c)
		}
	}
	fmt.Println("End: PlanetData.Runtime")
}

func parseEntityData(b *Buffer) {
	checkVersByte(b, 0, "EntityData")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt16le() // protoId = r.ReadInt16();
	b.GetInt16le() // modelIndex = r.ReadInt16();
	b.GetFloat32() // pos.x = r.ReadSingle();
	b.GetFloat32() // pos.y = r.ReadSingle();
	b.GetFloat32() // pos.z = r.ReadSingle();
	b.GetFloat32() // rot.x = r.ReadSingle();
	b.GetFloat32() // rot.y = r.ReadSingle();
	b.GetFloat32() // rot.z = r.ReadSingle();
	b.GetFloat32() // rot.w = r.ReadSingle();
	b.GetInt32le() // beltId = r.ReadInt32();
	b.GetInt32le() // splitterId = r.ReadInt32();
	b.GetInt32le() // storageId = r.ReadInt32();
	b.GetInt32le() // tankId = r.ReadInt32();
	b.GetInt32le() // minerId = r.ReadInt32();
	b.GetInt32le() // inserterId = r.ReadInt32();
	b.GetInt32le() // assemblerId = r.ReadInt32();
	b.GetInt32le() // fractionateId = r.ReadInt32();
	b.GetInt32le() // ejectorId = r.ReadInt32();
	b.GetInt32le() // siloId = r.ReadInt32();
	b.GetInt32le() // labId = r.ReadInt32();
	b.GetInt32le() // stationId = r.ReadInt32();
	b.GetInt32le() // powerNodeId = r.ReadInt32();
	b.GetInt32le() // powerGenId = r.ReadInt32();
	b.GetInt32le() // powerConId = r.ReadInt32();
	b.GetInt32le() // powerAccId = r.ReadInt32();
	b.GetInt32le() // powerExcId = r.ReadInt32();
	b.GetInt32le() // monsterId = r.ReadInt32();
}

func parsePrebuildData(b *Buffer) {
	checkVersByte(b, 0, "PrebuildData")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt16le() // protoId = r.ReadInt16();
	b.GetInt16le() // modelIndex = r.ReadInt16();
	b.GetFloat32() // pos.x = r.ReadSingle();
	b.GetFloat32() // pos.y = r.ReadSingle();
	b.GetFloat32() // pos.z = r.ReadSingle();
	b.GetFloat32() // rot.x = r.ReadSingle();
	b.GetFloat32() // rot.y = r.ReadSingle();
	b.GetFloat32() // rot.z = r.ReadSingle();
	b.GetFloat32() // rot.w = r.ReadSingle();
	b.GetFloat32() // pos2.x = r.ReadSingle();
	b.GetFloat32() // pos2.y = r.ReadSingle();
	b.GetFloat32() // pos2.z = r.ReadSingle();
	b.GetFloat32() // rot2.x = r.ReadSingle();
	b.GetFloat32() // rot2.y = r.ReadSingle();
	b.GetFloat32() // rot2.z = r.ReadSingle();
	b.GetFloat32() // rot2.w = r.ReadSingle();
	b.GetInt32le() // upEntity = r.ReadInt32();
	b.GetInt16le() // pickOffset = r.ReadInt16();
	b.GetInt16le() // insertOffset = r.ReadInt16();
	b.GetInt32le() // recipeId = r.ReadInt32();
	b.GetInt32le() // filterId = r.ReadInt32();
	b.GetInt32le() // refCount = r.ReadInt32();
}

func parseVegeData(b *Buffer) {
	checkVersByte(b, 0, "VegeData")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt16le() // protoId = r.ReadInt16()
	b.GetInt16le() // modelIndex = r.ReadInt16()
	b.GetInt16le() // hp = r.ReadInt16()
	b.GetFloat32() // pos.x = r.ReadSingle()
	b.GetFloat32() // pos.y = r.ReadSingle()
	b.GetFloat32() // pos.z = r.ReadSingle()
	b.GetFloat32() // rot.x = r.ReadSingle()
	b.GetFloat32() // rot.y = r.ReadSingle()
	b.GetFloat32() // rot.z = r.ReadSingle()
	b.GetFloat32() // rot.w = r.ReadSingle()
	b.GetFloat32() // scl.x = r.ReadSingle()
	b.GetFloat32() // scl.y = r.ReadSingle()
	b.GetFloat32() // scl.z = r.ReadSingle()
}

func parseVeinData(b *Buffer) {
	checkVersByte(b, 0, "VeinData")

	id := b.GetInt32le()         // id = r.ReadInt32();
	ty := b.GetInt16le()         // type = (EVeinType)r.ReadInt16();
	b.GetInt16le()               // modelIndex = r.ReadInt16();
	b.GetInt16le()               // groupIndex = r.ReadInt16();
	amount := b.GetInt32le()     // amount = r.ReadInt32();
	productID := b.GetInt32le()  // productId = r.ReadInt32();
	b.GetFloat32()               // pos.x = r.ReadSingle();
	b.GetFloat32()               // pos.y = r.ReadSingle();
	b.GetFloat32()               // pos.z = r.ReadSingle();
	minerCount := b.GetInt32le() // minerCount = r.ReadInt32();
	b.GetInt32le()               // minerId0 = r.ReadInt32();
	b.GetInt32le()               // minerId1 = r.ReadInt32();
	b.GetInt32le()               // minerId2 = r.ReadInt32();
	b.GetInt32le()               // minerId3 = r.ReadInt32();
	if id != 0 {
		fmt.Printf("  id %d, type %d (%s), amount %d, productID %d, minercount %d\n",
			id, ty, VeinType(ty), amount, productID, minerCount)
	}
}

func parseCargoContainer(b *Buffer) {
	checkVers(b, 0, "CargoContainer")

	poolCapacity := b.GetInt32le()
	cursor := b.GetInt32le()
	b.GetInt32le() // recycleBegin
	b.GetInt32le() // recycleEnd
	for i := 0; int32(i) < cursor; i++ {
		b.GetInt32le() // cargoPool.item
		b.GetFloat32() // cargoPool.position.x
		b.GetFloat32() // cargoPool.position.y
		b.GetFloat32() // cargoPool.position.z
		b.GetFloat32() // cargoPool.rotation.x
		b.GetFloat32() // cargoPool.rotation.y
		b.GetFloat32() // cargoPool.rotation.z
		b.GetFloat32() // cargoPool.rotation.w
	}
	for i := 0; int32(i) < poolCapacity; i++ {
		b.GetInt32le() // recycleID
	}
	fmt.Printf("cargo: capacity: %d, cursor %d\n", poolCapacity, cursor)
}

func parseCargoTraffic(b *Buffer) {
	checkVers(b, 0, "CargoTraffic")

	beltCursor := b.GetInt32le()
	b.GetInt32le() // beltCapacity
	beltRecycleCursor := b.GetInt32le()

	splitterCursor := b.GetInt32le()
	b.GetInt32le() // splitterCapacity
	splitterRecycleCursor := b.GetInt32le()

	pathCursor := b.GetInt32le()
	b.GetInt32le() // pathCapacity
	pathRecycleCursor := b.GetInt32le()

	for i := 1; int32(i) < beltCursor; i++ {
		parseBeltComponent(b)
	}
	for i := 0; int32(i) < beltRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < splitterCursor; i++ {
		parseSplitterComponent(b)
	}
	for i := 0; int32(i) < splitterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < pathCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			parseCargoPath(b)
		}
	}
	for i := 0; int32(i) < pathRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}

func parseBeltComponent(b *Buffer) {
	checkVers(b, 0, "BeltComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // speed = r.ReadInt32();
	b.GetInt32le() // segPathId = r.ReadInt32();
	b.GetInt32le() // segIndex = r.ReadInt32();
	b.GetInt32le() // segPivotOffset = r.ReadInt32();
	b.GetInt32le() // segLength = r.ReadInt32();
	b.GetInt32le() // outputId = r.ReadInt32();
	b.GetInt32le() // backInputId = r.ReadInt32();
	b.GetInt32le() // leftInputId = r.ReadInt32();
	b.GetInt32le() // rightInputId = r.ReadInt32();
}

func parseSplitterComponent(b *Buffer) {
	checkVers(b, 0, "SplitterComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // beltA = r.ReadInt32();
	b.GetInt32le() // beltB = r.ReadInt32();
	b.GetInt32le() // beltC = r.ReadInt32();
	b.GetInt32le() // beltD = r.ReadInt32();
	b.GetInt32le() // input0 = r.ReadInt32();
	b.GetInt32le() // input1 = r.ReadInt32();
	b.GetInt32le() // input2 = r.ReadInt32();
	b.GetInt32le() // input3 = r.ReadInt32();
	b.GetInt32le() // output0 = r.ReadInt32();
	b.GetInt32le() // output1 = r.ReadInt32();
	b.GetInt32le() // output2 = r.ReadInt32();
	b.GetInt32le() // output3 = r.ReadInt32();
	b.GetBoolean() // inPriority = r.ReadBoolean();
	b.GetBoolean() // outPriority = r.ReadBoolean();
	b.GetInt32le() // outFilter = r.ReadInt32();
}

func parseCargoPath(b *Buffer) {
	checkVers(b, 0, "CargoPath")

	b.GetInt32le() // id
	b.GetInt32le() // capacity
	bufferLength := b.GetInt32le()
	b.GetInt32le() // chunkCapacity
	chunkCount := b.GetInt32le()
	b.GetInt32le() // updateLen
	b.GetBoolean() // closed
	b.GetInt32le() // outputPathIdForImport
	b.GetInt32le() // outputIndex
	beltCount := b.GetInt32le()
	pathCount := b.GetInt32le()

	b.GetBytes(int(bufferLength)) // buffer...
	for i := 0; int32(i) < chunkCount; i++ {
		b.GetInt32le() // chunk component +0
		b.GetInt32le() // chunk component +1
		b.GetInt32le() // chunk component +2
	}
	for i := 0; int32(i) < bufferLength; i++ {
		b.GetFloat32() // pointPos.x
		b.GetFloat32() // pointPos.y
		b.GetFloat32() // pointPos.z
		b.GetFloat32() // pointRot.x
		b.GetFloat32() // pointRot.y
		b.GetFloat32() // pointRot.z
		b.GetFloat32() // pointRot.w
	}
	for i := 0; int32(i) < beltCount; i++ {
		b.GetInt32le() // belt ID?
	}
	for i := 0; int32(i) < pathCount; i++ {
		b.GetInt32le() // path ID?
	}
}

func parseTankComponent(b *Buffer) {

}

func parseFactoryStorage(b *Buffer) {
	checkVers(b, 0, "FactoryStorage")

	storageCursor := b.GetInt32le()
	storageCapacity := b.GetInt32le()
	storageRecycleCursor := b.GetInt32le()
	fmt.Printf("Storage cursor %d, capacity %d, recycleCursor %d\n",
		storageCursor, storageCapacity, storageRecycleCursor)
	for i := 1; int32(i) < storageCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			size := b.GetInt32le()
			fmt.Printf("  storage size %d\n", size)
			parseStorageComponent(b)
		}
	}
	for i := 0; int32(i) < storageRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	tankCursor := b.GetInt32le()
	tankCapacity := b.GetInt32le()
	tankRecycleCursor := b.GetInt32le()
	fmt.Printf("Tank cursor %d, capacity %d, recycleCursor %d\n",
		tankCursor, tankCapacity, tankRecycleCursor)
	for i := 1; int32(i) < tankCursor; i++ {
		parseTankComponent(b)
	}
	for i := 0; int32(i) < tankRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

}

func parsePowerSystem(b *Buffer) {
	checkVers(b, 0, "PowerSystem")

	genCapacity := b.GetInt32le()
	genCursor := b.GetInt32le()
	genRecycleCursor := b.GetInt32le()
	fmt.Printf(" generator: capacity %d, cursor %d, recycleCursor %d\n",
		genCapacity, genCursor, genRecycleCursor)
	for i := 1; int32(i) < genCursor; i++ {
		parsePowerGeneratorComponent(b)
	}
	for i := 0; int32(i) < genRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	nodeCapacity := b.GetInt32le()
	nodeCursor := b.GetInt32le()
	nodeRecycleCursor := b.GetInt32le()
	fmt.Printf(" node: capacity %d, cursor %d, recycleCursor %d\n",
		nodeCapacity, nodeCursor, nodeRecycleCursor)
	for i := 1; int32(i) < nodeCursor; i++ {
		parsePowerNodeComponent(b)
	}
	for i := 0; int32(i) < nodeRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	consumerCapacity := b.GetInt32le()
	consumerCursor := b.GetInt32le()
	consumerRecycleCursor := b.GetInt32le()
	fmt.Printf(" consumer: capacity %d, cursor %d, recycleCursor %d\n",
		consumerCapacity, consumerCursor, consumerRecycleCursor)
	for i := 1; int32(i) < consumerCursor; i++ {
		parsePowerConsumerComponent(b)
	}
	for i := 0; int32(i) < consumerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	accumulatorCapacity := b.GetInt32le()
	accumulatorCursor := b.GetInt32le()
	accumulatorRecycleCursor := b.GetInt32le()
	fmt.Printf(" accumulator: capacity %d, cursor %d, recycleCursor %d\n",
		accumulatorCapacity, accumulatorCursor, accumulatorRecycleCursor)
	for i := 1; int32(i) < accumulatorCursor; i++ {
		parsePowerAccumulatorComponent(b)
	}
	for i := 0; int32(i) < accumulatorRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	exchangerCapacity := b.GetInt32le()
	exchangerCursor := b.GetInt32le()
	exchangerRecycleCursor := b.GetInt32le()
	fmt.Printf(" exchanger: capacity %d, cursor %d, recycleCursor %d\n",
		exchangerCapacity, exchangerCursor, exchangerRecycleCursor)
	for i := 1; int32(i) < exchangerCursor; i++ {
		parsePowerExchangerComponent(b)
	}
	for i := 0; int32(i) < exchangerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	networkCapacity := b.GetInt32le()
	networkCursor := b.GetInt32le()
	networkRecycleCursor := b.GetInt32le()
	fmt.Printf(" network: capacity %d, cursor %d, recycleCursor %d\n",
		networkCapacity, networkCursor, networkRecycleCursor)
	for i := 1; int32(i) < networkCursor; i++ {
		if b.GetInt32le() == 1 {
			parsePowerNetwork(b)
		}
	}
	for i := 0; int32(i) < networkRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

}

func parsePowerGeneratorComponent(b *Buffer) {
	checkVers(b, 0, "PowerGeneratorComponent")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt32le() // entityId = r.ReadInt32()
	b.GetInt32le() // networkId = r.ReadInt32()
	b.GetBoolean() // photovoltaic = r.ReadBoolean()
	b.GetBoolean() // wind = r.ReadBoolean()
	b.GetBoolean() // gamma = r.ReadBoolean()
	b.GetInt64le() // genEnergyPerTick = r.ReadInt64()
	b.GetInt64le() // useFuelPerTick = r.ReadInt64()
	b.GetInt16le() // fuelMask = r.ReadInt16()
	b.GetInt64le() // fuelEnergy = r.ReadInt64()
	b.GetInt16le() // curFuelId = r.ReadInt16()
	b.GetInt16le() // fuelId = r.ReadInt16()
	b.GetInt16le() // fuelCount = r.ReadInt16()
	b.GetInt64le() // fuelHeat = r.ReadInt64()
	b.GetInt32le() // catalystId = r.ReadInt32()
	b.GetInt32le() // catalystPoint = r.ReadInt32()
	b.GetInt32le() // productId = r.ReadInt32()
	b.GetFloat32() // productCount = r.ReadSingle()
	b.GetInt64le() // productHeat = r.ReadInt64()
	b.GetFloat32() // warmup = r.ReadSingle()
	b.GetFloat32() // ionEnhance = r.ReadSingle()
	b.GetFloat32() // x = r.ReadSingle()
	b.GetFloat32() // y = r.ReadSingle()
	b.GetFloat32() // z = r.ReadSingle()
}

func parsePowerNodeComponent(b *Buffer) {
	checkVers(b, 0, "PowerNodeComponent")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt32le() // entityId = r.ReadInt32()
	b.GetInt32le() // networkId = r.ReadInt32()
	b.GetBoolean() // isCharger = r.ReadBoolean()
	b.GetInt32le() // workEnergyPerTick = r.ReadInt32()
	b.GetInt32le() // idleEnergyPerTick = r.ReadInt32()
	b.GetInt32le() // requiredEnergy = r.ReadInt32()
	b.GetFloat32() // powerPoint.x = r.ReadSingle()
	b.GetFloat32() // powerPoint.y = r.ReadSingle()
	b.GetFloat32() // powerPoint.z = r.ReadSingle()
	b.GetFloat32() // connectDistance = r.ReadSingle()
	b.GetFloat32() // coverRadius = r.ReadSingle()
}

func parsePowerConsumerComponent(b *Buffer) {
	checkVers(b, 0, "PowerConsumerComponent")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt32le() // entityId = r.ReadInt32()
	b.GetInt32le() // networkId = r.ReadInt32()
	b.GetFloat32() // plugPos.x = r.ReadSingle()
	b.GetFloat32() // plugPos.y = r.ReadSingle()
	b.GetFloat32() // plugPos.z = r.ReadSingle()
	b.GetInt64le() // requiredEnergy = r.ReadInt64()
	b.GetInt64le() // servedEnergy = r.ReadInt64()
	b.GetInt64le() // workEnergyPerTick = r.ReadInt64()
	b.GetInt64le() // idleEnergyPerTick = r.ReadInt64()
}

func parsePowerAccumulatorComponent(b *Buffer) {
	checkVers(b, 0, "PowerAccumulatorComponent")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt32le() // entityId = r.ReadInt32()
	b.GetInt32le() // networkId = r.ReadInt32()
	b.GetInt64le() // inputEnergyPerTick = r.ReadInt64()
	b.GetInt64le() // outputEnergyPerTick = r.ReadInt64()
	b.GetInt64le() // curEnergy = r.ReadInt64()
	b.GetInt64le() // maxEnergy = r.ReadInt64()
}

func parsePowerExchangerComponent(b *Buffer) {
	checkVers(b, 1, "PowerExchangerComponent")

	b.GetInt32le() // id = r.ReadInt32()
	b.GetInt32le() // entityId = r.ReadInt32()
	b.GetInt32le() // networkId = r.ReadInt32()
	b.GetInt16le() // emptyCount = r.ReadInt16()
	b.GetInt16le() // fullCount = r.ReadInt16()
	b.GetFloat32() // targetState = r.ReadSingle()
	b.GetFloat32() // state = r.ReadSingle()
	b.GetInt64le() // energyPerTick = r.ReadInt64()
	b.GetInt64le() // curPoolEnergy = r.ReadInt64()
	b.GetInt64le() // poolMaxEnergy = r.ReadInt64()
	b.GetInt32le() // emptyId = r.ReadInt32()
	b.GetInt32le() // fullId = r.ReadInt32()
	b.GetInt32le() // belt0 = r.ReadInt32()
	b.GetInt32le() // belt1 = r.ReadInt32()
	b.GetInt32le() // belt2 = r.ReadInt32()
	b.GetInt32le() // belt3 = r.ReadInt32()
	b.GetBoolean() // isOutput0 = r.ReadBoolean()
	b.GetBoolean() // isOutput1 = r.ReadBoolean()
	b.GetBoolean() // isOutput2 = r.ReadBoolean()
	b.GetBoolean() // isOutput3 = r.ReadBoolean()
	b.GetInt32le() // outputSlot = r.ReadInt32()
	b.GetInt32le() // inputSlot = r.ReadInt32()
	b.GetInt32le() // outputRectify = r.ReadInt32()
	b.GetInt32le() // inputRectify = r.ReadInt32()
}

func parsePowerNetwork(b *Buffer) {
	checkVers(b, 0, "PowerNetwork")

	id := b.GetInt32le() // id
	nodeCount := b.GetInt32le()
	consumersCount := b.GetInt32le()
	generatorsCount := b.GetInt32le()
	accumulatorsCount := b.GetInt32le()
	exchangersCount := b.GetInt32le()
	fmt.Printf("power network %d: %d nodes, %d consumers, %d generators, %d accumulators, %d exchangers\n",
		id, nodeCount, consumersCount, generatorsCount, accumulatorsCount, exchangersCount)

	for i := 0; int32(i) < nodeCount; i++ {
		parseNode(b)
	}

	for i := 0; int32(i) < consumersCount; i++ {
		b.GetInt32le() // id
	}

	for i := 0; int32(i) < generatorsCount; i++ {
		b.GetInt32le() // id
	}

	for i := 0; int32(i) < accumulatorsCount; i++ {
		b.GetInt32le() // id
	}

	for i := 0; int32(i) < exchangersCount; i++ {
		b.GetInt32le() // id
	}
}

func parseNode(b *Buffer) {
	os.Exit(1)
}

func parseFactorySystem(b *Buffer) {

}

func parsePlanetTransport(b *Buffer) {

}

func parseMonsterSystem(b *Buffer) {

}

func parsePlatformSystem(b *Buffer) {

}
