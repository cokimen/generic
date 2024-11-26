package main

import "fmt"

func show() {
	fmt.Println("Void Function")
}

func Add(a int16, b int16) int32 {
	return int32(a + b)
}

func lambda(s string) func() {
	return func() {
		println(s)
	}
}

func closure(s string) func() int32 {
	return func() int32 {
		return int32(len(s))
	}
}

func handler(fn func(s ...string) int) {
	var result int = fn("Ramanujan", "Tao Terence", "")
	fmt.Printf("Emitt Data Package %d \n", result)
}

func controller(fn func(s string)) {
	var s string = "erhan burhanudin"
	fn(s)
}

func main() {

	var a func() = lambda("erz")
	a()

	var b func() int32 = closure("burhanudin")

	fmt.Println(b())

	handler(func(s ...string) int {
		return len(s) + 10
	})

	var fn2 func(s string) = func(i int) func(string) {
		return func(s string) {
			fmt.Println(s, "Oke")
		}
	}(10)

	fn2("90")

	controller(func(i int) func(string) {

		fmt.Println("Tambahan dari gw 1")
		fmt.Println("Tambahan dari gw 2")
		fmt.Println("Tambahan dari gw 3")
		return func(s string) {
			fmt.Println(s, "Erbur")
		}

	}(10))

}
