package types

// StorageType indicates the type of storage slot, I suppose?
type StorageType int32

const (
	// StorageTypeDefault seems to be a 'generic' boring slot.
	StorageTypeDefault StorageType = 0

	// StorageTypeFuel seems to be something that requires a fuel item.
	StorageTypeFuel StorageType = 1

	// StorageTypeFiltered seems to be ... filtered?
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
