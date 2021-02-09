package main

import "fmt"

func parsePlayer(b *Buffer) {
	checkVers(b, 1, "Player")

	b.GetInt32le() // planetID
	b.GetFloat32() // position.x
	b.GetFloat32() // position.y
	b.GetFloat32() // uPosition.z
	b.GetFloat64() // uPosition.x
	b.GetFloat64() // uPosition.y
	b.GetFloat64() // uPosition.z
	b.GetFloat32() // uRotation.x
	b.GetFloat32() // uRotation.y
	b.GetFloat32() // uRotation.z
	b.GetFloat32() // uRotation.w
	b.GetInt32le() // movementState
	b.GetFloat32() // warpState
	b.GetBoolean() // warpCommand
	b.GetFloat64() // uVelocity.x
	b.GetFloat64() // uVelocity.y
	b.GetFloat64() // uVelocity.z
	b.GetInt32le() // inhandItemID
	b.GetInt32le() // inhandItemCount

	parseMecha(b)

	parseStorageComponent(b)

	parsePlayerNavigation(b)

	b.GetInt32le() // sandCount
}

func parseMecha(b *Buffer) {
	checkVers(b, 0, "Mecha")

	b.GetFloat64() // coreEnergyCap
	b.GetFloat64() // coreEnergy
	b.GetFloat64() // corePowerGen
	b.GetFloat64() // reactorPowerGen
	b.GetFloat64() // reactorEnergy
	b.GetInt32le() // reactorItemID

	// reactorStorage
	parseStorageComponent(b)

	// warpStorage
	parseStorageComponent(b)

	b.GetFloat64() // walkPower
	b.GetFloat64() // jumpPower

	b.GetFloat64() // thrustPowerPerAcc = r.ReadDouble();
	b.GetFloat64() // warpKeepingPowerPerSpeed = r.ReadDouble();
	b.GetFloat64() // warpStartPowerPerSpeed = r.ReadDouble();
	b.GetFloat64() // miningPower = r.ReadDouble();
	b.GetFloat64() // replicatePower = r.ReadDouble();
	b.GetFloat64() // researchPower = r.ReadDouble();
	b.GetFloat64() // droneEjectEnergy = r.ReadDouble();
	b.GetFloat64() // droneEnergyPerMeter = r.ReadDouble();
	b.GetInt32le() // coreLevel = r.ReadInt32();
	b.GetInt32le() // thrusterLevel = r.ReadInt32();
	b.GetFloat32() // miningSpeed = r.ReadSingle();
	b.GetFloat32() // replicateSpeed = r.ReadSingle();
	b.GetFloat32() // walkSpeed = r.ReadSingle();
	b.GetFloat32() // jumpSpeed = r.ReadSingle();
	b.GetFloat32() // maxSailSpeed = r.ReadSingle();
	b.GetFloat32() // maxWarpSpeed = r.ReadSingle();
	b.GetFloat32() // buildArea = r.ReadSingle();

	parseMechaForge(b)

	parseMechaLab(b)

	droneCount := b.GetInt32le()
	b.GetFloat32() // droneSpeed
	b.GetInt32le() // droneMovement
	for i := 0; int32(i) < droneCount; i++ {
		parseMechaDrone(b)
	}
}

func parseMechaDrone(b *Buffer) {
	checkVers(b, 0, "MechaDrone")

	b.GetInt32le() // stage = r.ReadInt32();
	b.GetFloat32() // position.x = r.ReadSingle();
	b.GetFloat32() // position.y = r.ReadSingle();
	b.GetFloat32() // position.z = r.ReadSingle();
	b.GetFloat32() // target.x = r.ReadSingle();
	b.GetFloat32() // target.y = r.ReadSingle();
	b.GetFloat32() // target.z = r.ReadSingle();
	b.GetFloat32() // forward.x = r.ReadSingle();
	b.GetFloat32() // forward.y = r.ReadSingle();
	b.GetFloat32() // forward.z = r.ReadSingle();
	b.GetFloat32() // speed = r.ReadSingle();
	b.GetInt32le() // movement = r.ReadInt32();
	b.GetInt32le() // targetObject = r.ReadInt32();
	b.GetFloat32() // progress = r.ReadSingle();
	b.GetFloat32() // initialVector.x = r.ReadSingle();
	b.GetFloat32() // initialVector.y = r.ReadSingle();
	b.GetFloat32() // initialVector.z = r.ReadSingle();
}

