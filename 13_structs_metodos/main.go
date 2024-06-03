package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado com sucesso \n", c.Nome)
}

func main() {
	bruno := Cliente{"Bruno", 30, true,
		Endereco{"Rua dos Bobos", 0, "São Paulo", "SP"}}
	bruno.Idade = 35
	bruno.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", bruno.Nome, bruno.Idade, bruno.Ativo)
}
