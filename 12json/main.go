package main

import (
	"encoding/json"
	"fmt"
)

type Surat struct {
	Name         string `json:"name"`
	NumberOfAyat int    `json:"number_of_ayat"`
}

func main() {
	// from string to object
	var str1 string = `{
		"name": "Al Baqarah",
		"number_of_ayat": 19
	}`

	var surat1 Surat

	if json.Unmarshal([]byte(str1), &surat1) == nil {
		fmt.Println(
			surat1,
		)
	}

	var listStr2 string = `[
	{
		"name": "Al Mumtahanah",
		"number_of_ayat": 13
	},
	{
		"name": "Al Munafiqun",
		"number_of_ayat": 11
	}
	]`

	var listSurat2 []Surat

	if json.Unmarshal([]byte(listStr2), &listSurat2) == nil {
		fmt.Println(listSurat2)
	}

	fmt.Println("----------------------------------------------------")

	// from object to string
	var listSurat = []Surat{
		{Name: "As Sajadah", NumberOfAyat: 30},
		{Name: "Muzammil", NumberOfAyat: 20},
	}

	if conversion, err := json.MarshalIndent(listSurat, "", " "); err == nil {
		fmt.Println(string(conversion))
	}

}
