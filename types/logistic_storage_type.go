package types

// LogisticStorageType holds the request/supply status.
type LogisticStorageType int32

const (
	// LogisticStorageTypeNone indicates this item is stored
	LogisticStorageTypeNone LogisticStorageType = 0

	// LogisticStorageTypeSupply indicates this item is supplied
	LogisticStorageTypeSupply LogisticStorageType = 1

	// LogisticStorageTypeDemand indicates this item is requested
	LogisticStorageTypeDemand LogisticStorageType = 2
)

func (t LogisticStorageType) String() string {
	return [...]string{"None", "Supply", "Demand"}[t]
}
