package main

import "fmt"

type Pessoa struct {
	name     string
	lastname string
	flavors  []string
}

func main() {
	pessoa1 := Pessoa{
		"Robson",
		"Silva",
		[]string{"Chocolate", "Morango"},
	}

	pessoa2 := Pessoa{
		"Aline",
		"França",
		[]string{"Milho", "Chocolate Branco"},
	}

	fmt.Println(pessoa1)
	for _, v := range pessoa1.flavors {
		fmt.Println(v)
	}

	fmt.Println("")

	fmt.Println(pessoa2)
	for _, v := range pessoa2.flavors {
		fmt.Println(v)
	}

}
