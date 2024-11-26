package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	firstname string
	lastname  string
	address   string
}

type Account struct {
	firstname string
	lastname  string
	address   string
}

type Record struct {
	id    int
	fname string
}

func main() {
	var us1 User = User{
		firstname: "Coki",
		lastname:  "Burhanudin",
		address:   "Jakarta",
	}
	fmt.Printf("User%v \n", us1)

	var ac1 Account

	ac1 = Account(us1)

	fmt.Printf("Account%v\n", ac1)

	// var rec1 Record = Record(ac1) selesain ini

	fmt.Println(
		fmt.Sprintf("%v", "Hello Please input your name and age"),
	)

	var z *bufio.Reader = bufio.NewReader(os.Stdin)

	inputResult, err := z.ReadString('\n')

	if err != nil {
		panic("data type error")
	}

	fmt.Println(inputResult)

	inputResult = strings.Replace(inputResult, "\n", "", -1)
	fmt.Println(len(inputResult))

	inputResultInt, err := strconv.ParseInt(inputResult, 0, 64)
	fmt.Printf("Conversion to int -> %v \n", inputResultInt)

}
