package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	firstname string
}

func (emp Employee) getFirstname() string {
	return emp.firstname
}

type Customer struct {
	custId int
}

func (c Customer) getCustId() int {
	return c.custId
}

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

func serializer(a interface{}) {
	fmt.Println("serializer part")
	switch a.(type) {
	case string:
		fmt.Println("string")
	case int32:
		fmt.Println("int32")
	case byte:
		fmt.Println("byte")
	case float32:
		fmt.Println("float32")
	case struct{}:
		fmt.Println("struct")
	default:
		fmt.Println("didnt defined")
	}
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

	fmt.Println("-----------------------------------------------")

	var emp1 Employee = Employee{firstname: "jaeger"}

	var cust1 Customer = Customer{custId: int(90)}

	var container []interface{} = []interface{}{emp1, cust1}

	for k, v := range container {
		fmt.Println("Index ", k)
		switch v.(type) {
		case Employee:
			fmt.Println(v.(Employee).getFirstname())
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(">>>>")
		case Customer:
			fmt.Println(v.(Customer).getCustId())
			fmt.Println(reflect.TypeOf(v))
			fmt.Println(">>>>")
		default:
			fmt.Println("Unknown")
		}
	}

	fmt.Println("---------------------------------------")

	serializer('a')
	serializer("Cokimen")
	serializer(int32(900))
	serializer(float32(90.99))
}
