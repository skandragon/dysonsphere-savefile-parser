package main

import (
	"math"

	"github.com/skandragon/dysonsphere/internal/cs"
	"github.com/skandragon/dysonsphere/internal/cs/mathf"
	"github.com/skandragon/dysonsphere/types"
)

// Universe -- star data, which is all auto-generated, not loaded
// from a save file.
type Universe struct {
	Stars     []*StarData `json:"stars"`
	Seed      int32       `json:"seed"`
	StarCount int32       `json:"star_count"`
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
	seed := random.Next()

	actualStarCount := generateTempPoses(seed, gd.StarCount, 4, 2.0, 2.3, 3.5, 0.18)

	stars := make([]*StarData, 0)

	universe := &Universe{
		Stars:     stars,
		StarCount: actualStarCount,
		Seed:      gd.GalaxySeed,
	}

	num2 := float32(random.NextDouble())
	num3 := float32(random.NextDouble())
	num4 := float32(random.NextDouble())
	num5 := float32(random.NextDouble())
	num6 := mathf.CeilToInt(0.01*float32(actualStarCount) + num2*0.3)
	num7 := mathf.CeilToInt(0.01*float32(actualStarCount) + num3*0.3)
	num8 := mathf.CeilToInt(0.016*float32(actualStarCount) + num4*0.4)
	num9 := mathf.CeilToInt(0.013*float32(actualStarCount) + num5*1.4)
	num10 := actualStarCount - num6
	num11 := num10 - num7
	num12 := num11 - num8
	num13 := (num12 - 1) / num9
	num14 := num13 / 2

	for i := int32(0); i < actualStarCount; i++ {
		seed2 := random.Next()
		if i == 0 {
			star := makeBirthStar(seed2, universe)
			universe.Stars = append(universe.Stars, star)
		} else {
			needSpectr := types.SpectralTypeX
			if i == 3 {
				needSpectr = types.SpectralTypeM
			} else if i == num12-1 {
				needSpectr = types.SpectralTypeO
			}
			needType := types.StarTypeMainSeqStar
			if i%num13 == num14 {
				needType = types.StarTypeGiantStar
			}
			if i >= num10 {
				needType = types.StarTypeBlackHole
			} else if i >= num11 {
				needType = types.StarTypeNeutronStar
			} else if i >= num12 {
				needType = types.StarTypeWhiteDwarf
			}
			star := createStar(universe, tmpPoses[i], i+1, seed2, needType, needSpectr)
			universe.Stars = append(universe.Stars, star)
		}
	}
	// TODO: astro poses
	// TODO: create planets
	// TODO: generate veins

	return universe
}

var (
	tmpPoses = make([]*VectorLF3, 0)
	tmpDrunk = make([]*VectorLF3, 0)
)

func generateTempPoses(seed int32, targetCount int32, iterCount int32, minDist float64, minStepLen float64, maxStepLen float64, flatten float64) int32 {
	tmpPoses = make([]*VectorLF3, 0)
	tmpDrunk = make([]*VectorLF3, 0)

	if iterCount < 1 {
		iterCount = 1
	} else if iterCount > 16 {
		iterCount = 16
	}
	randomPoses(seed, targetCount*iterCount, minDist, minStepLen, maxStepLen, flatten)

	for i := len(tmpPoses) - 1; i >= 0; i-- {
		if int32(i)%iterCount != 0 {
			copy(tmpPoses[i:], tmpPoses[i+1:])
			tmpPoses[len(tmpPoses)-1] = nil
			tmpPoses = tmpPoses[:len(tmpPoses)-1]
		}
		if int32(len(tmpPoses)) <= targetCount {
			break
		}
	}
	return int32(len(tmpPoses))
}

func randomPoses(seed int32, maxCount int32, minDist float64, minStepLen float64, maxStepLen float64, flatten float64) {
	random := cs.MakePRNGSequence(seed)
	num := random.NextDouble()
	tmpPoses = append(tmpPoses, VectorLF3Zero())
	num2 := 6
	num3 := 8
	num4 := int(num*float64(num3-num2) + float64(num2))
	for i := 0; i < num4; i++ {
		num5 := 0
		for num5 < 256 {
			num5++
			num6 := random.NextDouble()*2 - 1
			num7 := (random.NextDouble()*2 - 1) * flatten
			num8 := random.NextDouble()*2 - 1
			num9 := random.NextDouble()
			num10 := num6*num6 + num7*num7 + num8*num8
			if num10 < 1 {
				if num10 >= 1e-08 {
					num11 := math.Sqrt(num10)
					num9 = (num9*(maxStepLen-minStepLen) + minDist) / num11
					v := &VectorLF3{
						X: num6 * num9,
						Y: num7 * num9,
						Z: num8 * num9,
					}
					if !checkCollision(tmpPoses, v, minDist) {
						tmpPoses = append(tmpPoses, v)
						tmpDrunk = append(tmpDrunk, v)
						if int32(len(tmpPoses)) >= maxCount {
							return
						}
						break
					}
				}
			}
		}
	}
	num12 := 0
	for num12 < 256 {
		num12++
		for j := 0; j < len(tmpDrunk); j++ {
			num13 := random.NextDouble()
			if num13 < 0.7 {
				num14 := 0
				for num14 < 256 {
					num14++
					num15 := random.NextDouble()*2 - 1
					num16 := (random.NextDouble()*2 - 1) * flatten
					num17 := random.NextDouble()*2 - 1
					num18 := random.NextDouble()
					num19 := num15*num15 + num16*num16 + num17*num17
					if num19 <= 1 {
						if num19 >= 1e-08 {
							num20 := math.Sqrt(num19)
							num18 = (num18*(maxStepLen-minStepLen) + minDist) / num20
							v := &VectorLF3{
								X: tmpDrunk[j].X + num15*num18,
								Y: tmpDrunk[j].Y + num16*num18,
								Z: tmpDrunk[j].Z + num17*num18,
							}
							if !checkCollision(tmpPoses, v, minDist) {
								tmpDrunk[j] = v
								tmpPoses = append(tmpPoses, v)
								if int32(len(tmpPoses)) >= maxCount {
									return
								}
								break
							}
						}
					}
				}
			}
		}
	}
}

func checkCollision(pts []*VectorLF3, pt *VectorLF3, minDist float64) bool {
	distSquared := minDist * minDist
	for _, v := range pts {
		x := pt.X - v.X
		y := pt.Y - v.Y
		z := pt.Z - v.Z
		sq := x*x + y*y + z*z
		if sq < distSquared {
			return true
		}
	}
	return false
}
