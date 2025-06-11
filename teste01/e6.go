package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	valores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&valores[i])
	}

	sort.Ints(valores)

	for i, v := range valores {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
