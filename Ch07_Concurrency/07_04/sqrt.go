package sqrt

import (
	"errors"
)

var (
	ErrNegSqrt    = errors.New("norgate math: square root of negative number")
	ErrNoSolution = errors.New("no solution found")
)

// Abs returns the absolute value of x
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Sqrt returns the square root of x
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegSqrt
	}

	if x == 0.0 {
		return 0.0, nil
	}

	guess, epsilon := 1.0, 0.00001
	for i := 0; i < 10000; i++ {
		if Abs(guess*guess-x) <= epsilon {
			return guess, nil
		}
		guess = (x/guess + guess) / 2.0
	}

	return 0.0, ErrNoSolution
}
