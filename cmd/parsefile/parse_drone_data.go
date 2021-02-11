package main

func parseDroneData(b *Buffer) {
	checkVers(b, 0, "DroneData")

	b.GetFloat32() // begin.x = r.ReadSingle()
	b.GetFloat32() // begin.y = r.ReadSingle()
	b.GetFloat32() // begin.z = r.ReadSingle()
	b.GetFloat32() // end.x = r.ReadSingle()
	b.GetFloat32() // end.y = r.ReadSingle()
	b.GetFloat32() // end.z = r.ReadSingle()
	b.GetInt32le() // endId = r.ReadInt32()
	b.GetFloat32() // direction = r.ReadSingle()
	b.GetFloat32() // maxt = r.ReadSingle()
	b.GetFloat32() // t = r.ReadSingle()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // itemCount = r.ReadInt32()
	b.GetInt32le() // gene = r.ReadInt32()
}
