package main

import ("fmt")

func main() {
	var X float64
	fmt.Scan(&X)

	S := 0.0
	fat := 1.0 
	sinal := 1.0

	for n := 0; n < 20; n++ {
		if n > 0 {
			fat *= float64(n)
		}
		termo := sinal * pow(X, float64(n)) / fat
		S += termo
		sinal *= -1 
	}

	fmt.Printf("%.5f\n", S)
}

func pow(x float64, n float64) float64 {
	result := 1.0
	for i := 0; i < int(n); i++ {
		result *= x
	}
	return result
}
