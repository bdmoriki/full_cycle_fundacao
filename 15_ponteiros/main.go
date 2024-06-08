package main

func main() {
	// Memoria -> Endereço -> Valor

	a := 10
	println(a)
	println(&a)

	var ponteiro *int = &a
	*ponteiro = 20

	println(a)

	b := &a

	println(b)
	println(*b)
}
