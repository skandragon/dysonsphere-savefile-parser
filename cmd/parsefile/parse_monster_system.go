package main

func parseMonsterSystem(b *Buffer) {
	checkVers(b, 0, "MonsterSystem")

	b.GetInt32le() // monsterCapacity
	monsterCursor := b.GetInt32le()
	monsterRecycleCursor := b.GetInt32le()
	for i := int32(1); i < monsterCursor; i++ {
		parseMonsterComponent(b)
	}
	for i := int32(0); i < monsterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}
