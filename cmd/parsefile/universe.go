package main

import (
	"github.com/skandragon/dysonsphere/internal/cs"
)

// Universe -- star data, which is all auto-generated, not loaded
// from a save file.
type Universe struct {
	Stars []*StarData `json:"stars"`
	Seed  int32       `json:"seed"`
}

// UniqueStarname -- returns true if no other star has this name.
func (u *Universe) UniqueStarname(name string) bool {
	for _, star := range u.Stars {
		if star.Name == name {
			return false
		}
	}
	return true
}

// MakeUniverse generates the universe and all stars from a single PRNG seed.
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

	universe := &Universe{
		Stars: stars,
	}

	for i := int32(0); i < gd.StarCount; i++ {
		seed2 := random.Next()
		if i == 0 {
			stars = append(stars, makeBirthStar(seed2, universe))
		} else {
			//stars = append(stars, makeBirthStar(seed2))
		}
		//makeStar(seed2)
	}

	return universe
}

type VectorLF3 struct {
	X float64
	Y float64
	Z float64
}

var (
	tmpPoses  = make([]*VectorLF3, 0)
	tempDrunk = make([]*VectorLF3, 0)
)

func generateTempPoses(seed int32, targetCount int32, iterCount int32, minDist float64, minStepLen float64, maxStepLen float64, flatten float64) int32 {
	tmpPoses = make([]*VectorLF3, 0)
	tempDrunk = make([]*VectorLF3, 0)

	if iterCount < 1 {
		iterCount = 1
	} else if iterCount > 16 {
		iterCount = 16
	}
	randomPoses(seed, targetCount*iterCount, minDist, minStepLen, maxStepLen, flatten)

	for i := len(tmpPoses) - 1; i >= 0; i-- {
		if int32(i)%iterCount != 0 {
			copy(tmpPoses[i:], tmpPoses[i+1:])
			tmpPoses[len(tmpPoses)] = nil
			tmpPoses = tmpPoses[:len(tmpPoses)-1]
		}
		if int32(len(tmpPoses)) <= targetCount {
			break
		}
	}
	return int32(len(tmpPoses))
}
