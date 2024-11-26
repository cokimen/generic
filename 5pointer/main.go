package main

import "fmt"

func main() {

	var dt1 *struct{ firstname string } = new(struct{ firstname string })

	var dt2 struct{ firstname string } = struct{ firstname string }{firstname: "cokimen"}

	dt1 = &dt2

	fmt.Println(dt1)

	var dt3 *struct{ location string } = new(struct{ location string })
	dt3 = &struct{ location string }{location: "burhanudin"}

	fmt.Println(*dt3)

	var listUser [3]int8 = [...]int8{10, 20, 40}

	// array copy by value
	var listUsername [3]int8 = listUser

	fmt.Println(&listUser[0])
	fmt.Println(&listUser[1])
	fmt.Println(&listUser[2])

	listUsername[0] = 100
	fmt.Println(listUser)
	fmt.Println(listUsername)
}
