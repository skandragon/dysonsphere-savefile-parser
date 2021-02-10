package main

import (
	"fmt"
)

type EntityAnimation struct {
	time          float32
	prepareLength float32
	workingLength float32
	state         uint32
	power         float32
}

func (e *EntityAnimation) String() string {
	return fmt.Sprintf("EntityAnim{%f,%f,%f,%d,%f}",
		e.time, e.prepareLength, e.workingLength,
		e.state, e.power)
}

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
	entityAnimations := make([]*EntityAnimation, entityCursor)
	for i := 1; int32(i) < entityCursor; i++ {
		entityAnimations[i] = &EntityAnimation{
			time:          b.GetFloat32(),
			prepareLength: b.GetFloat32(),
			workingLength: b.GetFloat32(),
			state:         b.GetUInt32le(),
			power:         b.GetFloat32(),
		}
	}
	//fmt.Printf(" entityAnimations: %v\n", entityAnimations)
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
	checkVers(b, 0, "TankComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // lastTankId = r.ReadInt32();
	b.GetInt32le() // nextTankId = r.ReadInt32();
	b.GetInt32le() // belt0 = r.ReadInt32();
	b.GetInt32le() // belt1 = r.ReadInt32();
	b.GetInt32le() // belt2 = r.ReadInt32();
	b.GetInt32le() // belt3 = r.ReadInt32();
	b.GetBoolean() // isOutput0 = r.ReadBoolean();
	b.GetBoolean() // isOutput1 = r.ReadBoolean();
	b.GetBoolean() // isOutput2 = r.ReadBoolean();
	b.GetBoolean() // isOutput3 = r.ReadBoolean();
	b.GetInt32le() // fluidStorageCount = r.ReadInt32();
	b.GetInt32le() // currentCount = r.ReadInt32();
	b.GetInt32le() // fluidId = r.ReadInt32();
	b.GetBoolean() // outputSwitch = r.ReadBoolean();
	b.GetBoolean() // inputSwitch = r.ReadBoolean();
	b.GetBoolean() // isBottom = r.ReadBoolean();
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
			if id != int32(i) {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			size := b.GetInt32le()
			fmt.Printf("  storage size %d\n", size)
			parseStorageComponent(b)
		}
	}
	for i := int32(0); i < storageRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	tankCapacity := b.GetInt32le()
	tankCursor := b.GetInt32le()
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
	for i := 0; int32(i) < networkCursor; i++ {
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
	checkVers(b, 0, "Node")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetFloat32() // x = r.ReadSingle();
	b.GetFloat32() // y = r.ReadSingle();
	b.GetFloat32() // z = r.ReadSingle();
	b.GetFloat32() // connDistance2 = r.ReadSingle();
	b.GetFloat32() // coverRadius2 = r.ReadSingle();
	b.GetInt32le() // genId = r.ReadInt32();
	b.GetInt32le() // accId = r.ReadInt32();
	b.GetInt32le() // excId = r.ReadInt32();
	connectionCount := b.GetInt32le()
	lineCount := b.GetInt32le()
	consumerCount := b.GetInt32le()

	for i := 0; int32(i) < connectionCount; i++ {
		b.GetInt32le() // id
	}

	for i := 0; int32(i) < lineCount; i++ {
		b.GetInt32le() // id
	}

	for i := 0; int32(i) < consumerCount; i++ {
		b.GetInt32le() // id
	}
}

