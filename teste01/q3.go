package main

import (
	"fmt"
)

func main() {
	soma := 0

	fmt.Println("Números de 1 a 100:")
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
		soma += i
	}

	fmt.Println("\nSoma dos números de 1 a 100:", soma)
}