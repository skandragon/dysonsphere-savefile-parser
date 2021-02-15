package types

//MinerType indicates what type of thing a miner is.
type MinerType int32

const (
	// MinerTypeNone is used to indicate no miner is in a slot.
	MinerTypeNone MinerType = 0

	// MinerTypeWater can extrct water or acid from oceans.
	MinerTypeWater MinerType = 1

	// MinerTypeVein extracts materials from ore and similar veins.
	MinerTypeVein MinerType = 2

	// MinerTypeOil extracts oil from crude oil seeps.
	MinerTypeOil MinerType = 3
)

func (t MinerType) String() string {
	return [...]string{
		"None",
		"Water",
		"Vein",
		"Oil",
	}[t]
}
