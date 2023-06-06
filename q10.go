package main

import "fmt"

func ContarPares(slice []int) map[[2]int]int {
	contador := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			contador[pair]++
		}
	}

	return contador
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	frequenciaPares := ContarPares(slice)
	fmt.Println(frequenciaPares)
}
