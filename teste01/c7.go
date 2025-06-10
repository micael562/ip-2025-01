package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	var s float64 = 0.0
	var sinal float64 = 1.0
	var fatorial float64 = 1.0
	var potencia float64 = 1.0

	for i := 1; i <= 20; i++ {
		potencia *= x
		fatorial *= float64(i)
		sinal *= -1
		s += sinal * potencia / fatorial
	}

	s += x

	fmt.Printf("%.5f\n", s)
}
