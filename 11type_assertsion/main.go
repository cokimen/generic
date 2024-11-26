package main

import (
	"fmt"
	"reflect"
)

func main() {

	var genericsDatas []interface{} = make([]interface{}, 3)
	genericsDatas[0] = "Please" // you can use append here
	genericsDatas[1] = int32(900)
	genericsDatas[2] = float32(900.88)

	for k, v := range genericsDatas {
		fmt.Printf("Index num : %v \n ", k)
		switch v.(type) {
		case string:
			fmt.Println("this iss string")
		case int32:
			fmt.Println("this is int32")
		case float32:
			fmt.Println("this is float32")
		default:
			fmt.Println("this is ", reflect.TypeOf(v))
		}
	}

	var sample1 interface{} = int32(900)

	fmt.Println("--------------------------------------------------")
	if value, resBool := sample1.(string); resBool {
		fmt.Println(value)
	} else {
		fmt.Println("Not String")
		fmt.Printf("(interface{})(value) : %T \n", (interface{})(sample1))
		fmt.Printf("(interface{})(value) : %v \n", (interface{})(sample1))
		fmt.Printf("reflect.TypeOf(value) : %t \n", reflect.TypeOf(sample1))
		fmt.Printf("reflect.TypeOf(value) : %T \n", reflect.TypeOf(sample1))
		fmt.Printf("reflect.TypeOf(value) : %v \n", reflect.TypeOf(sample1))
	}

	fmt.Println("--------------------------------------------------")
	sample1 = string("Oke")
	if value, resBool := sample1.(string); resBool {
		fmt.Println("This is String")
		fmt.Printf("reflect.TypeOf(value) : %v \n", reflect.TypeOf(value))
		fmt.Println(value)
	} else {
		fmt.Println("Not String")
		fmt.Printf("(interface{})(value) : %V \n", (interface{})(value))
		fmt.Printf("reflect.TypeOf(value) : %v \n", reflect.TypeOf(value))
	}
}
