package main

import ("fmt")

func main() {
	var peso, altura float64

	_, err := fmt.Scan(&peso, &altura)
	if err != nil || altura <= 0 {
		return
	}

	imc := peso / (altura * altura)

	switch {
	case imc < 18.5:
		fmt.Println("Abaixo do peso")
	case imc < 25.0:
		fmt.Println("Peso normal")
	case imc < 30.0:
		fmt.Println("Sobrepeso")
	default:
		fmt.Println("Obesidade")
	}
}
