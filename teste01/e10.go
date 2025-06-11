package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	valores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&valores[i])
	}

	ehPalindromo := true
	for i := 0; i < N/2; i++ {
		if valores[i] != valores[N-1-i] {
			ehPalindromo = false
			break
		}
	}

	if ehPalindromo {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
