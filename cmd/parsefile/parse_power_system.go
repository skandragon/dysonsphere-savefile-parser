package main

func parsePowerSystem(b *Buffer) {
	checkVers(b, 0, "PowerSystem")

	b.GetInt32le() // genCapacity
	genCursor := b.GetInt32le()
	genRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < genCursor; i++ {
		parsePowerGeneratorComponent(b)
	}
	for i := 0; int32(i) < genRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // nodeCapacity
	nodeCursor := b.GetInt32le()
	nodeRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < nodeCursor; i++ {
		parsePowerNodeComponent(b)
	}
	for i := 0; int32(i) < nodeRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // consumerCapacity
	consumerCursor := b.GetInt32le()
	consumerRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < consumerCursor; i++ {
		parsePowerConsumerComponent(b)
	}
	for i := 0; int32(i) < consumerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // accumulatorCapacity
	accumulatorCursor := b.GetInt32le()
	accumulatorRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < accumulatorCursor; i++ {
		parsePowerAccumulatorComponent(b)
	}
	for i := 0; int32(i) < accumulatorRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // exchangerCapacity
	exchangerCursor := b.GetInt32le()
	exchangerRecycleCursor := b.GetInt32le()
	for i := 1; int32(i) < exchangerCursor; i++ {
		parsePowerExchangerComponent(b)
	}
	for i := 0; int32(i) < exchangerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // networkCapacity
	networkCursor := b.GetInt32le()
	networkRecycleCursor := b.GetInt32le()
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

	b.GetInt32le() // id
	nodeCount := b.GetInt32le()
	consumersCount := b.GetInt32le()
	generatorsCount := b.GetInt32le()
	accumulatorsCount := b.GetInt32le()
	exchangersCount := b.GetInt32le()

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
