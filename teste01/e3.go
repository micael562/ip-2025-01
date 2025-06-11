package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	original := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&original[i])
	}

	resultado := make([]int, N)

	for i := 0; i < N; i++ {
		switch {
		case N == 1:
			resultado[i] = 0
		case i == 0: 
			resultado[i] = original[i+1]
		case i == N-1: 
			resultado[i] = original[i-1]
		default: 
			resultado[i] = original[i-1] + original[i+1]
		}
	}

	for i, val := range resultado {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
