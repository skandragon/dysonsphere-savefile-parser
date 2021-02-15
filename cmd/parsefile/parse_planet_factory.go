package main

import (
	"fmt"
)

// EntityAnimation holds the current state of an entity's animation logic.
type EntityAnimation struct {
	Time          float32
	PrepareLength float32
	WorkingLength float32
	State         uint32
	Power         float32
}

func (e *EntityAnimation) String() string {
	return fmt.Sprintf("EntityAnim{%f,%f,%f,%d,%f}",
		e.Time, e.PrepareLength, e.WorkingLength,
		e.State, e.Power)
}

// PlanetFactory holds information about a single planet's factories.
type PlanetFactory struct {
	PlanetID     int32              `json:"planet_id"`
	Star         int32              `json:"star"`
	PlanetNumber int32              `json:"planet_number"`
	Runtime      *PlanetDataRuntime `json:"runtime"`
	Veins        []*VeinData        `json:"veins"`
}

func parsePlanetFactory(b *Buffer, i int) *PlanetFactory {
	const VERSION = 1
	checkVers(b, VERSION, "PlanetFactory")

	pf := &PlanetFactory{}

	pf.PlanetID = b.GetInt32le()
	//fmt.Printf("PlanetFactory %d, planetID: %d\n", i, pf.PlanetID)
	pf.Star = pf.PlanetID / 100
	pf.PlanetNumber = pf.PlanetID % 100
	pf.Runtime = parsePlanetDataRuntime(b)

	b.GetInt32le() // entityCapacity
	entityCursor := b.GetInt32le()
	entityRecycleCursor := b.GetInt32le()
	//fmt.Printf(" entityCapacity %d, cursor %d, recycleCursor %d\n",
	//	entityCapacity, entityCursor, entityRecycleCursor)
	for i := 1; int32(i) < entityCursor; i++ {
		parseEntityData(b)
	}
	entityAnimations := make([]*EntityAnimation, entityCursor)
	for i := 1; int32(i) < entityCursor; i++ {
		entityAnimations[i] = &EntityAnimation{
			Time:          b.GetFloat32(),
			PrepareLength: b.GetFloat32(),
			WorkingLength: b.GetFloat32(),
			State:         b.GetUInt32le(),
			Power:         b.GetFloat32(),
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

	b.GetInt32le() // prebuildCapacity
	prebuildCursor := b.GetInt32le()
	prebuildRecycleCursor := b.GetInt32le()
	//fmt.Printf(" prebuildCapacity %d, cursor %d, recycleCursor %d\n",
	//	prebuildCapacity, prebuildCursor, prebuildRecycleCursor)
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

	b.GetInt32le() // vegeCapacity
	vegeCursor := b.GetInt32le()
	vegeRecycleCursor := b.GetInt32le()
	//fmt.Printf(" vegeCapacity %d, cursor %d, recycleCursor %d\n",
	//	vegeCapacity, vegeCursor, vegeRecycleCursor)
	for i := 1; int32(i) < vegeCursor; i++ {
		parseVegeData(b)
	}
	for i := 0; int32(i) < vegeRecycleCursor; i++ {
		b.GetInt32le() // vegeRecycle[num6]);
	}

	veinCapacity := b.GetInt32le()
	veinCursor := b.GetInt32le()
	veinRecycleCursor := b.GetInt32le()
	//fmt.Printf(" veinCapacity %d, cursor %d, recycleCursor %d\n",
	//	veinCapacity, veinCursor, veinRecycleCursor)
	veins := make([]*VeinData, veinCapacity)
	for i := 1; int32(i) < veinCursor; i++ {
		veins[i] = parseVeinData(b)
	}
	pf.Veins = compressVeinDataSlice(veins)
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

	return pf
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

func parseFactorySystem(b *Buffer) {
	checkVers(b, 0, "FactorySystem")

	b.GetInt32le() // minerCapacity
	minerCursor := b.GetInt32le()
	minerRecycleCursor := b.GetInt32le()
	//fmt.Printf("miner: capacity %d, cursor %d, recycleCursor %d\n",
	//	minerCapacity, minerCursor, minerRecycleCursor)
	for i := uint32(1); i < uint32(minerCursor); i++ {
		parseMinerComponent(b)
	}
	for i := int32(0); i < minerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // inserterCapacity
	inserterCursor := b.GetInt32le()
	inserterRecycleCursor := b.GetInt32le()
	//fmt.Printf("inserter: capacity %d, cursor %d, recycleCursor %d\n",
	//	inserterCapacity, inserterCursor, inserterRecycleCursor)
	for i := int32(1); i < inserterCursor; i++ {
		parseInserterComponent(b)
	}
	for i := int32(0); i < inserterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // assemblerCapacity
	assemblerCursor := b.GetInt32le()
	assemblerRecycleCursor := b.GetInt32le()
	//fmt.Printf("assembler: capacity %d, cursor %d, recycleCursor %d\n",
	//	assemblerCapacity, assemblerCursor, assemblerRecycleCursor)
	for i := uint32(1); i < uint32(assemblerCursor); i++ {
		parseAssemblerComponent(b)
	}
	for i := int32(0); i < assemblerRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // fractionateCapacity
	fractionateCursor := b.GetInt32le()
	fractionateRecycleCursor := b.GetInt32le()
	//fmt.Printf("fractionate: capacity %d, cursor %d, recycleCursor %d\n",
	//	fractionateCapacity, fractionateCursor, fractionateRecycleCursor)
	for i := uint32(1); i < uint32(fractionateCursor); i++ {
		parseFractionateComponent(b)
	}
	for i := int32(0); i < fractionateRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // ejectorCapacity
	ejectorCursor := b.GetInt32le()
	ejectorRecycleCursor := b.GetInt32le()
	//fmt.Printf("ejector: capacity %d, cursor %d, recycleCursor %d\n",
	//	ejectorCapacity, ejectorCursor, ejectorRecycleCursor)
	for i := uint32(1); i < uint32(ejectorCursor); i++ {
		parseEjectorComponent(b)
	}
	for i := int32(0); i < ejectorRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // siloCapacity
	siloCursor := b.GetInt32le()
	siloRecycleCursor := b.GetInt32le()
	//fmt.Printf("silo: capacity %d, cursor %d, recycleCursor %d\n",
	//	siloCapacity, siloCursor, siloRecycleCursor)
	for i := uint32(1); i < uint32(siloCursor); i++ {
		parseSiloComponent(b)
	}
	for i := int32(0); i < siloRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	b.GetInt32le() // labCapacity
	labCursor := b.GetInt32le()
	labRecycleCursor := b.GetInt32le()
	//fmt.Printf("lab: capacity %d, cursor %d, recycleCursor %d\n",
	//	labCapacity, labCursor, labRecycleCursor)
	for i := uint32(1); i < uint32(labCursor); i++ {
		parseLabComponent(b)
	}
	for i := int32(0); i < labRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}

func parseMinerComponent(b *Buffer) {
	checkVers(b, 0, "MinerComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // pcId = r.ReadInt32();
	b.GetInt32le() // type = (EMinerType)r.ReadInt32();
	//fmt.Printf("Miner type %d (%s)\n", ty, MinerType(ty))
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
