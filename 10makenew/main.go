package main

import "fmt"

func main() {
	var a2 map[string]string // default if no init value , or using make its pointing to nil
	var a3 *map[string]string = new(map[string]string)
	fmt.Println(a2)
	fmt.Println(a3)
	a3 = &a2
	a2 = make(map[string]string)
	a2["firstnme"] = "adudu"
	fmt.Println(*a3)
}
