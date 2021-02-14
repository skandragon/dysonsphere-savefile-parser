package main

// StarType describes the type of star.
type StarType int32

const (
	// StarTypeMainSeqStar is the general 'normal' star.
	StarTypeMainSeqStar StarType = 0

	// StarTypeGiantStar is for gas or ice giants.
	StarTypeGiantStar StarType = 1

	// StarTypeWhiteDwarf ar for the little ones that barely shine.
	StarTypeWhiteDwarf StarType = 2

	// StarTypeNeutronStar are, well, neutron stars.
	StarTypeNeutronStar StarType = 3

	// StarTypeBlackHole really suck.
	StarTypeBlackHole StarType = 4
)

func (t StarType) String() string {
	return [...]string{"Main Sequence", "Giant", "White Dwarf", "Neutron", "Black Hole"}[t]
}
