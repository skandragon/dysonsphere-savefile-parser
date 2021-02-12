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
		inextp:    21,
		seedarray: make([]int32, 56),
	}

	var ii int32

	var subtraction int32
	if seed == math.MinInt32 {
		subtraction = math.MaxInt32
	} else {
		subtraction = iabs(seed)
	}

	mj := 161803398 - subtraction
	r.seedarray[55] = mj

	var mk int32 = 1
	var i int32
	for i = 1; i < 55; i++ {
		ii = (21 * i) % 55
		r.seedarray[ii] = mk
		mk = mj - mk
		if mk < 0 {
			mk += math.MaxInt32
		}
		mj = r.seedarray[ii]
	}
	for k := 1; k < 5; k++ {
		for i = 1; i < 56; i++ {
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
	return r.internalSample()
}

func (r *PRNGSequence) internalSample() int32 {
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
	if retVal == math.MaxInt32 {
		retVal--
	}
	if retVal < 0 {
		retVal += math.MaxInt32
	}
	r.seedarray[locINext] = retVal

	r.inext = locINext
	r.inextp = locINextp

	return retVal
}

func (r *PRNGSequence) NextDouble() float64 {
	return r.sample()
}

func (r *PRNGSequence) sample() float64 {
	return float64(r.internalSample()) * (1.0 / float64(math.MaxInt32))
}

func (r *PRNGSequence) NextWithMax(max int32) int32 {
	return int32(r.NextDouble() * float64(max))
}

func (r *PRNGSequence) NextRange(min int32, max int32) int32 {
	span := int64(max) - int64(min)
	if span <= math.MaxInt32 {
		return int32(r.sample()*float64(span)) + min
	}
	return int32(r.getSampleForLargeRange()*float64(span)) + min
}

func (r *PRNGSequence) getSampleForLargeRange() float64 {
	result := r.internalSample()
	negative := r.internalSample()%2 == 0
	if negative {
		result = -result
	}
	d := float64(result)
	d += (math.MaxInt32 - 1)
	d /= 2*math.MaxInt32 - 1
	return d
}
