package main

import "fmt"

func parsePowerStat(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 0 {
		panic(fmt.Sprintf("Unknown PowerStat version: %d", vers))
	}

	count := b.GetInt32le()
	fmt.Printf("Power energy (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt64le()
		fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power cursor (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt32le()
		fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power total (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		c := b.GetInt64le()
		fmt.Printf("  %d\n", c)
	}

}
