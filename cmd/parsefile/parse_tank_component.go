package main

func parseTankComponent(b *Buffer) {
	checkVers(b, 0, "TankComponent")

	b.GetInt32le() // id = r.ReadInt32();
	b.GetInt32le() // entityId = r.ReadInt32();
	b.GetInt32le() // lastTankId = r.ReadInt32();
	b.GetInt32le() // nextTankId = r.ReadInt32();
	b.GetInt32le() // belt0 = r.ReadInt32();
	b.GetInt32le() // belt1 = r.ReadInt32();
	b.GetInt32le() // belt2 = r.ReadInt32();
	b.GetInt32le() // belt3 = r.ReadInt32();
	b.GetBoolean() // isOutput0 = r.ReadBoolean();
	b.GetBoolean() // isOutput1 = r.ReadBoolean();
	b.GetBoolean() // isOutput2 = r.ReadBoolean();
	b.GetBoolean() // isOutput3 = r.ReadBoolean();
	b.GetInt32le() // fluidStorageCount = r.ReadInt32();
	b.GetInt32le() // currentCount = r.ReadInt32();
	b.GetInt32le() // fluidId = r.ReadInt32();
	b.GetBoolean() // outputSwitch = r.ReadBoolean();
	b.GetBoolean() // inputSwitch = r.ReadBoolean();
	b.GetBoolean() // isBottom = r.ReadBoolean();
}
