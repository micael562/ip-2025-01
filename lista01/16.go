package main

import "fmt"

func main() {
	var salario float64

	fmt.Scanf("%f", &salario)

	var salarioReajustado float64
	if salario <= 300.00 {
		salarioReajustado = salario * 1.50 
	} else {
		salarioReajustado = salario * 1.30 
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
}
