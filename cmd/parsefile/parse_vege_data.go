package main

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
