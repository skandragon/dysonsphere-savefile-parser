package main

import (
	"math"

	"github.com/skandragon/dysonsphere/internal/cs"
	"github.com/skandragon/dysonsphere/internal/cs/mathf"
)

// StarData holds all the statistics for a single star.
type StarData struct {
	Index              int32
	Level              float32
	ID                 int32
	Seed               int32
	ResourceCoef       float32
	Name               string
	Position           *VectorLF3
	UPosition          *VectorLF3
	Mass               float32
	Age                float32
	Lifetime           float32
	Temperature        float32
	Type               StarType
	Spectr             SpectralType
	Color              float32
	ClassFactor        float32
	Luminosity         float32
	Radius             float32
	AcdiskRadius       float32
	HabitableRadius    float32
	LightBalanceRadius float32
	OrbitScaler        float32
	DysonRadius        float32
	PhysicsRadius      float32
}

// Make the special starting star.
func makeBirthStar(seed int32, universe UniqueStarnameChecker) *StarData {
	star := &StarData{
		Index:        0,
		Level:        0,
		ID:           1,
		Seed:         seed,
		ResourceCoef: 0.6,
	}
	random := cs.MakePRNGSequence(seed)
	seed2 := random.Next()
	seed3 := random.Next()
	//	star.Name = randomName(seed2)
	random2 := cs.MakePRNGSequence(seed3)
	r := random2.NextDouble()
	r2 := random2.NextDouble()
	num := random2.NextDouble()
	rn := random2.NextDouble()
	rt := random2.NextDouble()
	num2 := random2.NextDouble()*0.2 + 0.9
	y := random2.NextDouble()*0.4 - 0.2
	num3 := math.Pow(2.0, y)
	num4 := randNormal(0, 0.08, r, r2)
	num4 = mathf.Clamp(num4, -0.2, 0.2)
	star.Mass = mathf.Pow(2, num4)
	d := 2.0 + 0.4*(1.0-float64(star.Mass))
	star.Lifetime = float32(10000 * math.Pow(0.1, math.Log10(float64(star.Mass)*0.5)/math.Log10(d)+1.0) * num2)
	star.Age = float32(num*0.4 + 0.3)

	num5 := (1 - mathf.Pow(mathf.Clamp01(star.Age), 20)*0.5) * star.Mass
	star.Temperature = float32(math.Pow(float64(num5), 0.56+0.14/(math.Log10(float64(num5)+4)/math.Log10(5.0)))*4450.0 + 1300.0)
	num6 := math.Log10((float64(star.Temperature)-1300.0)/4500.0)/math.Log10(2.6) - 0.5
	if num6 < 0.0 {
		num6 *= 4.0
	}
	if num6 > 2.0 {
		num6 = 2.0
	} else if num6 < -4.0 {
		num6 = -4.0
	}
	star.Spectr = SpectralType(mathf.RoundToInt(float32(num6) + 4))
	star.Color = mathf.Clamp01((float32(num6) + 3.5) * 0.2)
	star.ClassFactor = float32(num6)
	star.Luminosity = mathf.Pow(num5, 0.7)
	star.Radius = mathf.Pow(star.Mass, 0.4) * float32(num3)
	star.AcdiskRadius = 0
	p := float32(num6) + 2
	star.HabitableRadius = mathf.Pow(1.7, p) + 0.2*mathf.Min(1, star.OrbitScaler)
	star.LightBalanceRadius = mathf.Pow(1.7, p)
	star.OrbitScaler = mathf.Pow(1.35, p)
	if star.OrbitScaler < 1 {
		star.OrbitScaler = mathf.Lerp(star.OrbitScaler, 1, 0.6)
	}
	setStarAge(star, star.Age, rn, rt)
	star.DysonRadius = star.OrbitScaler * 0.28
	if star.DysonRadius*40000.0 < (star.PhysicsRadius * 1.5) {
		star.DysonRadius = ((star.PhysicsRadius * 1.5) / 40000.0)
	}
	star.Name = randomStarName(seed2, star, universe)
	return star
}

