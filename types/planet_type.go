package types

// PlanetType indicates the type of the planet's surface.
type PlanetType int32

const (
	// PlanetTypeNone means it doesn't exist I suppose?
	PlanetTypeNone PlanetType = 0

	// PlanetTypeVolcano are hot and likely rocky.
	PlanetTypeVolcano PlanetType = 1

	// PlanetTypeOcean means you'll need lots of landfill to mine.
	PlanetTypeOcean PlanetType = 2

	// PlanetTypeDesert are dry most likely.
	PlanetTypeDesert PlanetType = 3

	// PlanetTypeIce are ice giants.
	PlanetTypeIce PlanetType = 4

	// PlanetTypeGas are gas giants.
	PlanetTypeGas PlanetType = 5
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
