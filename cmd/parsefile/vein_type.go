package main

// VeinType describes the type of material that can be extracted.
type VeinType int32

const (
	// VeinTypeNone means the vein is toast, I suspect.
	VeinTypeNone VeinType = 0

	// VeinTypeIron - Iron Ore
	VeinTypeIron VeinType = 1

	// VeinTypeCopper - Copper Ore
	VeinTypeCopper VeinType = 2

	// VeinTypeSilicium - actually Silicon Ore
	VeinTypeSilicium VeinType = 3

	// VeinTypeTitanium - Titanium Ore
	VeinTypeTitanium VeinType = 4

	// VeinTypeStone - stone
	VeinTypeStone VeinType = 5

	// VeinTypeCoal - coal
	VeinTypeCoal VeinType = 6

	// VeinTypeOil - oil
	VeinTypeOil VeinType = 7

	// VeinTypeFireice - Fire Ice
	VeinTypeFireice VeinType = 8

	// VeinTypeDiamond - Kimberlite Ore
	VeinTypeDiamond VeinType = 9

	// VeinTypeFractal - Fractal Silicon
	VeinTypeFractal VeinType = 10

	// VeinTypeCrysrub - Organic Crystal
	VeinTypeCrysrub VeinType = 11

	// VeinTypeGrat - Optical Grating Crystal
	VeinTypeGrat VeinType = 12

	// VeinTypeBamboo - Spiniform Stalagmite Crystal
	VeinTypeBamboo VeinType = 13

	// VeinTypeMag -- Unipolar Magnet
	VeinTypeMag VeinType = 14

	// VeinTypeMax - the end of the types
	VeinTypeMax VeinType = 15
)

func (t VeinType) String() string {
	return [...]string{
		"None",
		"Iron",
		"Copper",
		"Silicon",
		"Titanium",
		"Stone",
		"Coal",
		"Oil",
		"Fire Ice",
		"Kimberlite Ore",
		"Fractal Silicon",
		"Fire Ice",
		"Optical Grating Crystal",
		"Spiniform Stalagmite Crystal",
		"Unipolar Magnet",
		"Max",
	}[t]
}
