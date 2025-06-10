package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)

	s := 0.0
	potencia := 1.0  
	fatorial := 1.0  
	sinal := 1.0   

	for i := 0; i < 20; i++ {
		if i == 0 {
			s += x 
			potencia = x
		} else {
			potencia *= x
			fatorial *= float64(i)
			s += sinal * potencia / fatorial
		}
		sinal *= -1
	}

	fmt.Printf("%.5f\n", s)
}
