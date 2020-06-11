package main

import "fmt"

func main() {
	anonymous := struct {
		newMap   map[int]string
		newSlice []string
	}{
		map[int]string{1: "Palmeiras", 2: "Corinthians"},
		[]string{"Morango", "Chocolate"},
	}

	fmt.Println(anonymous)
}
