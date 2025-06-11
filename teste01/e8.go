package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	valores := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&valores[i])
	}

	distintos := make(map[int]bool)
	for _, v := range valores {
		distintos[v] = true
	}

	fmt.Println(len(distintos))
}