func parseFactorySystem(b *Buffer) {
	checkVers(b, 0, "FactorySystem")

	minerCapacity := b.GetInt32le() // minerCapacity
	minerCursor := b.GetInt32le()
	minerRecycleCursor := b.GetInt32le()
	fmt.Printf("miner: capacity %d, cursor %d, recycleCursor %d\n",
		minerCapacity, minerCursor, minerRecycleCursor)
	for i := uint32(1); i < uint32(minerCursor); i++ {
		parseMinerComponent(b)
	}
	for i := int32(0); i < minerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	inserterCapacity := b.GetInt32le()
	inserterCursor := b.GetInt32le()
	inserterRecycleCursor := b.GetInt32le()
	fmt.Printf("inserter: capacity %d, cursor %d, recycleCursor %d\n",
		inserterCapacity, inserterCursor, inserterRecycleCursor)
	for i := int32(1); i < inserterCursor; i++ {
		parseInserterComponent(b)
	}
	for i := int32(0); i < inserterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	assemblerCapacity := b.GetInt32le()
	assemblerCursor := b.GetInt32le()
	assemblerRecycleCursor := b.GetInt32le()
	fmt.Printf("assembler: capacity %d, cursor %d, recycleCursor %d\n",
		assemblerCapacity, assemblerCursor, assemblerRecycleCursor)
	for i := uint32(1); i < uint32(assemblerCursor); i++ {
		parseAssemblerComponent(b)
	}
	for i := int32(0); i < assemblerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	fractionateCapacity := b.GetInt32le()
	fractionateCursor := b.GetInt32le()
	fractionateRecycleCursor := b.GetInt32le()
	fmt.Printf("fractionate: capacity %d, cursor %d, recycleCursor %d\n",
		fractionateCapacity, fractionateCursor, fractionateRecycleCursor)
	for i := uint32(1); i < uint32(fractionateCursor); i++ {
		parseFractionateComponent(b)
	}
	for i := int32(0); i < fractionateRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	ejectorCapacity := b.GetInt32le()
	ejectorCursor := b.GetInt32le()
	ejectorRecycleCursor := b.GetInt32le()
	fmt.Printf("ejector: capacity %d, cursor %d, recycleCursor %d\n",
		ejectorCapacity, ejectorCursor, ejectorRecycleCursor)
	for i := uint32(1); i < uint32(ejectorCursor); i++ {
		parseEjectorComponent(b)
	}
	for i := int32(0); i < ejectorRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	siloCapacity := b.GetInt32le()
	siloCursor := b.GetInt32le()
	siloRecycleCursor := b.GetInt32le()
	fmt.Printf("silo: capacity %d, cursor %d, recycleCursor %d\n",
		siloCapacity, siloCursor, siloRecycleCursor)
	for i := uint32(1); i < uint32(siloCursor); i++ {
		parseSiloComponent(b)
	}
	for i := int32(0); i < siloRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	labCapacity := b.GetInt32le()
	labCursor := b.GetInt32le()
	labRecycleCursor := b.GetInt32le()
	fmt.Printf("lab: capacity %d, cursor %d, recycleCursor %d\n",
		labCapacity, labCursor, labRecycleCursor)
	for i := uint32(1); i < uint32(labCursor); i++ {
		parseLabComponent(b)
	}
	for i := int32(0); i < labRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}

func parseMinerComponent(b *Buffer) {
	checkVers(b, 0, "MinerComponent")

	b.GetInt32le()       // id = r.ReadInt32();
	b.GetInt32le()       // entityId = r.ReadInt32();
	b.GetInt32le()       // pcId = r.ReadInt32();
	ty := b.GetInt32le() // type = (EMinerType)r.ReadInt32();
	fmt.Printf("Miner type %d (%s)\n", ty, MinerType(ty))
	b.GetInt32le() // speed = r.ReadInt32();
	b.GetInt32le() // time = r.ReadInt32();
	b.GetInt32le() // period = r.ReadInt32();
	b.GetInt32le() // insertTarget = r.ReadInt32();
	b.GetInt32le() // workstate = (EWorkState)r.ReadInt32();
	veinCount := b.GetInt32le()
	for i := int32(0); i < veinCount; i++ {
		b.GetInt32le() // veinID
	}

	b.GetInt32le() // currentVeinIndex = r.ReadInt32();
	b.GetInt32le() // minimumVeinAmount = r.ReadInt32();
	b.GetInt32le() // productId = r.ReadInt32();
	b.GetInt32le() // productCount = r.ReadInt32();
	b.GetInt32le() // seed = r.ReadUInt32();
}

