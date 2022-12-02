package calculator

import "errors"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return b - a
}
func Muti(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
