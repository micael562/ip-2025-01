package main

import "fmt"

func main() {
	var horas, minutos, segundos int

	fmt.Scanf("%d", &horas)
	fmt.Scanf("%d", &minutos)
	fmt.Scanf("%d", &segundos)

	totalSegundos := horas*3600 + minutos*60 + segundos

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos)
}
