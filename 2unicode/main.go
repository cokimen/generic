package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var s string = "สวัสดี"
	// fmt.Println(len(s))
	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// s = "春🎹"

	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// var german_characters string = "ÄÖÜß"
	// fmt.Println("Rune count:", utf8.RuneCountInString(german_characters))

	var a []rune = []rune("Καλημέρα κόσμε")
	fmt.Println("Length as Hardware eye Of Rune", len(a))

	var b string = "Καλημέρα κόσμε"
	fmt.Println("Length as Hardware eye", len(b))
	fmt.Println("Length as Human eye", utf8.RuneCountInString(b))

}