func parseInserterComponent(b *Buffer) {
	checkVers(b, 0, "InserterComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // pcId = r.ReadInt32();
	b.GetInt32le() // stage = (EInserterStage)r.ReadInt32();
	b.GetInt32le() // speed = r.ReadInt32();
	b.GetInt32le() // time = r.ReadInt32();
	b.GetInt32le() // stt = r.ReadInt32();
	b.GetInt32le() // delay = r.ReadInt32();
	b.GetInt32le() // pickTarget = r.ReadInt32();
	b.GetInt32le() // insertTarget = r.ReadInt32();
	b.GetBoolean() // careNeeds = r.ReadBoolean();
	b.GetBoolean() // canStack = r.ReadBoolean();
	b.GetInt16le() // pickOffset = r.ReadInt16();
	b.GetInt16le() // insertOffset = r.ReadInt16();
	b.GetInt32le() // filter = r.ReadInt32();
	b.GetInt32le() // itemId = r.ReadInt32();
	b.GetInt32le() // stackCount = r.ReadInt32();
	b.GetInt32le() // stackSize = r.ReadInt32();
	b.GetFloat32() // pos2.x = r.ReadSingle();
	b.GetFloat32() // pos2.y = r.ReadSingle();
	b.GetFloat32() // pos2.z = r.ReadSingle();
	b.GetFloat32() // rot2.x = r.ReadSingle();
	b.GetFloat32() // rot2.y = r.ReadSingle();
	b.GetFloat32() // rot2.z = r.ReadSingle();
	b.GetFloat32() // rot2.w = r.ReadSingle();
	b.GetInt16le() // t1 = r.ReadInt16();
	b.GetInt16le() // t2 = r.ReadInt16();
}

func parseAssemblerComponent(b *Buffer) {
	checkVers(b, 0, "AssemblerComponent")

	b.GetInt32le()             // id = r.ReadInt32();
	b.GetInt32le()             // entityId = r.ReadInt32();
	b.GetInt32le()             // pcId = r.ReadInt32();
	b.GetBoolean()             // replicating = r.ReadBoolean();
	b.GetBoolean()             // outputing = r.ReadBoolean();
	b.GetInt32le()             // speed
	b.GetInt32le()             // time = r.ReadInt32();
	recipeID := b.GetInt32le() // recipeId = r.ReadInt32();
	if recipeID > 0 {
		b.GetInt32le() // (ERecipeType)r.ReadInt32()
		b.GetInt32le() // timeSpent
		requireCount := b.GetInt32le()
		for i := int32(0); i < requireCount; i++ {
			b.GetInt32le() // require ID?
		}
		requiresCountsCount := b.GetInt32le()
		for i := int32(0); i < requiresCountsCount; i++ {
			b.GetInt32le() // requireCount
		}
		servedCount := b.GetInt32le()
		for i := int32(0); i < servedCount; i++ {
			b.GetInt32le() // served ID?
		}
		needsCount := b.GetInt32le()
		for i := int32(0); i < needsCount; i++ {
			b.GetInt32le() // needs ID?
		}
		productsCount := b.GetInt32le()
		for i := int32(0); i < productsCount; i++ {
			b.GetInt32le() // products ID?
		}
		productCountCount := b.GetInt32le()
		for i := int32(0); i < productCountCount; i++ {
			b.GetInt32le() // product count
		}
		producedCount := b.GetInt32le()
		for i := int32(0); i < producedCount; i++ {
			b.GetInt32le() // produced ID?
		}
	}
}

func parseFractionateComponent(b *Buffer) {
	checkVers(b, 0, "FractionateComponent")

	b.GetInt32le()  // id = r.ReadInt32();
	b.GetInt32le()  // entityId = r.ReadInt32();
	b.GetInt32le()  // pcId = r.ReadInt32();
	b.GetInt32le()  // belt0 = r.ReadInt32();
	b.GetInt32le()  // belt1 = r.ReadInt32();
	b.GetInt32le()  // belt2 = r.ReadInt32();
	b.GetBoolean()  // isOutput0 = r.ReadBoolean();
	b.GetBoolean()  // isOutput1 = r.ReadBoolean();
	b.GetBoolean()  // isOutput2 = r.ReadBoolean();
	b.GetBoolean()  // isWorking = r.ReadBoolean();
	b.GetFloat32()  // produceProb = r.ReadSingle();
	b.GetInt32le()  // need = r.ReadInt32();
	b.GetInt32le()  // product = r.ReadInt32();
	b.GetInt32le()  // needCurrCount = r.ReadInt32();
	b.GetInt32le()  // productCurrCount = r.ReadInt32();
	b.GetInt32le()  // oriProductCurrCount = r.ReadInt32();
	b.GetInt32le()  // progress = r.ReadInt32();
	b.GetBoolean()  // isRand = r.ReadBoolean();
	b.GetBoolean()  // fractionateSuccess = r.ReadBoolean();
	b.GetInt32le()  // needMaxCount = r.ReadInt32();
	b.GetInt32le()  // productMaxCount = r.ReadInt32();
	b.GetInt32le()  // oriProductMaxCount = r.ReadInt32();
	b.GetUInt32le() // seed = r.ReadUInt32();
}

func parseEjectorComponent(b *Buffer) {
	checkVers(b, 0, "EjectorComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // planetId = r.ReadInt32();
	b.GetInt32le() // pcId = r.ReadInt32();
	b.GetInt32le() // direction = r.ReadInt32();
	b.GetInt32le() // time = r.ReadInt32();
	b.GetBoolean() // fired = r.ReadBoolean();
	b.GetInt32le() // chargeSpend = r.ReadInt32();
	b.GetInt32le() // coldSpend = r.ReadInt32();
	b.GetInt32le() // bulletId = r.ReadInt32();
	b.GetInt32le() // bulletCount = r.ReadInt32();
	b.GetInt32le() // orbitId = r.ReadInt32();
	b.GetFloat32() // pivotY = r.ReadSingle();
	b.GetFloat32() // muzzleY = r.ReadSingle();
	b.GetFloat32() // localPosN.x = r.ReadSingle();
	b.GetFloat32() // localPosN.y = r.ReadSingle();
	b.GetFloat32() // localPosN.z = r.ReadSingle();
	b.GetFloat32() // localAlt = r.ReadSingle();
	b.GetFloat32() // localRot.x = r.ReadSingle();
	b.GetFloat32() // localRot.y = r.ReadSingle();
	b.GetFloat32() // localRot.z = r.ReadSingle();
	b.GetFloat32() // localRot.w = r.ReadSingle();
	b.GetFloat32() // localDir.x = r.ReadSingle();
	b.GetFloat32() // localDir.y = r.ReadSingle();
	b.GetFloat32() // localDir.z = r.ReadSingle();
}

func parseSiloComponent(b *Buffer) {
	checkVers(b, 0, "SiloComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // planetId = r.ReadInt32();
	b.GetInt32le() // pcId = r.ReadInt32();
	b.GetInt32le() // direction = r.ReadInt32();
	b.GetInt32le() // time = r.ReadInt32();
	b.GetBoolean() // fired = r.ReadBoolean();
	b.GetInt32le() // chargeSpend = r.ReadInt32();
	b.GetInt32le() // coldSpend = r.ReadInt32();
	b.GetInt32le() // bulletId = r.ReadInt32();
	b.GetInt32le() // bulletCount = r.ReadInt32();
	b.GetInt32le() // autoIndex = r.ReadInt32();
	b.GetBoolean() // hasNode = r.ReadBoolean();
	b.GetFloat32() // localPosN.x = r.ReadSingle();
	b.GetFloat32() // localPosN.y = r.ReadSingle();
	b.GetFloat32() // localPosN.z = r.ReadSingle();
	b.GetFloat32() // localRot.x = r.ReadSingle();
	b.GetFloat32() // localRot.y = r.ReadSingle();
	b.GetFloat32() // localRot.z = r.ReadSingle();
	b.GetFloat32() // localRot.w = r.ReadSingle();
}

func parseLabComponent(b *Buffer) {
	checkVers(b, 0, "LabComponent")

	b.GetInt32le()                 // id = r.ReadInt32();
	b.GetInt32le()                 // entityId = r.ReadInt32();
	b.GetInt32le()                 // pcId = r.ReadInt32();
	b.GetInt32le()                 // nextLabId = r.ReadInt32();
	b.GetBoolean()                 // replicating = r.ReadBoolean();
	b.GetBoolean()                 // outputing = r.ReadBoolean();
	b.GetInt32le()                 // time = r.ReadInt32();
	b.GetInt32le()                 // hashBytes = r.ReadInt32();
	researchMode := b.GetBoolean() // researchMode = r.ReadBoolean();
	recipeID := b.GetInt32le()     // recipeId = r.ReadInt32();
	b.GetInt32le()                 // techId = r.ReadInt32();
	if !researchMode && recipeID > 0 {
		b.GetInt32le() // timeSpent
		requireCount := b.GetInt32le()
		for i := int32(0); i < requireCount; i++ {
			b.GetInt32le() // require ID?
		}
		requiresCountsCount := b.GetInt32le()
		for i := int32(0); i < requiresCountsCount; i++ {
			b.GetInt32le() // requireCount
		}
		servedCount := b.GetInt32le()
		for i := int32(0); i < servedCount; i++ {
			b.GetInt32le() // served ID?
		}
		needsCount := b.GetInt32le()
		for i := int32(0); i < needsCount; i++ {
			b.GetInt32le() // needs ID?
		}
		productsCount := b.GetInt32le()
		for i := int32(0); i < productsCount; i++ {
			b.GetInt32le() // products ID?
		}
		productCountCount := b.GetInt32le()
		for i := int32(0); i < productCountCount; i++ {
			b.GetInt32le() // product count
		}
		producedCount := b.GetInt32le()
		for i := int32(0); i < producedCount; i++ {
			b.GetInt32le() // produced ID?
		}
	}
	if researchMode {
		matrixPoints := b.GetInt32le()
		for i := int32(0); i < matrixPoints; i++ {
			b.GetInt32le() // ?
		}
		matrixServed := b.GetInt32le()
		for i := int32(0); i < matrixServed; i++ {
			b.GetInt32le() // ?
		}
		needsCount := b.GetInt32le()
		for i := int32(0); i < needsCount; i++ {
			b.GetInt32le() // needs ID?
		}
	}
}

func parsePlanetTransport(b *Buffer) {
	checkVers(b, 0, "PlanetTransport")

	stationCursor := b.GetInt32le()
	stationCapackty := b.GetInt32le()
	stationRecycleCounter := b.GetInt32le()
	fmt.Printf("station: capacity %d, cursor %d, recycleCursor %d\n",
		stationCapackty, stationCursor, stationRecycleCounter)
	for i := int32(1); i < stationCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			if id != i {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			parseStationComponent(b)
		}
	}
	for i := 0; int32(i) < stationRecycleCounter; i++ {
		b.GetInt32le() // recycle id?
	}
}

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
	b.GetBoolean() // isStellar = r.ReadBoolean();
	hasName := b.GetInt32le()
	if hasName > 0 {
		b.GetString() // name = r.ReadString();
	}
	b.GetInt64le() // energy = r.ReadInt64();
	b.GetInt64le() // energyPerTick = r.ReadInt64();
	b.GetInt64le() // energyMax = r.ReadInt64();
	b.GetInt32le() // warperCount = r.ReadInt32();
	b.GetInt32le() // warperMaxCount = r.ReadInt32();
	b.GetInt32le() // idleDroneCount = r.ReadInt32();
	workDroneCount := b.GetInt32le()
	b.GetInt32le() // workDroneDatas.Length
	for i := int32(0); i < workDroneCount; i++ {
		parseDroneData(b)
	}
	for i := int32(0); i < workDroneCount; i++ {
		parseLocalLogisticOrder(b)
	}
	b.GetInt32le()                  // idleShipCount = r.ReadInt32();
	workShipCount := b.GetInt32le() // workShipCount = r.ReadInt32();
	b.GetInt64le()                  // idleShipIndices = r.ReadUInt64();
	b.GetInt64le()                  // workShipIndices = r.ReadUInt64();
	b.GetInt32le()                  // workShipDatas.Length
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
	b.GetInt32le() // isCollector = r.ReadBoolean();
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

func parseDroneData(b *Buffer) {
	checkVers(b, 0, "DroneData")

	b.GetInt32le() // begin.x = r.ReadSingle()
	b.GetInt32le() // begin.y = r.ReadSingle()
	b.GetInt32le() // begin.z = r.ReadSingle()
	b.GetInt32le() // end.x = r.ReadSingle()
	b.GetInt32le() // end.y = r.ReadSingle()
	b.GetInt32le() // end.z = r.ReadSingle()
	b.GetInt32le() // endId = r.ReadInt32()
	b.GetInt32le() // direction = r.ReadSingle()
	b.GetInt32le() // maxt = r.ReadSingle()
	b.GetInt32le() // t = r.ReadSingle()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // itemCount = r.ReadInt32()
	b.GetInt32le() // gene = r.ReadInt32()
}

func parseLocalLogisticOrder(b *Buffer) {
	checkVers(b, 0, "LocalLogisticsOrder")

	b.GetInt32le() // otherStationId = r.ReadInt32()
	b.GetInt32le() // thisIndex = r.ReadInt32()
	b.GetInt32le() // otherIndex = r.ReadInt32()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // thisOrdered = r.ReadInt32()
	b.GetInt32le() // otherOrdered = r.ReadInt32()
}

func parseShipData(b *Buffer) {
	checkVers(b, 0, "ShipData")

	b.GetInt32le() // stage = r.ReadInt32()
	b.GetInt32le() // planetA = r.ReadInt32()
	b.GetInt32le() // planetB = r.ReadInt32()
	b.GetFloat64() // uPos.x = r.ReadDouble()
	b.GetFloat64() // uPos.y = r.ReadDouble()
	b.GetFloat64() // uPos.z = r.ReadDouble()
	b.GetFloat32() // uVel.x = r.ReadSingle()
	b.GetFloat32() // uVel.y = r.ReadSingle()
	b.GetFloat32() // uVel.z = r.ReadSingle()
	b.GetFloat32() // uSpeed = r.ReadSingle()
	b.GetFloat32() // warpState = r.ReadSingle()
	b.GetFloat32() // uRot.x = r.ReadSingle()
	b.GetFloat32() // uRot.y = r.ReadSingle()
	b.GetFloat32() // uRot.z = r.ReadSingle()
	b.GetFloat32() // uRot.w = r.ReadSingle()
	b.GetFloat32() // uAngularVel.x = r.ReadSingle()
	b.GetFloat32() // uAngularVel.y = r.ReadSingle()
	b.GetFloat32() // uAngularVel.z = r.ReadSingle()
	b.GetFloat32() // uAngularSpeed = r.ReadSingle()
	b.GetFloat64() // pPosTemp.x = r.ReadDouble()
	b.GetFloat64() // pPosTemp.y = r.ReadDouble()
	b.GetFloat64() // pPosTemp.z = r.ReadDouble()
	b.GetFloat32() // pRotTemp.x = r.ReadSingle()
	b.GetFloat32() // pRotTemp.y = r.ReadSingle()
	b.GetFloat32() // pRotTemp.z = r.ReadSingle()
	b.GetFloat32() // pRotTemp.w = r.ReadSingle()
	b.GetInt32le() // otherGId = r.ReadInt32()
	b.GetInt32le() // direction = r.ReadInt32()
	b.GetFloat32() // t = r.ReadSingle()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // itemCount = r.ReadInt32()
	b.GetInt32le() // gene = r.ReadInt32()
	b.GetInt32le() // shipIndex = r.ReadInt32()
	b.GetInt32le() // warperCnt = r.ReadInt32()
}

func parseRemoteLogisticOrder(b *Buffer) {
	checkVers(b, 0, "RemoteLogisticOrder")

	b.GetInt32le() // otherStationGId = r.ReadInt32()
	b.GetInt32le() // thisIndex = r.ReadInt32()
	b.GetInt32le() // otherIndex = r.ReadInt32()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // thisOrdered = r.ReadInt32()
	b.GetInt32le() // otherOrdered = r.ReadInt32()
}

func parseStationStore(b *Buffer) {
	checkVers(b, 0, "StationStore")

	b.GetInt32le() // itemId = r.ReadInt32();
	b.GetInt32le() // count = r.ReadInt32();
	b.GetInt32le() // localOrder = r.ReadInt32();
	b.GetInt32le() // remoteOrder = r.ReadInt32();
	b.GetInt32le() // max = r.ReadInt32();
	b.GetInt32le() // localLogic = (ELogisticStorage)r.ReadInt32();
	b.GetInt32le() // remoteLogic = (ELogisticStorage)r.ReadInt32();
}

func parseMonsterSystem(b *Buffer) {
	checkVers(b, 0, "MonsterSystem")

	b.GetInt32le() // monsterCapacity
	monsterCursor := b.GetInt32le()
	monsterRecycleCursor := b.GetInt32le()
	for i := int32(1); i < monsterCursor; i++ {
		parseMonsterComponent(b)
	}
	for i := int32(0); i < monsterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}

func parseMonsterComponent(b *Buffer) {
	checkVers(b, 0, "MonsterComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetFloat32() // walkSpeed = r.ReadSingle();
	b.GetFloat32() // point0.x = r.ReadSingle();
	b.GetFloat32() // point0.y = r.ReadSingle();
	b.GetFloat32() // point0.z = r.ReadSingle();
	b.GetFloat32() // point1.x = r.ReadSingle();
	b.GetFloat32() // point1.y = r.ReadSingle();
	b.GetFloat32() // point1.z = r.ReadSingle();
	b.GetFloat32() // point2.x = r.ReadSingle();
	b.GetFloat32() // point2.y = r.ReadSingle();
	b.GetFloat32() // point2.z = r.ReadSingle();
	b.GetInt32le() // direction = r.ReadInt32();
	b.GetFloat32() // stopTime = r.ReadSingle();
	b.GetFloat32() // t = r.ReadSingle();
	b.GetFloat32() // stopCurrentTime = r.ReadSingle();
	b.GetInt32le() // monsterState = (EMonsterState)r.ReadInt32();
	b.GetFloat32() // stepDistance = r.ReadSingle();
}

func parsePlatformSystem(b *Buffer) {
	checkVers(b, 0, "PlatformSystem")

	reformDataLength := b.GetInt32le()
	b.GetBytes(int(reformDataLength)) // reformData
	reformOffsetsCount := b.GetInt32le()
	for i := int32(0); i < reformOffsetsCount; i++ {
		b.GetInt32le() // reformOffset
	}
}
