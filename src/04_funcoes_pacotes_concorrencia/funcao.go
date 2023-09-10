package main

import "fmt"

func main() { // programa que calcula a média de uma sala

	lista := []float64{98, 93, 77, 82, 83} //notas

	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	fmt.Println(total / float64(len(lista)))
}