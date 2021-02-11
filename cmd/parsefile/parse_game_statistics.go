package main

func parseGameStatistics(b *Buffer) {
	checkVers(b, 0, "GameStatistics")

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		// fmt.Printf("   %d\n", id)
	}

	parseProductionStatistics(b)
}
