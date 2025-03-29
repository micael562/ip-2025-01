package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var totalPessoas int
		var percPopular, percGeral, percArquibancada, percCadeiras float64

		fmt.Scan(&totalPessoas, &percPopular, &percGeral, &percArquibancada, &percCadeiras)

		populares := float64(totalPessoas) * (percPopular / 100.0)
		gerais := float64(totalPessoas) * (percGeral / 100.0)
		arquibancadas := float64(totalPessoas) * (percArquibancada / 100.0)
		cadeiras := float64(totalPessoas) * (percCadeiras / 100.0)

		renda := (populares * 1.0) + (gerais * 5.0) + (arquibancadas * 10.0) + (cadeiras * 20.0)

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
