package main

type MinerType int32

const (
	MinerTypeNone  MinerType = 0
	MinerTypeWater MinerType = 1
	MinerTypeVein  MinerType = 2
	MinerTypeOil   MinerType = 3
)

func (t MinerType) String() string {
	return [...]string{
		"None",
		"Water",
		"Vein",
		"Oil",
	}[t]
}
