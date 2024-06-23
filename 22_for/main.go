package main

func main() {
	// Go tem apenas FOR

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	// for {
	// 	println("Hello World")
	// }
}
