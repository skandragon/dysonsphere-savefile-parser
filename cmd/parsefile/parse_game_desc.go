package main

import "fmt"

type gameDesc struct {
	galaxyAlgo         int32
	galaxySeed         int32
	starCount          int32
	playerPronto       int32
	resourceMultiplier float32
	themeIds           []int32
}

func (g *gameDesc) Dump() {
	fmt.Printf("galaxy algorithm: %d\n", g.galaxyAlgo)
	fmt.Printf("galaxy seed: %d\n", g.galaxySeed)
	fmt.Printf("star count: %d\n", g.starCount)
	fmt.Printf("playerPronto: %d\n", g.playerPronto)
	fmt.Printf("resource multiplier: %f\n", g.resourceMultiplier)
	for _, tid := range g.themeIds {
		fmt.Printf("  theme ID: %d\n", tid)
	}
}

func parseGameDesc(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameDesc version: %d", vers))
	}

	galaxyAlgo := b.GetInt32le()
	galaxySeed := b.GetInt32le()
	starCount := b.GetInt32le()
	playerPronto := b.GetInt32le()
	resourceMultiplier := b.GetFloat32()

	tidlen := b.GetInt32le()
	tids := make([]int32, int(tidlen))
	for i := 0; int32(i) < tidlen; i++ {
		tids[i] = b.GetInt32le()
	}

	gameDesc := &gameDesc{
		galaxyAlgo:         galaxyAlgo,
		galaxySeed:         galaxySeed,
		starCount:          starCount,
		playerPronto:       playerPronto,
		resourceMultiplier: resourceMultiplier,
		themeIds:           tids,
	}
	gameDesc.Dump()
}
