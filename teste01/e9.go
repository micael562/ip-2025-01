package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	original := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&original[i])
	}

	somaCumulativa := make([]int, N)
	soma := 0

	for i := 0; i < N; i++ {
		soma += original[i]
		somaCumulativa[i] = soma
	}

	for i, val := range somaCumulativa {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
