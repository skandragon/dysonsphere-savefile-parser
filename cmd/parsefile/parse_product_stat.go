package main

func parseProductStat(b *Buffer) {
	checkVers(b, 0, "ProductStat")

	// count array
	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d %d\n", i, c)
	}

	// cursor array
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d %d\n", i, c)
	}

	// total array
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		// this.total[k] = ((num8 >= 0) ? num8 : (-num8));
		//fmt.Printf("  %d %d\n", i, c)
	}

	b.GetInt32le() // itemID
}
