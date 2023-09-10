package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é Par")
		} else {
			fmt.Println(i, "é Impar")
		}
	}
}

// Verifica se o número é impar
