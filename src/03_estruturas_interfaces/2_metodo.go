package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// Metodo "area" possui um tipo "retangulo"
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r *retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {

	r := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimetro: ", r.perimetro())
}