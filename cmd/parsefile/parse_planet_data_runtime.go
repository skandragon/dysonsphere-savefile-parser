package main

type VeinSummary struct {
	Type   VeinType `json:"type"`
	Amount int64    `json:"amount"`
}

type VeinGroup struct {
	Type     VeinType        `json:"type"`
	Position *PositionSingle `json:"-"`
	Count    int32           `json:"count"`
	Amount   int64           `json:"amount"`
}

type PlanetDataRuntime struct {
	VeinSummaries []*VeinSummary `json:"vein_summaries"`
	VeinGroups    []*VeinGroup   `json:"vein_groups"`
}

func parsePlanetDataRuntime(b *Buffer) *PlanetDataRuntime {
	count := b.GetInt32le()
	b.GetBytes(int(count)) // modData

	veinSummaryCount := b.GetInt32le()
	resourceSummary := make([]*VeinSummary, 0)
	for i := 0; int32(i) < veinSummaryCount; i++ {
		amount := b.GetInt64le()
		if amount > 0 {
			resourceSummary = append(resourceSummary, &VeinSummary{
				Type:   VeinType(i),
				Amount: amount,
			})
		}
	}

	veinGroupCount := b.GetInt32le()
	veinGroups := make([]*VeinGroup, veinGroupCount)
	for i := 0; int32(i) < veinGroupCount; i++ {
		veinGroups[i] = &VeinGroup{
			Type: VeinType(b.GetInt32le()),
			Position: &PositionSingle{
				X: b.GetFloat32(),
				Y: b.GetFloat32(),
				Z: b.GetFloat32(),
			},
			Count:  b.GetInt32le(),
			Amount: b.GetInt64le(),
		}
	}

	return &PlanetDataRuntime{
		VeinSummaries: resourceSummary,
		VeinGroups:    veinGroups,
	}
}
