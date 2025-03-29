package main

import (
	"fmt"
)

func aprovadoOuReprovado(n1, n2, n3 float64) {
	media := (n1 + n2 + n3) / 3
	fmt.Printf("MEDIA = %.2f\n", media)
	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
