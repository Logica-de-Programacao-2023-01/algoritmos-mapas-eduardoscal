package main

import "fmt"

func SomarValores(mapa map[string]int) int {
	soma := 0

	for _, valor := range mapa {
		soma += valor
	}

	return soma
}

func main() {
	mapa := map[string]int{"a": 1, "b": 2, "c": 3}

	soma := SomarValores(mapa)
	fmt.Println(soma)
}
