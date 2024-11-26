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

type CustomErrorFajri struct {
	ErrorMessage string
}

func (cst *CustomErrorFajri) Error() string {
	return cst.ErrorMessage
}

func main() {
	var abc LocalNewError = CustomError{ErrorMessage: "validation Error"}
	var z error = abc
	fmt.Println(z)

	var abd error = CustomError{ErrorMessage: "validation Error for Username"}
	fmt.Println(abd)

	//
	var cstErr error = &CustomErrorFajri{ErrorMessage: "Apapun Itu"}
	fmt.Println(cstErr)

	if cstErr != nil {
		panic("Error Bro")
	}
}
