package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)

	s := 0.0
	pot := 1.0 
	fat := 1.0  
	sinal := 1.0

	for i := 0; i < 20; i++ {
		if i > 0 {
			pot *= x
			fat *= float64(i)
		}
		termo := sinal * pot / fat
		if i == 0 {
			termo = x
		}
		s += termo
		sinal *= -1
	}

	fmt.Printf("%.5f\n", s)
}
