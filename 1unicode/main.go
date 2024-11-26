package main

import (
	"fmt"
	"unicode/utf8"
)

var a1 byte
var a2 rune
var a3 string
var a4 string

// for container unicode and ascii byte

func main() {
	// the most high unicode
	a2 = 999998
	fmt.Println(a1, string(a2), a3)

	// len(string(a2)) , 4 byte
	fmt.Println(a1, len(string(a2)), a3)

	fmt.Println(a1, len([]byte(string(a2))), a3)

	var conversion []byte = []byte(string(a2))

	for i := 0; i < len(string(a2)); i++ {
		fmt.Println(conversion[i])
	}

	a4 = "happen发生"

	fmt.Println("----")

	for _, v := range a4 {
		fmt.Println(string(v))
	}

	var a5 int = utf8.RuneCountInString(a4)

	fmt.Println(a5)

	// all ascii
	var a6 string = "kali-bata"
	fmt.Println(utf8.RuneCountInString(a6))

	fmt.Println("----")
	fmt.Println(utf8.RuneCount([]byte(a6)))

}
