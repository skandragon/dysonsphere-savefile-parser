package main

import "fmt"

func parseFactoryProductionStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 1 {
		panic(fmt.Sprintf("Unknown FactoryProductionStat version: %d\n", vers))
	}

	productCapacity := b.GetInt32le()
	productCursor := b.GetInt32le()
	fmt.Printf("productCursor (count) %d, productCapacity %d\n", productCursor, productCapacity)
	for i := 0; int32(i) < productCursor; i++ {
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
		c := b.GetInt32le()
		fmt.Printf("  %d %d", i, c)
	}

	powerProduction := b.GetInt64le()
	fmt.Printf("Power production: %d\n", powerProduction)
}
