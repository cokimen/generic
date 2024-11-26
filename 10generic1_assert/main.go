package main

import (
	"fmt"
)

func handler[T string | int32 | interface{}](data T) T {
	// var retVal T
	if _, ok := (interface{})(data).(string); ok {
		fmt.Println("Data Type is String")
	} else if _, ok := (interface{})(data).(int32); ok {
		fmt.Println("Data Type is int32")
	} else if _, ok := (interface{})(data).(struct{ firstname string }); ok {
		fmt.Println("Data Type is struct{firstname string}")
	} else if _, ok := (interface{})(data).(interface{}); ok {
		fmt.Println("Data Type is interface")
	}
	return data
}

func main() {

	var a string = handler("Oke")
	fmt.Println(a)
	var b int32 = handler(int32(1234))
	fmt.Println(b)

	var c struct {
		firstname string
	} = handler(struct {
		firstname string
	}{
		firstname: "erz",
	})

	fmt.Println(c)

	var d []string = handler([]string{"erhan", "burhanudin"})

	fmt.Println(d)
}
