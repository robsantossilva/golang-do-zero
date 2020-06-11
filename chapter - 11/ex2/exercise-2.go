package main

import "fmt"

type Pessoa struct {
	name     string
	lastname string
	flavors  []string
}

func main() {

	newMap := make(map[string]Pessoa)

	pessoa1 := Pessoa{
		"Robson",
		"Silva",
		[]string{"Chocolate", "Morango"},
	}

	newMap[pessoa1.lastname] = pessoa1

	pessoa2 := Pessoa{
		"Aline",
		"França",
		[]string{"Milho", "Chocolate Branco"},
	}

	newMap[pessoa2.lastname] = pessoa2

	for _, v := range newMap {

		fmt.Println(v.name, v.lastname)

		for _, s := range v.flavors {
			fmt.Println(s)
		}
		fmt.Println("")
	}

}
