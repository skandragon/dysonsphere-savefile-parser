/*
 * Copyright 2021-2022 Michael Graff
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cs

import "math"

//
// PRNGSequence implements the really lame, but apparently frequently used c# Random class.
//
type PRNGSequence struct {
	inext     int32
	inextp    int32
	seedarray []int32
}

func iabs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//
// MakePRNGSequence takes a seed, and produces (via other methods) a "random" stream.  It's really
// not random, but given the same seed it will produce the same sequence.
//
func MakePRNGSequence(seed int32) *PRNGSequence {
	r := &PRNGSequence{
		inext:     0,
		inextp:    31,
		seedarray: make([]int32, 56),
	}

	var ii int32

	mj := 161803398 - iabs(seed)
	r.seedarray[55] = mj

	var mk int32 = 1
	for i := int32(1); i < 55; i++ {
		ii = 21 * i % 55
		r.seedarray[ii] = mk
		mk = mj - mk
		if mk < 0 {
			mk += math.MaxInt32
		}
		mj = r.seedarray[ii]
	}
	for k := 1; k < 5; k++ {
		for i := 1; i < 56; i++ {
			r.seedarray[i] -= r.seedarray[1+(i+30)%55]
			if r.seedarray[i] < 0 {
				r.seedarray[i] += math.MaxInt32
			}
		}
	}

	return r
}

//
// Next returns the next value in the sequence.
//
func (r *PRNGSequence) Next() int32 {
	return int32(r.internalSample() * math.MaxInt32)
}

//
// NextDouble returns a floating point value between 0 and 1, inclusive.
//
func (r *PRNGSequence) NextDouble() float64 {
	return r.internalSample()
}

//
// NextWithMax returns a random value between 0 and the max value provided,
// inclusive.
//
func (r *PRNGSequence) NextWithMax(max int32) int32 {
	return int32(r.internalSample() * float64(max))
}

//
// NextRange returns a value between min and max, inclusive.
//
func (r *PRNGSequence) NextRange(min int32, max int32) int32 {
	span := uint32(max - min)
	if span <= 1 {
		return min
	}
	return int32(uint64(uint32(r.internalSample()*float64(span))) + uint64(int64(min)))
}

func (r *PRNGSequence) internalSample() float64 {
	var locINext = r.inext
	var locINextp = r.inextp

	locINext++
	if locINext >= 56 {
		locINext = 1
	}
	locINextp++
	if locINextp >= 56 {
		locINextp = 1
	}
	retVal := r.seedarray[locINext] - r.seedarray[locINextp]
	if retVal < 0 {
		retVal += math.MaxInt32
	}
	r.seedarray[locINext] = retVal

	r.inext = locINext
	r.inextp = locINextp

	return float64(retVal) * 4.656612875245797e-10
}
