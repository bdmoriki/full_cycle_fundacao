package main

import "fmt"

type Conta struct {
	Saldo float64
}

func NewConta() *Conta {
	return &Conta{Saldo: 0.00}
}

func (c *Conta) simular(valor float64) float64 {
	c.Saldo += valor
	fmt.Println(c.Saldo)
	return c.Saldo
}

func main() {
	conta := Conta{Saldo: 100.00}
	conta.simular(200)
	fmt.Println(conta.Saldo)
}
