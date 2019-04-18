package sieve

import "math"

// A Naive is a naive implementation of the Sieve of Eratosthenes.
type Naive struct {
	max    int64
	primes []int64
}

// NewNaive creates a new naive implementation of the Sieve of Eratosthenes.
// An error is returned if a maximum value of less then 2.
func NewNaive(max int64) (*Naive, error) {
	if max < 2 {
		return nil, &OutOfBoundsError{
			Max: math.MaxInt64,
			Min: 2,
			N:   max,
		}
	}
	return &Naive{
		max: max,
	}, nil
}

// Max returns the upper bound of the search space.
func (n *Naive) Max() int64 {
	return n.max
}

// Primes returns all prime numbers in [2, Max()].
// The calculation is only performed once.
func (n *Naive) Primes() []int64 {
	if n.primes == nil {
		n.primes = n.calculatePrimes()
	}

	return n.primes
}

func (n *Naive) calculatePrimes() []int64 {
	numbers := make([]bool, n.max+1)

	for i := int64(2); i <= n.max; i++ {
		if numbers[i] {
			continue
		}

		for j := 2 * i; j <= n.max; j += i {
			numbers[j] = true
		}
	}

	primes := make([]int64, 0)
	for i := int64(2); i <= n.max; i++ {
		if numbers[i] {
			continue
		}

		primes = append(primes, i)
	}

	return primes
}
