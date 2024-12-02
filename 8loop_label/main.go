package main

import "fmt"

type Baams struct {
	firstname string
}

func main() {
	var rabil interface{} = float32(10.99)

	if value, isTrue := rabil.(string); isTrue {
		fmt.Print(value, "its string")
	} else if value, isTrue := rabil.(float32); isTrue {
		fmt.Println(value, "its float32")
	}

	defer func() {
		recover()
		fmt.Println("I am not panic")
	}()
	var simple int = 0
LOOP1:
	for {
		fmt.Println("oKE")
		simple++
		if simple >= 10 {
			fmt.Println("Cokimen", simple)
			break
		}
		goto LOOP1
	}

	var b1 *Baams = new(Baams)
	b1.firstname = "cokimen"
	fmt.Println(b1)

	b1 = nil

	fmt.Println(b1.firstname)

	fmt.Println("Cirendeu")

}
