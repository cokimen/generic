package main

import (
	"fmt"
	"strconv"
)

func getOrders() []string {
	var tmp []string
	var total_items_prices int16
	fmt.Println("What you want to Order")
	fmt.Println("Please write down like below")
	fmt.Println("Udang 9000")
	var item, amount string

	for true {
		fmt.Scanln(&item, &amount)
		var v, _ = strconv.Atoi(amount)
		total_items_prices += int16(v)
		tmp = append(tmp, item+" "+amount)

		fmt.Println("Apakah ingin pesan lagi ")
		var isWantMoreOrder string

		fmt.Scanln(&isWantMoreOrder)

		// clean variable
		item, amount = "", ""
		if isWantMoreOrder == "N" {
			break
		}

	}

	fmt.Println("Total ", total_items_prices)

	return tmp

}

func main() {
	var z []string = getOrders()
	fmt.Println(z)

}
