package main

import "fmt"

type Address struct {
	address string
	number  string
	city    string
}

type People struct {
	name    string
	age     int
	address Address
}

func main() {

	address := Address{
		"Rua bósnia",
		"43",
		"Carapicuiba",
	}

	robson := People{
		"Robson",
		28,
		address,
	}

	fmt.Println(robson.address.city)
}
