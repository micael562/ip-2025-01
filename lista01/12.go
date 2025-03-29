package main

import "fmt"

func main() {
	var horas int

	fmt.Scan(&horas)

	var valor float64

	if horas <= 3 {
		valor = 10.00
	} else {
		valor = 10.00 + float64(horas-3)*5.00
	}

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
