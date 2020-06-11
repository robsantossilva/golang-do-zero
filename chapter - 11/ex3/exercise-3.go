package main

import "fmt"

type Veiculo struct {
	doors string
	color string
}

type Caminhonete struct {
	veiculo        Veiculo
	fourWheelDrive bool
}

type Sedan struct {
	veiculo     Veiculo
	luxuryModel bool
}

func main() {
	caminhonete := Caminhonete{
		Veiculo{
			"4 portas",
			"Azul",
		},
		true,
	}

	sedan := Caminhonete{
		Veiculo{
			"2 portas",
			"Branco",
		},
		false,
	}

	fmt.Println(caminhonete)
	fmt.Println(sedan)

	fmt.Println(caminhonete.veiculo.color)
	fmt.Println(sedan.veiculo.doors)
}
