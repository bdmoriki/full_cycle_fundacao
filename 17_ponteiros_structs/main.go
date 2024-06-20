package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c Cliente) andou() {
	c.Nome = "Bruno D"
	fmt.Printf("O cliente %s andou \n", c.Nome)
}

func main() {
	bruno := Cliente{Nome: "Bruno"}
	bruno.andou()
	fmt.Printf("O valor da struct com nome %v", bruno.Nome)
}
