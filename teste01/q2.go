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

	if numero >= 20 && numero <= 90 {
		fmt.Println("O número está compreendido entre 20 e 90.")
	} else {
		fmt.Println("O número não está compreendido entre 20 e 90.")
	}
}