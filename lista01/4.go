package main

import (
	"fmt"
)

func main() {
	var salarioMinimo, kwGasto float64

	fmt.Scan(&salarioMinimo, &kwGasto)

	valorKw := (0.7 * salarioMinimo) / 100

	custoConsumo := kwGasto * valorKw

	custoComDesconto := custoConsumo * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", valorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}
