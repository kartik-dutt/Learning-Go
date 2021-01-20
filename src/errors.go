package main

import (
	"errors"
	"fmt"
)

// Error handling -> Return error.
func divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}

	return -1, errors.New("Division by 0 isn't possible.")
}

func div(c, d int) {
	a, b := divide(c, d)
	if b == nil {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

// By default for a class error function is called.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	div(6, 3)
	div(9, 0)
	fmt.Println(f2(42))
}
