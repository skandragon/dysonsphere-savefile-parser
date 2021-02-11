package main

func parseFactoryProductionStat(b *Buffer) {
	checkVers(b, 1, "FactoryProductionStat")

	b.GetInt32le() // productCapacity
	productCursor := b.GetInt32le()
	for i := 1; int32(i) < productCursor; i++ {
		parseProductStat(b)
	}

	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		parsePowerStat(b)
	}
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
	}

	b.GetInt64le() // energyConsumption
}
