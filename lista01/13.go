package main

import "fmt"

func main() {
	var nota float64

	fmt.Scan(&nota)

	var conceito string

	if nota >= 9.0 && nota <= 10.0 {
		conceito = "A"
	} else if nota >= 7.5 && nota < 9.0 {
		conceito = "B"
	} else if nota >= 6.0 && nota < 7.5 {
		conceito = "C"
	} else {
		conceito = "D"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
