package sieve

import (
	"fmt"
)

// OutOfBoundsError represents and out of bounds error.
type OutOfBoundsError struct {
	Max, Min, N int64
}

func (e *OutOfBoundsError) Error() string {
	return fmt.Sprintf("%d out of bounds [%d, %d]", e.N, e.Min, e.Max)
}
