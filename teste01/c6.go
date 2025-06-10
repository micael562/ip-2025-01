package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	var totalAprovados, totalExame, totalReprovados int
	var somaMedias float64

	for i := 1; i <= N; i++ {
		var nota1, nota2 float64
		fmt.Scan(&nota1, &nota2) 

		media := (nota1 + nota2) / 2
		somaMedias += media

		var situacao string
		if media <= 3 {
			situacao = "Reprovado"
			totalReprovados++
		} else if media < 7 {
			situacao = "Exame"
			totalExame++
		} else {
			situacao = "Aprovado"
			totalAprovados++
		}

		fmt.Printf("Aluno %d: %s\n", i, situacao)
	}

	mediaClasse := somaMedias / float64(N)
	fmt.Printf("\nTotal Aprovados: %d\n", totalAprovados)
	fmt.Printf("Total Exame: %d\n", totalExame)
	fmt.Printf("Total Reprovados: %d\n", totalReprovados)
	fmt.Printf("Media da classe: %.1f\n", mediaClasse)
}
