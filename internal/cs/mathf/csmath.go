package mathf

import "math"

// Lerp interpolates between a and b by t.  t is clamped between 0 and 1
func Lerp(a float32, b float32, t float32) float32 {
	return a + (b-a)*Clamp01(t)
}

// Clamp returns min if v < min, max if v > max, else v.
func Clamp(v float32, min float32, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Clamp01 is the same as Clamp(v, 0, 1)
func Clamp01(v float32) float32 {
	return Clamp(v, 0, 1)
}

// CeilToInt - return int32(math.Ceil(v))
func CeilToInt(v float32) int32 {
	return int32(math.Ceil(float64(v)))
}

func Pow(x float32, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func RoundToInt(x float32) int32 {
	return int32(math.Round(float64(x)))
}

func Min(x float32, y float32) float32 {
	if x < y {
		return x
	}
	return y
}
func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func Log(x float32) float32 {
	return float32(math.Log(float64(x)))
}

func Log10(x float32) float32 {
	return float32(math.Log10(float64(x)))
}

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}
