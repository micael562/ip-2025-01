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

	resultado := make([][]int, N)
	for i := range resultado {
		resultado[i] = make([]int, M)
	}

	direcoes := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1},  {1, 0}, {1, 1},
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			soma := 0
			for _, dir := range direcoes {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < N && nj >= 0 && nj < M {
					soma += matriz[ni][nj]
				}
			}
			resultado[i][j] = soma
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(resultado[i][j])
		}
		fmt.Println()
	}
}
