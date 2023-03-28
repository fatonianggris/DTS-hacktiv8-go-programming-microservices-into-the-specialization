package main

import (
	"fmt"
)

func main() {

	var bilangan int

	var i int

	fmt.Print("Masukan bilangan: ")
	fmt.Scanln(&bilangan)

	for x := 1; x <= bilangan; x++ {
		i = 0
		for y := 1; y <= bilangan; y++ {
			if x%y == 0 {
				i++
				fmt.Println(i)
			}
		}

		if (i == 2) && (x != 1) {
			fmt.Printf("No. %d FizzBuzz \n", x)
		} else {
			fmt.Printf("No. %d \n", x)
		}
	}
}
