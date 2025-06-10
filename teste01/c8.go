package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	x = math.Round(x*10) / 10

	for A := 0.0; A <= x+1e-9; A += 0.1 {

		A = math.Round(A*10) / 10
		sinA := maclaurinSin(A)
		fmt.Printf("%.1f  %.4f\n", A, sinA)
	}
}

func maclaurinSin(A float64) float64 {
	A3 := math.Pow(A, 3) / 6      
	A5 := math.Pow(A, 5) / 120     
	A7 := math.Pow(A, 7) / 5040   
	return A - A3 + A5 - A7
}
