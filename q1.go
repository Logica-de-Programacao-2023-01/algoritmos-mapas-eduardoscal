package main

import (
	"fmt"
	"strings"
)

func ContarPalavras(texto string) map[string]int {
	palavras := strings.Fields(texto)
	contador := make(map[string]int)

	for _, palavra := range palavras {
		contador[palavra]++
	}

	return contador
}

func main() {
	texto := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	ocorrencias := ContarPalavras(texto)
	fmt.Println(ocorrencias)
}
