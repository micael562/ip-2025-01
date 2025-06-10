package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)

	s := x            
	potencia := x    
	fatorial := 1.0  
	sinal := -1.0    

	for i := 1; i < 20; i++ {
		fatorial *= float64(i)
		termo := sinal * potencia / fatorial
		s += termo
		potencia *= x     
		sinal *= -1    
	}

	fmt.Printf("%.5f\n", s)
}
