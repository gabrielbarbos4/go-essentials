package internal

import (
	"errors"
	"fmt"
)

func Functions() {
	fmt.Print()

	var result, remainder, err = intDivision(11, 15)

	if err != nil {
		fmt.Printf(err.Error())
	}

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		var err error = errors.New("Cannot ")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
