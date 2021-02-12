package main

type PlanetType int32

const (
	PlanetTypeNone    PlanetType = 0
	PlanetTypeVolcano PlanetType = 1
	PlanetTypeOcean   PlanetType = 2
	PlanetTypeDesert  PlanetType = 3
	PlanetTypeIce     PlanetType = 4
	PlanetTypeGas     PlanetType = 5
)

func (t PlanetType) String() string {
	return [...]string{
		"None",
		"Volcano",
		"Ocean",
		"Desert",
		"Ice",
		"Gas",
	}[t]
}
