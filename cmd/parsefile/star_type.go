package main

type StarType int32

const (
	StarTypeMainSeqStar StarType = 0
	StarTypeGiantStar   StarType = 1
	StarTypeWhiteDwarf  StarType = 2
	StarTypeNeutronStar StarType = 3
	StarTypeBlackHole   StarType = 4
)

func (t StarType) String() string {
	return [...]string{"Main Sequence", "Giant", "White Dwarf", "Neutron", "Black Hole"}[t]
}
