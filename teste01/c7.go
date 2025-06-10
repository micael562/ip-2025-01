package main

import (
	"fmt"
)

func main() {
	var X float64
	fmt.Scan(&X) 

	S := 0.0
	fat := 1.0  
	sinal := 1.0 

	for n := 0; n < 20; n++ {
		xn := 1.0
		for i := 0; i < n; i++ {
			xn *= X
		}

		termo := sinal * xn / fat
		S += termo

		sinal *= -1
		if n > 0 {
			fat *= float64(n + 1)
		}
	}

	fmt.Printf("%f\n", S) 
}
