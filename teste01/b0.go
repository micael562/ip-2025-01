package main

import ("fmt")

func main() {
	var x int
	fmt.Print("Digite um número inteiro maior que 1: ")
	fmt.Scan(&x)

	if x > 1 {
		fmt.Printf("%d pratos de trigo pra %d tigres tristes\n", x, x)
	} else {
		fmt.Println("O número deve ser maior que 1.")
	}
}
