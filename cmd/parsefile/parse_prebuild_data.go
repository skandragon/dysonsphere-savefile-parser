package main

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
