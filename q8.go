package main

import "fmt"

func CalcularDivisaoGastos(gastos map[string]float64) map[string]float64 {
	totalGastos := 0.0

	for _, valor := range gastos {
		totalGastos += valor
	}

	divisao := make(map[string]float64)

	for nome, valor := range gastos {
		porcentagem := valor / totalGastos
		divisao[nome] = porcentagem - 1.0/float64(len(gastos))
	}

	return divisao
}

func main() {
