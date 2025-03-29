package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Print("Digite um número: ")
	if _, err := fmt.Scan(&numero); err != nil {
		fmt.Println("Erro ao ler o número. Certifique-se de inserir um valor válido.")
		return
	}

	if numero > 0 {
		fmt.Println("O número é positivo.")
	} else if numero < 0 {
		fmt.Println("O número é negativo.")
	} else {
		fmt.Println("O número é nulo (zero).")
	}
}
