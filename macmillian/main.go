package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type Class struct {
	firstname string
}

func (c Class) create() interface{} {
	return c
}

type Meta struct {
	width int
}

func (m Meta) create() interface{} {
	return m
}

type Action interface {
	create() interface{}
}

type CustomerStruct struct {
	firstname string
	birthdate time.Month
}

type Customer CustomerStruct

func main() {

	var act Action = Meta{width: 900}
	fmt.Println(act)

	fmt.Println(
		reflect.TypeOf(act),
	)
	act = Class{firstname: "Erhan"}
	fmt.Println(act)

	fmt.Println(
		reflect.TypeOf(act),
	)

	var cust Customer = Customer(struct {
		firstname string
		birthdate time.Month
	}{
		firstname: "cokie",
		birthdate: time.Month(90),
	})

	fmt.Println(cust)

	// map[bool]int32{true: int32(10), false: string("Oke")}

	var listData map[bool]int32 = map[bool]int32{true: 100, false: -777}

	fmt.Println(listData["coki" == "erhan"])

	var llbro map[bool](interface{}) = map[bool]interface{}{true: int32(100), false: errors.New("Error Beyd")}
	fmt.Println(llbro["udin" == "jono"])

	var olama map[bool]interface{} = map[bool]interface{}{true: int32(100), false: errors.New("Error Beyd")}
	fmt.Println(olama["udin" == "jono"])

	var benefits_teamcenter = []string{
		"Increase Deliver To Market",
		"Streamline development process",
	}

	fmt.Println(benefits_teamcenter)

	type Employee struct {
		firstname string `json:"Firstname"`
		lastname  string `json:"Lastname"`
	}

	var emp2 Employee

}
