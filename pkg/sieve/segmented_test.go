package sieve_test

import (
	"github.com/go-test/deep"
	"github.com/samlitowitz/eratosthenes/pkg/sieve"
	"math"
	"testing"
)

func TestNewSegmented(t *testing.T) {
	testData := []struct {
		n        int64
		expected struct {
			isSegmentedNil bool
			max            int64
			primes         []int64
			err            error
		}
	}{
		{
			n: 1,
			expected: struct {
				isSegmentedNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isSegmentedNil: true,
				max:        math.MaxInt64,
				primes:     []int64{},
				err:        &sieve.OutOfBoundsError{math.MaxInt64, 2, 1},
			},
		},
		{
			n: 2,
			expected: struct {
				isSegmentedNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isSegmentedNil: false,
				max:        2,
				primes:     []int64{2},
				err:        nil,
			},
		},
		{
			n: 5,
			expected: struct {
				isSegmentedNil bool
				max        int64
				primes     []int64
				err        error
			}{
				isSegmentedNil: false,
				max:        5,
				primes:     []int64{2, 3, 5},
				err:        nil,
			},
		},
	}

	for _, data := range testData {
		segmented, err := sieve.NewSegmented(data.n, 64)

		if data.expected.isSegmentedNil {
			if segmented != nil {
				t.Errorf("expected segmented to be nil")
			}
		} else {
			if data.expected.max != segmented.Max() {
				t.Errorf("expected %d to equal %d", data.expected.max, segmented.Max())
			}

			if diff := deep.Equal(data.expected.primes, segmented.Primes()); diff != nil {
				t.Error(diff)
			}
		}

		if diff := deep.Equal(data.expected.err, err); diff != nil {
			t.Error(diff)
		}

	}
}

func TestSegmented_Max(t *testing.T) {
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
		segmented, _ := sieve.NewSegmented(data.n, 64)

		if data.expected != segmented.Max() {
			t.Errorf("expected %d to equal %d", data.expected, segmented.Max())
		}
	}
}

func TestSegmented_Primes(t *testing.T) {
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
		segmented, _ := sieve.NewNaive(data.n)

		if diff := deep.Equal(data.expected.primes, segmented.Primes()); diff != nil {
			t.Error(diff)
		}
	}
}

var resultSegmented []int64

func benmarkSegmented_Primes(max int64, b *testing.B) {
	var p []int64
	segmented, _ := sieve.NewSegmented(max, 64)
	for n := 0; n < b.N; n++ {
		p = segmented.Primes()
		segmented.Reset()
	}
	resultSegmented = p
}

func BenchmarkSegmented_Primes256(b *testing.B) {
	benmarkSegmented_Primes(256, b)
}

func BenchmarkSegmented_Primes512(b *testing.B) {
	benmarkSegmented_Primes(512, b)
}
