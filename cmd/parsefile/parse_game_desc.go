package main

type GameDesc struct {
	GalaxyAlgo         int32   `json:"galaxy_algo"`
	GalaxySeed         int32   `json:"galaxy_seed"`
	StarCount          int32   `json:"star_count"`
	PlayerPronto       int32   `json:"player_pronto"`
	ResourceMultiplier float32 `json:"resource_multiplier"`
	ThemeIDs           []int32 `json:"theme_ids"`
}

func parseGameDesc(b *Buffer) *GameDesc {
	checkVers(b, 2, "GameDesc")

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

	gameDesc := &GameDesc{
		GalaxyAlgo:         galaxyAlgo,
		GalaxySeed:         galaxySeed,
		StarCount:          starCount,
		PlayerPronto:       playerPronto,
		ResourceMultiplier: resourceMultiplier,
		ThemeIDs:           tids,
	}
	return gameDesc
}
