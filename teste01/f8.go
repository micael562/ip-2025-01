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

	softmax := make([][]float64, N)
	for i := 0; i < N; i++ {
		softmax[i] = make([]float64, M)

		exponentials := make([]float64, M)
		var sum float64
		max := findMax(matriz[i]) 
		
		for j := 0; j < M; j++ {
			exponentials[j] = math.Exp(float64(matriz[i][j] - max))
			sum += exponentials[j]
		}

		for j := 0; j < M; j++ {
			softmax[i][j] = exponentials[j] / sum
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%.6f", softmax[i][j])
		}
		fmt.Println()
	}
}

func findMax(row []int) int {
	max := row[0]
	for _, v := range row {
		if v > max {
			max = v
		}
	}
	return max
}
