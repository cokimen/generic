package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Customer struct {
	Firstname string `json:"fname"`
	Lastname  string `json:"lname"`
}

type TypeListCustomer []Customer

type Buyer struct {
	id int
}

type TypeListBuyer []Buyer

type Q interface {
	insert() (any, error)
}

type QInherit interface {
	Q
	update() (any, error)
}

func (c *Customer) insert() (any, error) {
	return "Customer insert", nil
}

func (c *Customer) update() (any, error) {
	return "Customer update", nil
}

func (buyer *Buyer) insert() (any, error) {
	return fmt.Sprintf("Buyer insert"), nil
}

func main() {
	var listCust TypeListCustomer = make(TypeListCustomer, 0)

	var data string = `[{
		"fname": "echo",
		"lname": "depok"
	},{
		"fname": "gorim",
		"lname": "rimba"
	}]`

	if err := json.Unmarshal([]byte(data), &listCust); err != nil {
		panic("Error")
	}

	var singleBuyer Buyer = Buyer{id: 10}

	var iface []Q = []Q{&listCust[0], &listCust[1], &singleBuyer}

	log.Println(listCust)

	for k, v := range iface {
		switch v.(type) {
		case *Customer:
			fmt.Println(k, v.(*Customer).Firstname)
		case *Buyer:
			fmt.Println(k, v.(*Buyer).id)
		}

		switch v.(type) {
		case QInherit:
			val1, _ := v.(QInherit).update()
			fmt.Println("<case block 2>", val1)
		case Q:
			val1, _ := v.(Q).insert()
			fmt.Println("<case block 2>", val1)

		}

		switch v.(interface{}) {

		case v.(Q):
			val1, _ := v.(Q).insert()
			fmt.Println("<case block 3>", val1)
		case v.(QInherit):
			val1, _ := v.(QInherit).update()
			fmt.Println("<case block 3>", val1)

		}
		fmt.Println("---------------")
	}

}
