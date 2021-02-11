package main

func parseProductionStatistics(b *Buffer) {
	checkVers(b, 0, "ProductionStatistics")

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		parseFactoryProductionStat(b)
	}

	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // id
	}

	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // id
	}
}
