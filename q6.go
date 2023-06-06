package main

import "fmt"

func SomarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, contagem := range contagens {
		for palavra, quantidade := range contagem {
			soma[palavra] += quantidade
		}
	}

	return soma
}

func main() {
	contagens := []map[string]int{
		{"a": 1, "b": 2},
		{"b": 3, "c": 4},
		{"c": 5, "d": 6},
	}

	somaTotal := SomarContagens(contagens)
	fmt.Println(somaTotal)
}
