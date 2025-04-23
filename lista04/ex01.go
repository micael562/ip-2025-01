package main
import "fmt"

func main() {
  var vetor [10]int

	// Preenchendo o vetor com entrada do usuário
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	// Verificando números maiores que 50
	encontrou := false
	for i, valor := range vetor {
		if valor > 50 {
			fmt.Printf("Número %d na posição %d é maior que 50\n", valor, i)
			encontrou = true
		}
	}

	if !encontrou {
		fmt.Println("Nenhum número superior a 50 foi encontrado.")
	}
}
