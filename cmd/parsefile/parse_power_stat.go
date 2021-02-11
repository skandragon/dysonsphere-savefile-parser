package main

func parsePowerStat(b *Buffer) {
	checkVers(b, 0, "PowerStat")

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt64le()
		//fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt64le()
		//fmt.Printf("  %d\n", c)
	}

}
