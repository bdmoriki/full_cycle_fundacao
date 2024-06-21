package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Bruno": 10, "João": 20, "Maria": 30}
	m2 := map[string]float64{"Bruno": 10.00, "João": 20.00, "Maria": 30.00}
	m3 := map[string]MyNumber{"Bruno": 10, "João": 20, "Maria": 30}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.0))
}
