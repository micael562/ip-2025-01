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

	xmin, xmax := math.Inf(1), math.Inf(-1)
	for _, x := range valores {
		if x < xmin {
			xmin = x
		}
		if x > xmax {
			xmax = x
		}
	}

	normalizados := make([]float64, N)
	for i, x := range valores {
		if xmax == xmin {
			normalizados[i] = 0.00
		} else {
			normalizados[i] = (x - xmin) / (xmax - xmin)
		}
	}

	for i, x := range normalizados {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", x)
	}
	fmt.Println()
}
