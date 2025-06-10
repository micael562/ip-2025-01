package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N) 

	var aprovados, exame, reprovados int
	var somaMedias float64

	for i := 1; i <= N; i++ {
		var nota1, nota2 float64
		fmt.Scan(&nota1, &nota2)

		media := (nota1 + nota2) / 2
		somaMedias += media

		var situacao string
		if media <= 3.0 {
			situacao = "Reprovado"
			reprovados++
		} else if media < 7.0 {
			situacao = "Exame"
			exame++
		} else {
			situacao = "Aprovado"
			aprovados++
		}

		fmt.Printf("Aluno %d: %s\n", i, situacao)
	}

	fmt.Printf("\nTotal Aprovados: %d\n", aprovados)
	fmt.Printf("Total Exame: %d\n", exame)
	fmt.Printf("Total Reprovados: %d\n", reprovados)
}
