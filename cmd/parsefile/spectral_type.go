package main

// SpectralType holds the spectral class of the star.
type SpectralType int32

const (
	// SpectralTypeM are type M.
	SpectralTypeM SpectralType = 0

	// SpectralTypeK are type K.
	SpectralTypeK SpectralType = 1

	// SpectralTypeG are type G.
	SpectralTypeG SpectralType = 2

	// SpectralTypeF are type F.
	SpectralTypeF SpectralType = 3

	// SpectralTypeA are type A.
	SpectralTypeA SpectralType = 4

	// SpectralTypeB are type B.
	SpectralTypeB SpectralType = 5

	// SpectralTypeO are type O.
	SpectralTypeO SpectralType = 6

	// SpectralTypeX are type X.
	SpectralTypeX SpectralType = 7
)

func (t SpectralType) String() string {
	return [...]string{"M", "K", "G", "F", "A", "B", "O", "X"}[t]
}
