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
	}{id: 1, name: "Capucino", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 2, name: "Americano", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 3, name: "Cafelate", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 4, name: "Lemon Tea", price: 9000},
	struct {
		id    int
		name  string
		price int
	}{id: 5, name: "Coconut Lemon", price: 9000},
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
		val := strings.Replace(stringInput, "\n", "", -1)
		ff = string(val)
		fmt.Println("Debug", ff)
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
	var total int32
	for k, v := range list {
		fmt.Println(k+1, v)
		total += int32(v.amount)
	}

	fmt.Printf("Amount: %6d| PPN: %6d|\n", total, 345)
	fmt.Println(limiter)
	return ""
}

// example this form is allow in golang
func diamondOperator[T int32 | string](f func(s string, t T) T) {

}

func main() {

	showListmenu()

	var sInput string

	var iInput int32
	var listItemBuy []itemBuy = make([]itemBuy, 0)

	for {
		fmt.Println("Please input number menu you want to buy")
		iInput := genericInput(inputStream, iInput)
		product := findRepo(int(iInput))
		listItemBuy = append(listItemBuy, itemBuy{
			id:       product.id,
			itemName: product.name,
			qty:      int(iInput),
			amount:   int(iInput) * product.price,
		})

		fmt.Println(">do you want buy more")
		sInput = genericInput(inputStream, sInput)
		if strings.ToUpper(strings.Trim(sInput, " ")) != "Y" {
			billFormater(listItemBuy)
			break
		}
	}

}
