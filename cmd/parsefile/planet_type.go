package main

type PlanetType int32

const (
	PlanetTypeNone PlanetType = 0
	PlanetTypeVolcano
	PlanetTypeOcean
	PlanetTypeDesert
	PlanetTypeIce
	PlanetTypeGas
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
