package main

type VeinType int32

const (
	VeinTypeNone     VeinType = 0
	VeinTypeIron     VeinType = 1
	VeinTypeCopper   VeinType = 2
	VeinTypeSilicium VeinType = 3
	VeinTypeTitanium VeinType = 4
	VeinTypeStone    VeinType = 5
	VeinTypeCoal     VeinType = 6
	VeinTypeOil      VeinType = 7
	VeinTypeFireice  VeinType = 8
	VeinTypeDiamond  VeinType = 9
	VeinTypeFractal  VeinType = 10
	VeinTypeCrysrub  VeinType = 11
	VeinTypeGrat     VeinType = 12
	VeinTypeBamboo   VeinType = 13
	VeinTypeMag      VeinType = 14
	VeinTypeMax      VeinType = 15
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
