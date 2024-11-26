package main

import "fmt"

func init() {
	fmt.Println("Always Call in when you use package")
}

func onpurpose_throw() {
	panic("this is not incident")
}

func main() {

	defer func() {
		fmt.Println("I am executed even error is happen")
	}()

	fmt.Println("running")

	// just comment this to make sure how defer works
	onpurpose_throw()

}
