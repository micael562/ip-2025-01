package main

import ("fmt")

func main() {
	var vendasBrutas float64

	_, err := fmt.Scan(&vendasBrutas)
	if err != nil {
		return
	}

	const salarioBase = 500.0
	const comissao = 0.09
	const bonus = 800.0

	salarioFinal := salarioBase + (vendasBrutas * comissao)

	if vendasBrutas > 15000 {
		salarioFinal += bonus
	}

	fmt.Printf("%.2f", salarioFinal)
}
