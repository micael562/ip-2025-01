package main

import (
  "fmt"
	"math")

func main() {
	var X float64
	fmt.Scan(&X) 

	S := 0.0
	sinal := 1.0 

	for n := 0; n < 20; n++ {
		termo := math.Pow(X, float64(n)) / factorial(float64(n))
		S += sinal * termo
		sinal *= -1 
	}

	fmt.Printf("%.4f\n", S) 
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	result := 1.0
	for i := 1.0; i <= n; i++ {
		result *= i
	}
	return result
}
