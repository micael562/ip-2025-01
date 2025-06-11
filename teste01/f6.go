package main

import ("fmt")

func main() {
	var N, K, M int
	fmt.Scan(&N, &K, &M)

	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, K)
		for j := 0; j < K; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	B := make([][]int, K)
	for i := 0; i < K; i++ {
		B[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	C := make([][]int, N)
	for i := 0; i < N; i++ {
		C[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			sum := 0
			for k := 0; k < K; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(C[i][j])
		}
		fmt.Println()
	}
}
