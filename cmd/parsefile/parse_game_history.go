package main

import "fmt"

func parseGameHistory(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameHistory version: %d", vers))
	}

	// TODO: parse recipeUnlocked
	count := b.GetInt32le()
	fmt.Printf("Unlocked recipe count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse tutorialUnlocked
	count = b.GetInt32le()
	fmt.Printf("Tutorial unlocked count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse featureKeys
	count = b.GetInt32le()
	fmt.Printf("Feature keys count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	// TODO: parse techStates
	count = b.GetInt32le()
	fmt.Printf("Tech states count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		unlocked := b.GetBoolean()
		curLevel := b.GetInt32le()
		maxLevel := b.GetInt32le()
		hashUnloaded := b.GetInt64le()
		hashNeeded := b.GetInt64le()
		fmt.Printf("  id %d: unlocked=%v curLevel %d, maxLevel %d, hashUploaded %d, hashNeeded %d\n",
			id, unlocked, curLevel, maxLevel, hashUnloaded, hashNeeded)
	}

	b.GetBoolean() // automanageLabItems
	b.GetInt32le() // currentTech

	// techQueue
	count = b.GetInt32le()
	fmt.Printf("Tech queue length: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		fmt.Printf("   id: %d\n", id)
	}

	b.GetInt32le() // universeObserveLevel
	b.GetFloat32() // solarSailLife
	b.GetFloat32() // solarEnergyLossRate
	b.GetBoolean() // useIonLayer
	b.GetInt32le() // inserterStackCount
	b.GetFloat32() // logisticDroneSpeed
	b.GetFloat32() // logisticDroneSpeedScale
	b.GetInt32le() // logisticDroneCarries
	b.GetFloat32() // logisticShipSailSpeed
	b.GetFloat32() // logisticShipWarpSpeed
	b.GetFloat32() // logisticShipSpeedScale
	b.GetBoolean() // logisticShipWarpDrive
	b.GetInt32le() // logisticShipCarries
	b.GetFloat32() // miningCostRate
	b.GetFloat32() // miningSpeedScale
	b.GetInt32le() // storageLevel
	b.GetInt32le() // labLevel
	b.GetInt32le() // techSpeed
	b.GetFloat32() // dysonNodeLatitide
	b.GetInt64le() // universeMatrixPointUploaded
	b.GetBoolean() // missionAccomplished
}