func setStarAge(star *StarData, age float32, rn float64, rt float64) {
	num := float32(rn*0.1 + 0.95)
	num2 := float32(rt*0.4 + 0.8)
	num3 := float32(rt*9.0 + 1.0)
	star.Age = age
	if age >= 1 {
		if star.Mass >= 18 {
			star.Type = StarTypeBlackHole
			star.Spectr = SpectralTypeX
			star.Mass *= 2.5 * num2
			star.Radius *= 1
			star.AcdiskRadius = star.Radius * 5
			star.Temperature = 0
			star.Luminosity *= 0.001 * num
			star.HabitableRadius = 0
			star.LightBalanceRadius *= 0.4 * num
		} else if star.Mass >= 7 {
			star.Type = StarTypeNeutronStar
			star.Spectr = SpectralTypeX
			star.Mass *= 0.2 * num
			star.Radius *= 0.15
			star.AcdiskRadius = star.Radius * 9
			star.Temperature = num3 * 10000000
			star.Luminosity *= 0.1 * num
			star.HabitableRadius = 0
			star.LightBalanceRadius *= 3 * num
			star.OrbitScaler *= 1.5 * num
		} else {
			star.Type = StarTypeWhiteDwarf
			star.Spectr = SpectralTypeX
			star.Mass *= 0.2 * num
			star.Radius *= 0.2
			star.AcdiskRadius = 0
			star.Temperature = num2 * 150000
			star.Luminosity *= 0.04 * num2
			star.HabitableRadius *= 0.15 * num2
			star.LightBalanceRadius *= 0.2 * num
		}
	} else if age >= 0.96 {
		num4 := mathf.Pow(5.0, mathf.Abs(mathf.Log10(star.Mass)-0.7)) * 5.0
		if num4 > 10 {
			num4 = float32((mathf.Log(num4*0.1) + 1) * 10)
		}
		num5 := float32(1 - mathf.Pow(star.Age, 30)*0.5)
		star.Type = StarTypeGiantStar
		star.Mass = num5 * star.Mass
		star.Radius = num4 * num2
		star.AcdiskRadius = 0
		star.Temperature = num5 * star.Temperature
		star.Luminosity = 1.6 * star.Luminosity
		star.HabitableRadius = 9 * star.HabitableRadius
		star.LightBalanceRadius = 3 * star.HabitableRadius
		star.OrbitScaler = 3.3 * star.OrbitScaler
	}
}

