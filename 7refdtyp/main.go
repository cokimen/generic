package main

import "fmt"

func main() {

	// zero value of data type below is null
	var dt1 *int = nil
	var dt2 interface{} = nil
	var dt3 []string = nil
	var dt4 chan string = nil
	var dt5 func(s string) int = nil
	var dt6 map[string]string = nil

	fmt.Println(
		dt1,
		dt2,
		dt3,
		dt4,
		dt5,
		dt6,
	)

	var listUser [3]string = [...]string{"erhan", "burhanudin", "cokimen"}
	fmt.Println(listUser)

	var listPair [3][2]string = [...][2]string{{"Erlang", "C"}, {"Haskell", "Java"}, {"Rust", "Fortran"}}

	fmt.Printf("%T %v \n", listPair, listPair)

	// slices
	var listScientist []string = []string{"Sri Ramanujan", "Bertrand Russel", "Alan Turing", "Carl Gauss", "Nikola Tesla", "William Slidis"}

	fmt.Println(
		fmt.Sprintf("variable name listScientist , Data Type: %T, Value: %v", listScientist, listScientist),
	)

	fmt.Println(len(listScientist))
	fmt.Println(cap(listScientist))

	listScientist = append(listScientist, "Euler")

	fmt.Println(len(listScientist))
	fmt.Println(cap(listScientist))

	// preallocate capacity of slice
	var listSurahQuran []string = make([]string, 200)

	listSurahQuran[0] = "Assajadah"
	fmt.Println(len(listSurahQuran))
	fmt.Println(cap(listSurahQuran))

	listSurahQuran = append(listSurahQuran, "Al Baqarah")

	fmt.Println(listSurahQuran[200])
	fmt.Println(len(listSurahQuran))
	fmt.Println(cap(listSurahQuran))

	var object map[string]string = map[string]string{}
	object["username"] = "stduser0"
	object["password"] = "passroot01"

	value, ok := object["location"]

	// fmt.Println(ok)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Ok", ok)
	}

}
