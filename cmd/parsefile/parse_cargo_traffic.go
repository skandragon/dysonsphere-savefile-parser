package main

func parseCargoTraffic(b *Buffer) {
	checkVers(b, 0, "CargoTraffic")

	beltCursor := b.GetInt32le()
	b.GetInt32le() // beltCapacity
	beltRecycleCursor := b.GetInt32le()

	splitterCursor := b.GetInt32le()
	b.GetInt32le() // splitterCapacity
	splitterRecycleCursor := b.GetInt32le()

	pathCursor := b.GetInt32le()
	b.GetInt32le() // pathCapacity
	pathRecycleCursor := b.GetInt32le()

	for i := 1; int32(i) < beltCursor; i++ {
		parseBeltComponent(b)
	}
	for i := 0; int32(i) < beltRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < splitterCursor; i++ {
		parseSplitterComponent(b)
	}
	for i := 0; int32(i) < splitterRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}

	for i := 1; int32(i) < pathCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			parseCargoPath(b)
		}
	}
	for i := 0; int32(i) < pathRecycleCursor; i++ {
		b.GetInt32le() // recycle id?
	}
}
