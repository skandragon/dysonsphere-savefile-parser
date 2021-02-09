package main

import "fmt"

func parseGameStatistics(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown GameStatistics version: %d", vers))
	}

	count := b.GetInt32le()
	fmt.Printf("Statistics: count of hashedTechHistory: %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		// fmt.Printf("   %d\n", id)
	}

	parseProductionStatistics(b)
}
