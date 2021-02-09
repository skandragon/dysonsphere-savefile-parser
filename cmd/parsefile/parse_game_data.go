package main

import "fmt"

func parseGameData(b *Buffer) {
	vers := b.GetInt32le()
	if vers != 2 {
		panic(fmt.Sprintf("Unknown GameData version: %d", vers))
	}

	gamename := b.GetString()
	fmt.Printf("Save file name: %s\n", gamename)

	parseGameDesc(b)
	gameTicks := b.GetInt64le()
	fmt.Printf("Game ticks: %d\n", gameTicks)

	parseGamePreferences(b)
	parseGameHistory(b)

	b.GetBoolean() // hidePlayerModel
	b.GetBoolean() // disableController

	parseGameStatistics(b)

	b.GetInt32le() // planetId

	parsePlayer(b)

	factoryCount := b.GetInt32le()

	parseGalacticTransport(b) // yes, this is out of order in the code too...

	for i := 0; int32(i) < factoryCount; i++ {
		parsePlanetFactory(b, i)
	}

	starcount := b.GetInt32le()
	fmt.Printf("Star count (again?) %d\n", starcount)

	// dyson sphere mess...
}
