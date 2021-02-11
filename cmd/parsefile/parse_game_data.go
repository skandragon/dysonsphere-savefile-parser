package main

type GameData struct {
	Name            string           `json:"name"`
	Descriptor      *GameDesc        `json:"descriptor"`
	StarCount       int32            `json:"star_count"`
	PlanetFactories []*PlanetFactory `json:"planet_factories"`
	GameTicks       int64            `json:"game_ticks"`
}

func parseGameData(b *Buffer) *GameData {
	const VERSION = 2
	checkVers(b, VERSION, "GameData")

	ret := &GameData{
		Name:       b.GetString(),
		Descriptor: parseGameDesc(b),
	}

	ret.GameTicks = b.GetInt64le()

	parseGamePreferences(b)
	parseGameHistory(b)

	b.GetBoolean() // hidePlayerModel
	b.GetBoolean() // disableController

	parseGameStatistics(b)

	b.GetInt32le() // planetId

	parsePlayer(b)

	factoryCount := b.GetInt32le()

	parseGalacticTransport(b) // yes, this is out of order in the code too...

	ret.PlanetFactories = make([]*PlanetFactory, int(factoryCount))
	for i := 0; int32(i) < factoryCount; i++ {
		ret.PlanetFactories[i] = parsePlanetFactory(b, i)
	}

	ret.StarCount = b.GetInt32le()

	// dyson sphere mess...

	return ret
}
