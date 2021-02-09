package main

import "fmt"

func parseFactoryProductionStat(b *Buffer) {
	checkVers(b, 1, "FactoryProductionStat")

	productCapacity := b.GetInt32le()
	productCursor := b.GetInt32le()
	fmt.Printf("productCursor (count) %d, productCapacity %d\n", productCursor, productCapacity)
	for i := 1; int32(i) < productCursor; i++ {
		parseProductStat(b)
	}

	count := b.GetInt32le()
	fmt.Printf("powerPool count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		parsePowerStat(b)
	}
	count = b.GetInt32le()
	fmt.Printf("productionIndicies (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d %d\n", i, c)
	}

	energyConsumption := b.GetInt64le()
	fmt.Printf("energyConsumption: %d\n", energyConsumption)
}
