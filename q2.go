package main

import "fmt"

func MesclarMapas(mapa1, mapa2 map[string]int) map[string]int {
	mapaResultado := make(map[string]int)

	for chave, valor := range mapa1 {
		mapaResultado[chave] = valor
	}

	for chave, valor := range mapa2 {
		mapaResultado[chave] = valor
	}

	return mapaResultado
}

func main() {
	mapa1 := map[string]int{"a": 1, "b": 2}
	mapa2 := map[string]int{"b": 3, "c": 4}

	mapaMesclado := MesclarMapas(mapa1, mapa2)
	fmt.Println(mapaMesclado)
}
