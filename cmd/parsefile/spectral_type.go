package main

type SpectralType int32

const (
	SpectralTypeM SpectralType = 0
	SpectralTypeK SpectralType = 1
	SpectralTypeG SpectralType = 2
	SpectralTypeF SpectralType = 3
	SpectralTypeA SpectralType = 4
	SpectralTypeB SpectralType = 5
	SpectralTypeO SpectralType = 6
	SpectralTypeX SpectralType = 7
)

func (t SpectralType) String() string {
	return [...]string{"M", "K", "G", "F", "A", "B", "O", "X"}[t]
}
