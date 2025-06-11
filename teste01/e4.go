package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	valores := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&valores[i])
	}

	var S float64
	for i := 0; i < N/2; i++ {
		diff := valores[i] - valores[N-1-i]
		S += math.Pow(diff, 3)
	}

	fmt.Printf("%.2f\n", S)
}
