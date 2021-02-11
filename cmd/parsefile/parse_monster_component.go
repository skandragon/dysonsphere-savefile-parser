package main

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
