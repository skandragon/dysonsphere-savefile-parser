package main

import (
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/skandragon/dysonsphere/internal/cs"
)

type Universe struct {
	Stars []*StarData `json:"stars"`
	Seed  int32       `json:"seed"`
}

func MakeUniverse(gd *GameDesc) *Universe {
	galaxySeed := gd.GalaxySeed
	random := cs.MakePRNGSequence(galaxySeed)
	random.Next() // seed

	//num := generateTempPoses(seed, gd.StarCount, 4, 2.0, 2.3, 3.5, 0.18)

	random.NextDouble() // num2
	random.NextDouble() // num3
	random.NextDouble() // num4
	random.NextDouble() // num5

	stars := make([]*StarData, 0)
	for i := int32(0); i < gd.StarCount; i++ {
		seed2 := random.Next()
		if i == 0 {
			stars = append(stars, makeBirthStar(seed2))
		} else {
			//stars = append(stars, makeBirthStar(seed2))
		}
		//makeStar(seed2)
	}

	return &Universe{
		Stars: stars,
	}
}

// hacking, trying to find out how random seeding actually works...
func runSearch(galaxySeed int32) {
	var wg sync.WaitGroup
	cpuCount := runtime.NumCPU()
	wg.Add(cpuCount)
	index := int32(-1)
	for job := 0; job < cpuCount; job++ {
		go func() {
			jobIndex := atomic.AddInt32(&index, 1)
			slice := int32(math.MaxInt32 / cpuCount)
			startVal := slice * jobIndex
			endVal := startVal + slice
			log.Printf("core %d, start %d, end %d", jobIndex, startVal, endVal)
			search(galaxySeed, startVal, endVal)
			wg.Done()
		}()
	}
	wg.Wait()
	os.Exit(0)

}

func search(seed int32, startVal int32, endVal int32) {
	seeds := make([]int32, 0)

	tstar := &StarData{Type: StarTypeMainSeqStar}
	for i := startVal; i < endVal; i++ {
		rtest := cs.MakePRNGSequence(i)
		//rtest.Next()
		//rtest.NextDouble()
		//rtest.NextDouble()
		//rtest.NextDouble()
		rtest.NextDouble()
		tseed := rtest.Next()
		if randomStarName(tseed, tstar, []string{}) == "Bunda" {
			seeds = append(seeds, i)
		}
	}
	log.Printf("Found: %v", seeds)
}

func randNormal(averageValue float64, standardDeviation float64, r1 float64, r2 float64) float64 {
	return averageValue + standardDeviation*math.Sqrt(-2.0*math.Log(1.0-r1))*math.Sin(6.283185307179586*r2)
}

func clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func clamp01(v float64) float64 {
	return clamp(v, 0, 1)
}

var (
	orbitRadius = []float64{
		0, 0.4, 0.7, 1, 1.4, 1.9, 2.5, 3.3, 4.3, 5.5, 6.9, 8.4, 10,
		11.7, 13.5, 15.4, 17.5,
	}
)

type StarData struct {
	Index        int32
	Level        float64
	ID           int32
	Seed         int32
	ResourceCoef float64
	Name         string
	//Position           PositionSingle
	Mass               float64
	Age                float64
	Lifetime           float64
	Temperatore        float64
	Temperature        float64
	Type               StarType
	Spectr             SpectralType
	Color              float64
	ClassFactor        float64
	Luminosity         float64
	Radius             float64
	AcdiskRadius       float64
	HabitableRadius    float64
	LightBalanceRadius float64
	OrbitScaler        float64
	DysonRadius        float64
	PhysicsRadius      float64
}

func makeBirthStar(seed int32) *StarData {
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
	num4 = clamp(num4, -0.2, 0.2)
	star.Mass = math.Pow(2, num4)
	d := 2.0 + 0.4*(1.0-star.Mass)
	star.Lifetime = (10000.0 * math.Pow(0.1, math.Log10(star.Mass*0.5)/math.Log10(d)+1.0) * num2)
	star.Age = num*0.4 + 0.3

	num5 := (1 - math.Pow(clamp01(star.Age), 20)*0.5) * star.Mass
	star.Temperature = math.Pow(num5, 0.56+0.14/(math.Log10(num5+4)/math.Log10(5.0)))*4450.0 + 1300.0
	num6 := math.Log10((star.Temperature-1300.0)/4500.0)/math.Log10(2.6) - 0.5
	if num6 < 0.0 {
		num6 *= 4.0
	}
	if num6 > 2.0 {
		num6 = 2.0
	} else if num6 < -4.0 {
		num6 = -4.0
	}
	star.Spectr = SpectralType(math.Round(num6 + 4))
	star.Color = clamp01((num6 + 3.5) * 0.2)
	star.ClassFactor = num6
	star.Luminosity = math.Pow(float64(num5), 0.7)
	star.Radius = math.Pow(float64(star.Mass), 0.4) * num3
	star.AcdiskRadius = 0
	p := num6 + 2
	star.HabitableRadius = math.Pow(1.7, p) + 0.2*math.Min(1, star.OrbitScaler)
	star.LightBalanceRadius = math.Pow(1.7, p)
	star.OrbitScaler = math.Pow(1.35, p)
	if star.OrbitScaler < 1 {
		star.OrbitScaler = lerp(star.OrbitScaler, 1, 0.6)
	}
	setStarAge(star, star.Age, rn, rt)
	star.DysonRadius = star.OrbitScaler * 0.28
	if star.DysonRadius*40000.0 < (star.PhysicsRadius * 1.5) {
		star.DysonRadius = ((star.PhysicsRadius * 1.5) / 40000.0)
	}
	galaxy := []string{}
	star.Name = randomStarName(seed2, star, galaxy)
	return star
}

func setStarAge(star *StarData, age float64, rn float64, rt float64) {
	num := (rn*0.1 + 0.95)
	num2 := (rt*0.4 + 0.8)
	num3 := (rt*9.0 + 1.0)
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
		num4 := (math.Pow(5.0, math.Abs(math.Log10(star.Mass)-0.7)) * 5.0)
		if num4 > 10 {
			num4 = (math.Log(num4*0.1) + 1) * 10
		}
		num5 := 1 - math.Pow(star.Age, 30)*0.5
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

func lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*t
}
