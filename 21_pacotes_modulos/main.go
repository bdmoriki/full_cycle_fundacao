package main

import (
	"fmt"

	"curso-go/matematica"
)

func main() {
	fmt.Printf("Resultado: %v \n", matematica.Soma(10, 30))
	fmt.Println(matematica.A)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())
}
