package main

import (
	"fmt"
	"sort"
	"strings"
)

func OrdenarPalavra(palavra string) string {
	letras := strings.Split(palavra, "")
	sort.Strings(letras)
	return strings.Join(letras, "")
}

func AgruparAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		chave := OrdenarPalavra(palavra)
		anagramas[chave] = append(anagramas[chave], palavra)
	}

	return anagramas
}

func main() {
	palavras := []string{"amor", "roma", "carro", "arco", "casa"}

	gruposAnagramas := AgruparAnagramas(palavras)
	fmt.Println(gruposAnagramas)
}
