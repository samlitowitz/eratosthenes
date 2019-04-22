package sieve

import (
	"math"
)

type Segmented struct {
	max, segmentLength int64
	primes             []int64
}

// NewSegmented creates a new segmented implementation of the
// Sieve of Eratosthenes. An error is returned if a maximum
// value of less then 2.
func NewSegmented(max, segmentLength int64) (*Segmented, error) {
	if max < 2 {
		return nil, &OutOfBoundsError{
			Max: math.MaxInt64,
			Min: 2,
			N:   max,
		}
	}
	if segmentLength < 1 {
		return nil, &OutOfBoundsError{
			Max: math.MaxInt64,
			Min: 1,
			N:   int64(segmentLength),
		}
	}
	return &Segmented{
		max:           max,
		segmentLength: segmentLength,
	}, nil
}

// Max returns the upper bound of the search space.
func (n *Segmented) Max() int64 {
	return n.max
}

// Primes returns all prime numbers in [2, Max()].
// The result of the calculation is stored and
// referenced until Reset is called.
func (n *Segmented) Primes() []int64 {
	if n.primes == nil {
		n.primes = n.calculatePrimes()
	}

	return n.primes
}

// Reset clears out the calculated primes forcing them
// to be recalculated when calling Primes().
func (n *Segmented) Reset() {
	n.primes = nil
}

func (n *Segmented) calculatePrimes() []int64 {
	primesMap := make(map[int64]int64)
	segmentCount := n.max / n.segmentLength + 1

	for i := int64(0); i < segmentCount; i++ {
		segment := make([]bool, n.segmentLength+1)
		min, max := i*n.segmentLength, (i+1)*n.segmentLength
		if max > n.max {
			max = n.max
		}
		if min < 2 {
			min = 2
		}

		// mark multiples of existing primesMap in current segment
		for p, next := range primesMap {
			for j := next; j <= max; j += p {
				segment[j] = true
				next = j;
			}
			primesMap[p] = next + p
		}

		// sift for new primesMap and mark their multiples in current segment
		for j := min; j <= max; j++ {
			if segment[j] {
				continue
			}

			next := 2 * j
			primesMap[j] = next
			for k := 2 * j; k <= max; k += j {
				segment[k] = true
				next = k
			}
			primesMap[j] = next + j
		}
	}

	primes := make([]int64, 0)
	for i := range primesMap {
		primes = append(primes, i)
	}

	return primes
}