func createStar(galaxy *Universe, pos *VectorLF3, id int32, seed int32, needType StarType, needSpectr SpectralType) *StarData {
	star := &StarData{
		Index:    id - 1,
		Seed:     seed,
		ID:       id,
		Position: pos,
	}
	if galaxy.StarCount > 1 {
		star.Level = float32(star.Index) / float32(galaxy.StarCount-1)
	}
	random := cs.MakePRNGSequence(seed)
	seed2 := random.Next()
	seed3 := random.Next()
	num := float32(pos.Magnitude())
	num2 := num / 32
	if num2 > 1 {
		num2 = mathf.Log(num2) + 1
		num2 = mathf.Log(num2) + 1
		num2 = mathf.Log(num2) + 1
		num2 = mathf.Log(num2) + 1
		num2 = mathf.Log(num2) + 1
	}
	star.ResourceCoef = mathf.Pow(7, num2) * 0.6
	random2 := cs.MakePRNGSequence(seed3)
	num3 := random2.NextDouble()
	num4 := random2.NextDouble()
	num5 := random2.NextDouble()
	rn := random2.NextDouble()
	rt := random2.NextDouble()
	num6 := (random2.NextDouble() - 0.5) * 0.2
	num7 := random2.NextDouble()*0.2 + 0.9
	num8 := random2.NextDouble()*0.4 - 0.2
	num9 := math.Pow(2.0, num8)
	num10 := mathf.Lerp(-0.98, 0.88, star.Level)
	if num10 < 0 {
		num10 -= 0.65
	} else {
		num10 += 0.65
	}
	standardDeviation := float32(0.33)

	if needType == StarTypeGiantStar {
		if num8 <= -0.08 {
			num10 = 1.6
		} else {
			num10 = -1.5
		}
		standardDeviation = 0.3
	}
	num11 := randNormal(num10, standardDeviation, num3, num4)
	if needSpectr == SpectralTypeM {
		num11 = -3
	} else if needSpectr == SpectralTypeO {
		num11 = 3
	}
	if num11 > 0 {
		num11 *= 2
	}
	num11 = mathf.Clamp(num11, -2.4, 4.65) + float32(num6) + 1
	if needType == StarTypeBlackHole {
		star.Mass = 18 + float32(num3*num4)*30
	} else if needType == StarTypeNeutronStar {
		star.Mass = 7 + float32(num3)*11
	} else if needType == StarTypeWhiteDwarf {
		star.Mass = 1 + float32(num4)*5
	} else {
		star.Mass = mathf.Pow(2, num11)
	}
	d := 5.0
	if star.Mass < 2 {
		d = 2 + 0.4*(1-float64(star.Mass))
	}
	star.Lifetime = float32(10000.0 * math.Pow(0.1, math.Log10(float64(star.Mass)*0.5)/math.Log10(d)+1) * num7)
	if needType == StarTypeGiantStar {
		star.Lifetime = float32(10000.0 * math.Pow(0.1, math.Log10(float64(star.Mass)*0.58)/math.Log10(d)+1) * num7)
		star.Age = float32(num5)*0.04 + 0.96
	} else if needType == StarTypeBlackHole || needType == StarTypeNeutronStar || needType == StarTypeWhiteDwarf {
		star.Age = float32(num5)*0.4 + 1
		if needType == StarTypeWhiteDwarf {
			star.Lifetime += 10000
		} else if needType == StarTypeNeutronStar {
			star.Lifetime += 1000
		}
	} else if star.Mass < 0.5 {
		star.Age = float32(num5)*0.12 + 0.02
	} else if star.Mass < 0.8 {
		star.Age = float32(num5)*0.4 + 0.1
	} else {
		star.Age = float32(num5)*0.7 + 0.2
	}
	num12 := star.Lifetime * star.Age
	if num12 > 5000 {
		num12 = (mathf.Log(num12/5000) + 1) * 5000
	}
	if num12 > 8000 {
		num13 := num12 / 8000
		num13 = mathf.Log(num13) + 1
		num13 = mathf.Log(num13) + 1
		num13 = mathf.Log(num13) + 1
		num12 = num13 * 8000
	}
	star.Lifetime = num12 / star.Age
	num14 := (1 - mathf.Pow(mathf.Clamp01(star.Age), 20)*0.5) * star.Mass
	star.Temperature = float32(math.Pow(float64(num14), 0.56+0.14/(math.Log10(float64(num14+4))/math.Log10(5.0)))*44.50 + 1300)
	num15 := math.Log10((float64(star.Temperature)-1300)/4500)/math.Log10(2.6) - 0.5
	if num15 < 0 {
		num15 *= 4
	}
	if num15 > 2 {
		num15 = 2
	} else if num15 < -4 {
		num15 = -4
	}
	star.Spectr = SpectralType(mathf.RoundToInt(float32(num15) + 4))
	star.Color = mathf.Clamp01((float32(num15) + 3.5) * 0.2)
	star.ClassFactor = float32(num15)
	star.Luminosity = mathf.Pow(num14, 0.7)
	star.Radius = float32(math.Pow(float64(star.Mass), 0.4) * num9)
	star.AcdiskRadius = 0
	p := float32(num15) + 2
	star.HabitableRadius = mathf.Pow(1.7, p) + 0.25*mathf.Min(1, star.OrbitScaler)
	star.LightBalanceRadius = mathf.Pow(1.7, p)
	star.OrbitScaler = mathf.Pow(1.35, p)
	if star.OrbitScaler < 1 {
		star.OrbitScaler = mathf.Lerp(star.OrbitScaler, 1, 0.6)
	}
	setStarAge(star, star.Age, rn, rt)
	star.DysonRadius = star.OrbitScaler * 0.28
	if star.DysonRadius*40000 < star.PhysicsRadius*1.5 {
		star.DysonRadius = star.PhysicsRadius * 1.5 / 40000
	}
	star.UPosition = star.Position.Times(2400000)

	star.Name = randomStarName(seed2, star, galaxy)
	return star
}

func randNormal(averageValue float32, standardDeviation float32, r1 float64, r2 float64) float32 {
	return averageValue + standardDeviation*mathf.Sqrt(-2.0*mathf.Log(float32(1.0-r1)))*mathf.Sin(float32(6.283185307179586*r2))
}

var (
	orbitRadius = []float64{
		0, 0.4, 0.7, 1, 1.4, 1.9, 2.5, 3.3, 4.3, 5.5, 6.9, 8.4, 10,
		11.7, 13.5, 15.4, 17.5,
	}
)
