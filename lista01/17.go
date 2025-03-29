package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanf("%d %d", &x, &y)

	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	} else {
		for i := 0; i < y; i++ {
			fmt.Print(x + i*2)

			if i < y-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
