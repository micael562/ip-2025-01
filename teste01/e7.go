package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	produto := 0
	for i := 0; i < N; i++ {
		produto += A[i] * B[i]
	}

	fmt.Println(produto)
}
