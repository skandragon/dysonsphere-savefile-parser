package main

import "fmt"

func parseProductionStatistics(b *Buffer) {
	checkVers(b, 0, "ProductionStatistics")

	count := b.GetInt32le()
	fmt.Printf("factory production stat count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		parseFactoryProductionStat(b)
	}

	count = b.GetInt32le()
	fmt.Printf("firstCreateIds count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		if id != 0 {
			fmt.Printf("   id: %d\n", id)
		}
	}

	count = b.GetInt32le()
	fmt.Printf("favoriteIds count: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		id := b.GetInt32le()
		if id != 0 {
			fmt.Printf("   id: %d\n", id)
		}
	}
}
