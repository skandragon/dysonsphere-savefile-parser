package main

type VeinType int32

const (
	VeinTypeNone VeinType = 0
	VeinTypeIron
	VeinTypeCopper
	VeinTypeSilicium
	VeinTypeTitanium
	VeinTypeStone
	VeinTypeCoal
	VeinTypeOil
	VeinTypeFireice
	VeinTypeDiamond
	VeinTypeFractal
	VeinTypeCrysrub
	VeinTypeGrat
	VeinTypeBamboo
	VeinTypeMag
	VeinTypeMax
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
		"Diamond",
		"Fractal",
		"Crysrub",
		"Grat",
		"Bamboo",
		"Mag",
		"Max",
	}[t]
}
