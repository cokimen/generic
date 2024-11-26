package main

import "fmt"

func main() {

	// 0001
	var a int32 = 1

	// 0100 -> 4
	var b int32 = 4

	// 0001
	// 0100
	fmt.Println(a & b)

	// 0111 -> 7
	// 0001 -> 1

	fmt.Println(7 & 1)
	fmt.Println(1 & 7)
	fmt.Println(1 < 7)
	fmt.Println(7 | 1)

	switch 11 {
	case 1, 2, 3:
		fmt.Println("1,2,3")
	case 10, 11:
		fmt.Println("10,11")
		fallthrough
	case 12, 13:
		fmt.Println("12,13")
	default:
		fmt.Println("separate")
	}

}
