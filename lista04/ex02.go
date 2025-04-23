package main
import "fmt"

func main() {
  var vetor1 [10]int
	var vetor2 [5]int

	// Preenchimento do primeiro vetor
	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	// Preenchimento do segundo vetor
	fmt.Println("Digite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	// Calcular a soma de todos os elementos do segundo vetor
	somaVetor2 := 0
	for _, num := range vetor2 {
		somaVetor2 += num
	}

	// Criar os vetores resultantes
	var paresResultantes []int
	var imparesResultantes []int

	for _, num := range vetor1 {
		if num%2 == 0 {
			paresResultantes = append(paresResultantes, num+somaVetor2)
		} else {
			imparesResultantes = append(imparesResultantes, num+somaVetor2)
		}
	}

	// Exibir resultados
	fmt.Println("Primeiro vetor resultante (pares + soma do segundo vetor):", paresResultantes)
	fmt.Println("Segundo vetor resultante (ímpares + soma do segundo vetor):", imparesResultantes)
}
