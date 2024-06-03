package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	bruno := Cliente{"Bruno", 30, true}
	bruno.Idade = 35

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", bruno.Nome, bruno.Idade, bruno.Ativo)
}
