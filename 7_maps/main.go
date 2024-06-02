package main

import "fmt"

func main() {
	salarios := map[string]float64{"Bruno": 1000.0, "João": 2000.5, "Maria": 3000.0}

	fmt.Println(salarios["Bruno"])

	delete(salarios, "Bruno")
	salarios["Bru"] = 5000

	fmt.Println(salarios["Bru"])

	// sal := make(map[string]float64)
	// sal1 := map[string]float64{}

	for nome, valor := range salarios {
		fmt.Printf("O salario do %s é %v\n", nome, valor)
	}

	for _, valor := range salarios {
		fmt.Printf("O salario é %v\n", valor)
	}
}
