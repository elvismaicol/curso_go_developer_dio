package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	tempF := ebulicaoF
	tempC := (ebulicaoF - 32) * 5 / 9
	fmt.Printf("A temperatura de ebulição da água em ºF = %g  e a temperatura de ebulição da água em ºC = %g ", tempF, tempC)
	fmt.Printf("A temperatura de ebulição da água em ºF = %g (%T) e a temperatura de ebulição da água em ºC = %g (%T)", tempF, tempC, tempF, tempC)

}
