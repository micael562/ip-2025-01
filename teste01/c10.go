package main

import (
	"fmt"
)

func main() {
	var (
		P_inicial, D, DeltaP, P_min float64
		Q_inicial, DeltaQ           int
	)

	fmt.Scan(&P_inicial, &Q_inicial, &D, &DeltaP, &DeltaQ, &P_min)

	maxLucro := -1.0
	precoMax := 0.0
	ingressosMax := 0

	fmt.Printf("Preco    Ingressos Vendidos   Lucro\n")

	iterations := 0
	for P := P_inicial; P >= P_min-1e-9; P -= DeltaP {
		Q := Q_inicial + DeltaQ*iterations
		iterations++

		lucro := (P * float64(Q)) - D

		fmt.Printf("%.2f    %d    %.2f\n", P, Q, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			precoMax = P
			ingressosMax = Q
		}
	}

	fmt.Printf("Lucro maximo: %.2f\n", maxLucro)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.\n", precoMax, ingressosMax)
}
