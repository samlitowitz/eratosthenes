package sieve

// The Sieve interface is implements by types which calculate
// prime numbers using the Sieve of Eratosthenes algorithm.
// See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes for
// more information.
type Sieve interface {
	// Max returns the upper bound of the search space.
	Max() int64
	// Primes returns all prime numbers in [2, Max()].
	Primes() []int64
}
