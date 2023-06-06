package main

import "fmt"

func ContarCaracteres(texto string) map[rune]int {
	contador := make(map[rune]int)

	for _, caractere := range texto {
		contador[caractere]++
	}

	return contador
}

func main() {
	texto := "Lorem ipsum dolor sit amet"
	frequencia := ContarCaracteres(texto)
	fmt.Println(frequencia)
}
