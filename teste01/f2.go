package main

import ("fmt")

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	matriz := make([][]int, N)
	for i := 0; i < N; i++ {
		matriz[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	somasLinhas := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			somasLinhas[i] += matriz[i][j]
		}
	}

	somasColunas := make([]int, M)
	for j := 0; j < M; j++ {
		for i := 0; i < N; i++ {
			somasColunas[j] += matriz[i][j]
		}
	}

	for i, soma := range somasLinhas {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(soma)
	}
	fmt.Println()

	for j, soma := range somasColunas {
		if j > 0 {
			fmt.Print(" ")
		}
		fmt.Print(soma)
	}
	fmt.Println()
}
