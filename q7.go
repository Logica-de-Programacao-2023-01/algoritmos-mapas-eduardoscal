package main

import (
	"fmt"
	"strings"
)

func ContarLetrasPorPalavra(frase string) map[string]map[rune]int {
	palavras := strings.Fields(frase)
	contador := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		contagem := make(map[rune]int)

		for _, letra := range palavra {
			contagem[letra]++
		}

		contador[palavra] = contagem
	}

	return contador
}

func main() {
	frase := "A bola azul é bonita"
	letrasPorPalavra := ContarLetrasPorPal
