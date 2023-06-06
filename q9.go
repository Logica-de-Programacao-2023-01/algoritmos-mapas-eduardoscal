package main

import "fmt"

func GerarSequenciaFibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	a, b := 0, 1

	for i := 0; i <= n; i++ {
		sequencia[i] = a
		a, b = b, a+b
	}

	return sequencia
}

func main() {
	n := 10
	sequenciaFibonacci := GerarSequenciaFibonacci(n)
	fmt.Println(sequenciaFibonacci)
