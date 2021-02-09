package main

import "fmt"

func parseGamePreferences(b *Buffer) {
	checkVers(b, 2, "GamePreferences")

	cameraX := b.GetFloat64()
	cameraY := b.GetFloat64()
	cameraZ := b.GetFloat64()
	fmt.Printf("Camera: %f %f %f\n", cameraX, cameraY, cameraZ)
	camRotX := b.GetFloat32()
	camRotY := b.GetFloat32()
	camRotZ := b.GetFloat32()
	camRotW := b.GetFloat32()
	fmt.Printf("Camera rotation: %f %f %f %f\n", camRotX, camRotY, camRotZ, camRotW)

	reformCursorSize := b.GetInt32le()
	fmt.Printf("Reform cursor size: %d\n", reformCursorSize)

	// replicationMultipliers, a key/value list
	count := b.GetInt32le()
	fmt.Printf("replicationMultipliers count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		key := b.GetInt32le()
		value := b.GetInt32le()
		fmt.Printf("replicationMultiplier: key %d, value %d\n", key, value)
	}

	detailPower := b.GetBoolean()
	detailVein := b.GetBoolean()
	detailSpaceGuide := b.GetBoolean()
	detailSign := b.GetBoolean()
	detailIcon := b.GetBoolean()
	fmt.Printf("detailPower=%v detailVein=%v detailSpaceGuide=%v detailSign=%v detailIcon=%v\n",
		detailPower, detailVein, detailSpaceGuide, detailSign, detailIcon)

	// currently displayed tutorials
	count = b.GetInt32le()
	fmt.Printf("Displayed tutorial count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		tutorial := b.GetInt32le()
		fmt.Printf("  tutorial: %d\n", tutorial)
	}
}
