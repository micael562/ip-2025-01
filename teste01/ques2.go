package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var pi float64 = 0.0
	sign := 1.0

	for i := 0; i < N; i++ {
		term := 4.0 / float64(2*i+1)
		pi += sign * term
		sign *= -1 // alterna o sinal
	}

	fmt.Printf("%.6f\n", pi)
}
