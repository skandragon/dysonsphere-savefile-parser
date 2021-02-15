package cs

// Lerp interpolates between a and b by t.  t is clamped between 0 and 1
func Lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*Clamp01(t)
}

// Clamp returns min if v < min, max if v > max, else v.
func Clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Clamp01 is the same as Clamp(v, 0, 1)
func Clamp01(v float64) float64 {
	return Clamp(v, 0, 1)
}
