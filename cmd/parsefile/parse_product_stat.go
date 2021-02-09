package main

import "fmt"

func parseProductStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown ProductStat version: %d\n", vers))
	}

	// count array
	count := b.GetInt32le()
	fmt.Printf("Production count (count array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d %d\n", i, c)
	}

	// cursor array
	count = b.GetInt32le()
	fmt.Printf("Production count (cursor array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d %d\n", i, c)
	}

	// total array
	count = b.GetInt32le()
	fmt.Printf("Production count (total array): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d %d\n", i, c)
	}

	itemID := b.GetInt32le()
	fmt.Printf("Production itemId: %d\n", itemID)
}
