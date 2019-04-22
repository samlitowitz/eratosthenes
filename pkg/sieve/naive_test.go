package sieve_test

import (
	"math"
	"testing"

	"github.com/go-test/deep"
	"github.com/samlitowitz/eratosthenes/pkg/sieve"
)

func TestNewNaive(t *testing.T) {
	testData := []struct {
		n        int64
		expected struct {
			isNaiveNil bool
			max        int64
			primes     []int64
			err        error
		}
	}{
		{
			n: 1,
			expected: struct {
				isNaiveNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isNaiveNil: true,
				max:        math.MaxInt64,
				primes:     []int64{},
				err:        &sieve.OutOfBoundsError{math.MaxInt64, 2, 1},
			},
		},
		{
			n: 2,
			expected: struct {
				isNaiveNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isNaiveNil: false,
				max:        2,
				primes:     []int64{2},
				err:        nil,
			},
		},
		{
			n: 5,
			expected: struct {
				isNaiveNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isNaiveNil: false,
				max:        5,
				primes:     []int64{2, 3, 5},
				err:        nil,
			},
		},
	}

	for _, data := range testData {
		naive, err := sieve.NewNaive(data.n)

		if data.expected.isNaiveNil {
			if naive != nil {
				t.Errorf("expected naive to be nil")
			}
		} else {
			if data.expected.max != naive.Max() {
				t.Errorf("expected %d to equal %d", data.expected.max, naive.Max())
			}

			if diff := deep.Equal(data.expected.primes, naive.Primes()); diff != nil {
				t.Error(diff)
			}
		}

		if diff := deep.Equal(data.expected.err, err); diff != nil {
			t.Error(diff)
		}

	}
}

func TestNaive_Max(t *testing.T) {
	boundsTests := []struct {
		n        int64
		expected int64
	}{
		{
			n:        2,
			expected: 2,
		},
		{
			n:        3,
			expected: 3,
		},
		{
			n:        4,
			expected: 4,
		},
	}

	for _, data := range boundsTests {
		naive, _ := sieve.NewNaive(data.n)

		if data.expected != naive.Max() {
			t.Errorf("expected %d to equal %d", data.expected, naive.Max())
		}
	}
}

func TestNaive_Primes(t *testing.T) {
	testData := []struct {
		n        int64
		expected struct {
			primes []int64
		}
	}{
		{
			n: 2,
			expected: struct {
				primes []int64
			}{
				primes: []int64{2},
			},
		},
		{
			n: 5,
			expected: struct {
				primes []int64
			}{
				primes: []int64{2, 3, 5},
			},
		},
	}

	for _, data := range testData {
		naive, _ := sieve.NewNaive(data.n)

		if diff := deep.Equal(data.expected.primes, naive.Primes()); diff != nil {
			t.Error(diff)
		}
	}
}

var resultNaive []int64

func benchmarkNaive_Primes(max int64, b *testing.B) {
	var p []int64
	naive, _ := sieve.NewNaive(max)
	for n := 0; n < b.N; n++ {
		p = naive.Primes()
		naive.Reset()
	}
	resultNaive = p
}

func BenchmarkNaive_Primes256(b *testing.B) {
	benchmarkNaive_Primes(256, b)
}

func BenchmarkNaive_Primes512(b *testing.B) {
	benchmarkNaive_Primes(512, b)
}
