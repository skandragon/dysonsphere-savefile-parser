package main

import "fmt"

func parsePowerStat(b *Buffer) {
	checkVers(b, 0, "PowerStat")

	count := b.GetInt32le()
	fmt.Printf("Power energy (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt64le()
		//fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power cursor (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le()
		//fmt.Printf("  %d\n", c)
	}

	count = b.GetInt32le()
	fmt.Printf("Power total (count): %d\n", count)
	for i := 0; int32(i) < count; i++ {
		b.GetInt64le()
		//fmt.Printf("  %d\n", c)
	}

}
