package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	cosSeries := maclaurinCos(x, 20)
	
	cosStd := math.Cos(x)
	
	diff := cosSeries - cosStd
	
	fmt.Printf("%.4f %.4f %.4f\n", cosSeries, cosStd, diff)
}

func maclaurinCos(x float64, n int) float64 {
	result := 0.0
	sign := 1.0
	
	for k := 0; k < n; k++ {
		term := 2 * k

		termValue := sign * math.Pow(x, float64(term)) / factorial(term)
		result += termValue
		sign *= -1 
	}
	
	return result
}

func factorial(n int) float64 {
	if n == 0 || n == 1 {
		return 1.0
	}
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result
}
