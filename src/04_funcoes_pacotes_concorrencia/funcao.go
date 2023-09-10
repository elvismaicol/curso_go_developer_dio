package main

import "fmt"

func media (lista []float64)float64{
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() { // programa que calcula a média de uma sala

	lista := []float64{98, 93, 55, 60, 83} //notas
	fmt.Println(media(lista))
}