package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	valores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&valores[i])
	}

	if N == 1 {
		fmt.Println(1)
		return
	}

	maxLen := 1  
	currentLen := 1

	for i := 1; i < N; i++ {
		if valores[i] > valores[i-1] {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 1
		}
	}

	fmt.Println(maxLen)
}
