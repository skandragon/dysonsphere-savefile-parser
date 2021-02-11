package main

func parseCargoContainer(b *Buffer) {
	checkVers(b, 0, "CargoContainer")

	poolCapacity := b.GetInt32le()
	cursor := b.GetInt32le()
	b.GetInt32le() // recycleBegin
	b.GetInt32le() // recycleEnd
	for i := 0; int32(i) < cursor; i++ {
		b.GetInt32le() // cargoPool.item
		b.GetFloat32() // cargoPool.position.x
		b.GetFloat32() // cargoPool.position.y
		b.GetFloat32() // cargoPool.position.z
		b.GetFloat32() // cargoPool.rotation.x
		b.GetFloat32() // cargoPool.rotation.y
		b.GetFloat32() // cargoPool.rotation.z
		b.GetFloat32() // cargoPool.rotation.w
	}
	for i := 0; int32(i) < poolCapacity; i++ {
		b.GetInt32le() // recycleID
	}
}
