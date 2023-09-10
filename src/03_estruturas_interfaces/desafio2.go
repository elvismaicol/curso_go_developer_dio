package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			//fmt.Println(i)
			fmt.Println("Pin e Pan")
			fmt.Println("--**--")
		} else if i%3 == 0 && i%5 != 0 {
			//fmt.Println(i)
			fmt.Println("Pin")
			fmt.Println("--**--")
		} else if i%3 != 0 && i%5 == 0 {
			//fmt.Println(i)
			fmt.Println("Pan")
			fmt.Println("--**--")
		} else {
			continue
		}
	}
}