package main

// SpectralType holds the spectral class of the star.
type SpectralType int32

const (
	// SpectralTypeM are red, colder stars, < 3500K, and dim.
	SpectralTypeM SpectralType = 0

	// SpectralTypeK are orange to red, 3500 - 5000K, and dimmer than our sun.
	SpectralTypeK SpectralType = 1

	// SpectralTypeG are white to yellow, 5000 - 6000K.
	SpectralTypeG SpectralType = 2

	// SpectralTypeF are blue to white, 6000 - 7500K, 6x brighter than our sun.
	SpectralTypeF SpectralType = 3

	// SpectralTypeA are blue, 7500 - 11000K, 80x brighter than our sun.
	SpectralTypeA SpectralType = 4

	// SpectralTypeB are blue, larger, and 20,000x brighter than our sun.
	SpectralTypeB SpectralType = 5

	// SpectralTypeO are blue giants, over 25000K, and 1,400,000x brighter than our sun.
	SpectralTypeO SpectralType = 6

	// SpectralTypeX are special, in that they are neutron stars or black holes.
	SpectralTypeX SpectralType = 7
)

func (t SpectralType) String() string {
	return [...]string{"M", "K", "G", "F", "A", "B", "O", "X"}[t]
}