func parseStorageComponent(b *Buffer) {
	checkVers(b, 1, "StorageComponent")

	id := b.GetInt32le()
	entityID := b.GetInt32le()
	previous := b.GetInt32le()
	next := b.GetInt32le()
	bottom := b.GetInt32le()
	top := b.GetInt32le()
	fmt.Printf("id %d, entityId %d, previous %d, next %d, top %d, bottom %d\n",
		id, entityID, previous, next, bottom, top)
	storageType := b.GetInt32le() // type = (EStorageType)r.ReadInt32();
	size := b.GetInt32le()
	fmt.Printf("  storage type %d, size %d\n", storageType, size)
	b.GetInt32le() // bans = r.ReadInt32();
	for i := 0; int32(i) < size; i++ {
		b.GetInt32le() // grids[i].itemId = r.ReadInt32();
		b.GetInt32le() // grids[i].filter = r.ReadInt32();
		b.GetInt32le() // grids[i].count = r.ReadInt32();
		b.GetInt32le() // grids[i].stackSize = r.ReadInt32();
	}

}

func parsePlayerNavigation(b *Buffer) {
	checkVers(b, 0, "PlayerNavigation")

	b.GetBoolean() // navigating = r.ReadBoolean();
	b.GetInt32le() // naviAstroId = r.ReadInt32();
	b.GetFloat64() // naviTarget.x = r.ReadDouble();
	b.GetFloat64() // naviTarget.y = r.ReadDouble();
	b.GetFloat64() // naviTarget.z = r.ReadDouble();
	b.GetBoolean() // useFly = r.ReadBoolean();
	b.GetBoolean() // useSail = r.ReadBoolean();
	b.GetBoolean() // useWarp = r.ReadBoolean();
	b.GetInt32le() // stage = (ENaviStage)r.ReadInt32();
	b.GetFloat64() // flyThreshold = r.ReadDouble();
	b.GetFloat64() // sailThreshold = r.ReadDouble();
	b.GetFloat64() // warpThreshold = r.ReadDouble();
	b.GetFloat64() // maxSailSpeed = r.ReadDouble();
}

func parseMechaForge(b *Buffer) {
	checkVers(b, 0, "MechaForge")

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		parseForgeTask(b)
	}
}

func parseForgeTask(b *Buffer) {
	checkVers(b, 0, "ForgeTask")

	b.GetInt32le() // recipeId = r.ReadInt32();
	b.GetInt32le() // count = r.ReadInt32();
	b.GetInt32le() // tick = r.ReadInt32();
	b.GetInt32le() // tickSpend = r.ReadInt32();
	itemIDCount := b.GetInt32le()
	productIDCount := b.GetInt32le()
	for i := 0; i < int(itemIDCount); i++ {
		b.GetInt32le() // itemIds[i] = r.ReadInt32();
		b.GetInt32le() // itemCounts[i] = r.ReadInt32();
		b.GetInt32le() // served[i] = r.ReadInt32();
	}
	for i := 0; i < int(productIDCount); i++ {
		b.GetInt32le() // productIds[j] = r.ReadInt32();
		b.GetInt32le() // productCounts[j] = r.ReadInt32();
		b.GetInt32le() // produced[j] = r.ReadInt32();
	}
	b.GetInt32le() // parentTaskIndex = r.ReadInt32();
}

func parseMechaLab(b *Buffer) {
	checkVers(b, 0, "MechaLab")

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // key
		b.GetInt32le() // value
	}
}
