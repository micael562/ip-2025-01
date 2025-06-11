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

	transposta := make([][]int, M)
	for j := 0; j < M; j++ {
		transposta[j] = make([]int, N)
		for i := 0; i < N; i++ {
			transposta[j][i] = matriz[i][j]
		}
	}

	for j := 0; j < M; j++ {
		for i := 0; i < N; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(transposta[j][i])
		}
		fmt.Println()
	}
}
