package main

func parseGamePreferences(b *Buffer) {
	checkVers(b, 2, "GamePreferences")

	b.GetFloat64() // cameraX
	b.GetFloat64() // cameraY
	b.GetFloat64() // cameraZ
	b.GetFloat32() // camRotX
	b.GetFloat32() // camRotY
	b.GetFloat32() // camRotZ
	b.GetFloat32() // camRotW

	b.GetInt32le() //reformCursorSize

	// replicationMultipliers, a key/value list
	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // key (this is the item the user can replicate)
		b.GetInt32le() // the per-click craft count, from 1-10 currently
	}

	b.GetBoolean() // detailPower
	b.GetBoolean() // detailVein
	b.GetBoolean() // detailSpaceGuide
	b.GetBoolean() // detailSign
	b.GetBoolean() // detailIcon

	// currently displayed tutorials
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // tutorialID
	}
}
