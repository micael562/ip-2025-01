package main

import (
	"fmt"
	"math"
)

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

	normalizada := make([][]float64, N)
	for i := 0; i < N; i++ {
		normalizada[i] = make([]float64, M)

		max := math.Inf(-1)
		for j := 0; j < M; j++ {
			if float64(matriz[i][j]) > max {
				max = float64(matriz[i][j])
			}
		}

		for j := 0; j < M; j++ {
			normalizada[i][j] = float64(matriz[i][j]) / max
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%.6f", normalizada[i][j])
		}
		fmt.Println()
	}
}
