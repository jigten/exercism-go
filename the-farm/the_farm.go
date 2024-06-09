package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, numOfCows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(numOfCows)

	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return (fodderAmount * fatteningFactor) / float64(numOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numOfCows int) (float64, error) {
	if numOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalculator, numOfCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.customMessage)
}

func ValidateNumberOfCows(numOfCows int) error {
	if numOfCows < 0 {
		return &InvalidCowsError{numOfCows, "there are no negative cows"}
	}

	if numOfCows == 0 {
		return &InvalidCowsError{numOfCows, "no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
