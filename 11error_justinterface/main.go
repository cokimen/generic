package main

import "fmt"

type CustomError struct {
	ErrorMessage string
}

func (s CustomError) Error() string {
	return s.ErrorMessage
}

type LocalNewError interface {
	Error() string
}

func main() {
	var abc LocalNewError = CustomError{ErrorMessage: "validation Error"}
	var z error = abc
	fmt.Println(z)

	var abd error = CustomError{ErrorMessage: "validation Error for Username"}
	fmt.Println(abd)

}
