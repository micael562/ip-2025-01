package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	matriz := make([][]int, N)
	for i := 0; i < N; i++ {
		matriz[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	rotacionada := make([][]int, N)
	for i := 0; i < N; i++ {
		rotacionada[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			rotacionada[j][N-1-i] = matriz[i][j]
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rotacionada[i][j])
		}
		fmt.Println()
	}
}
