package grains

import (
	"errors"
)

// Square returns the number of grains on the square for the number passed in.
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("square number must be 1 through 64")
	}

	return 1 << (num - 1), nil
}

// Total return the sum of the squares from 1 to 64, which is the number of squares on a chess board.
func Total() uint64 {
	return 1<<64 - 1
}
