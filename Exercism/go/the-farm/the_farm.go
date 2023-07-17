package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calculator FodderCalculator, amount int) (float64, error) {
	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	fodder, err := calculator.FodderAmount(amount)
	if err != nil {
		return 0, err
	}
	return (fodder * factor) / float64(amount), nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, amount int) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calculator, amount)
}

type InvalidCowsError struct {
	message string
}

func (err *InvalidCowsError) Error() string {
	return err.message
}
func ValidateNumberOfCows(amount int) error {
	if amount < 0 {
		msg := fmt.Sprintf("%d cows are invalid: there are no negative cows", amount)
		return &InvalidCowsError{message: msg}
	} else if amount == 0 {
		msg := fmt.Sprintf("%d cows are invalid: no cows don't need food", amount)
		return &InvalidCowsError{message: msg}
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
