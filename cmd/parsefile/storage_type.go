package main

type StorageType int32

const (
	StorageTypeDefault  StorageType = 0
	StorageTypeFuel     StorageType = 1
	StorageTypeFiltered StorageType = 9
)

func (t StorageType) String() string {
	return [...]string{
		"Default",
		"Fuel",
		"Unknown2",
		"Unknown3",
		"Unknown4",
		"Unknown5",
		"Unknown6",
		"Unknown7",
		"Unknown8",
		"Filtered",
	}[t]
}
