package main

// VeinSummary is a list of type to total amounts
type VeinSummary struct {
	Type   VeinType `json:"type"`
	Amount int64    `json:"amount"`
}

// VeinGroup is a single value representation of a group of veins
type VeinGroup struct {
	Type     VeinType `json:"type"`
	Position *Vector3 `json:"-"`
	Count    int32    `json:"count"`
	Amount   int64    `json:"amount"`
}

// PlanetDataRuntime is data that is frequently updated.
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
			Type:     VeinType(b.GetInt32le()),
			Position: ParseVector3(b),
			Count:    b.GetInt32le(),
			Amount:   b.GetInt64le(),
		}
	}
	veinGroups = compressVeinGroups(veinGroups)

	return &PlanetDataRuntime{
		VeinSummaries: resourceSummary,
		VeinGroups:    veinGroups,
	}
}

func compressVeinGroups(i []*VeinGroup) []*VeinGroup {
	o := make([]*VeinGroup, 0)
	for _, vg := range i {
		if vg != nil && vg.Amount > 0 {
			o = append(o, vg)
		}
	}
	return o
}
