package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type product struct {
	id    int
	name  string
	price int
}

var listProduct []struct {
	id    int
	name  string
	price int
} = []struct {
	id    int
	name  string
	price int
}{
	struct {
		id    int
		name  string
		price int
	}{id: 0, name: "Capucino", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 1, name: "Americano", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 2, name: "Cafelate", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 3, name: "Lemon Tea", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 4, name: "Coconut Lemon", price: 9000},
}

type itemBuy struct {
	id       int
	itemName string
	qty      int
	amount   int
}

type BillInfo struct {
	CustomerName string
	items        []itemBuy
}

var inputStream *bufio.Reader = bufio.NewReader(os.Stdin)

func findRepo(id int) product {
	var item *product
	for _, v := range listProduct {
		if v.id == id {
			item = &product{id: v.id, name: v.name, price: v.price}
		}
	}
	return *item
}

func genericInput[T int32 | string](inpStr *bufio.Reader, retVal T) T {
	stringInput, err := inpStr.ReadString('\n')
	if err != nil {
		panic("Error Stream")
	}

	// var ff interface{} = (interface{})(retVal)
	var ff interface{}

	switch (interface{})(retVal).(type) {
	case int32:
		val, err := strconv.Atoi(strings.Replace(stringInput, "\n", "", -1))
		if err != nil {
			panic("Please type number")
		}
		ff = int32(val)
		retVal = ff.(T)
		// map[bool](interface{}){true: retVal, false: panic("Error COnversion")}[err == nil]
	case string:
		ff = string("sada")
		retVal = ff.(T)
	}
	return retVal

}

func showListmenu() {
	fmt.Println("Welcome to Menu Online")
	fmt.Println("1) Capucino")
	fmt.Println("2) Americano")
	fmt.Println("3) Cafelate")
	fmt.Println("4) Lemon Tea")
	fmt.Println("5) Coconut Lemon")

}

func inputConsole() []itemBuy {

	_, err := inputStream.ReadString('\n')
	if err != nil {
		panic("Error Stream")
	}

	var result []itemBuy

	fmt.Println(listProduct)

	return result

}

func billFormater(list []itemBuy) string {
	var limiter string = strings.Repeat("-", 40)
	fmt.Println(limiter)
	fmt.Printf("%-20s:%19s\n", "Customer Name", "Ucup87899")
	fmt.Printf("%-20s:%19s\n", "Order Date ", "2024-11-11")
	fmt.Printf("%-20s:%19s\n", "Items ", "")
	for k, v := range list {
		fmt.Println(k, v)
	}
	fmt.Printf("Amount: %6d| PPN: %6d|\n", 12, 345)
	fmt.Println(limiter)
	return ""
}

func diamondOperator[T int32 | string](f func(s string, t T) T) {

}

func main() {

	showListmenu()
	fmt.Println("Please input number menu you want to buy")

	var sInput string
	var iInput int32
	genericInput(inputStream, iInput)
	genericInput(inputStream, sInput)

	inputConsole()

	var list1 []itemBuy = []itemBuy{
		itemBuy{
			id:       1,
			itemName: "cok",
			qty:      90,
			amount:   9000,
		},
	}
	billFormater(list1)

}
