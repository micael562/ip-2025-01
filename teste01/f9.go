package main

import ("fmt")

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	F := make([][]int, 3)
	for i := 0; i < 3; i++ {
		F[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&F[i][j])
		}
	}

	C := make([][]int, N-2)
	for i := 0; i < N-2; i++ {
		C[i] = make([]int, M-2)
		for j := 0; j < M-2; j++ {
			sum := 0
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					sum += A[i+1+di][j+1+dj] * F[di+1][dj+1]
				}
			}
			C[i][j] = sum
		}
	}

	for i := 0; i < N-2; i++ {
		for j := 0; j < M-2; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(C[i][j])
		}
		fmt.Println()
	}
}
