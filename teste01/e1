package main

import ("fmt")

func main() {
	var N int
	fmt.Scan(&N)

	alturas := make([]float64, N)
	var soma float64

	for i := 0; i < N; i++ {
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media := soma / float64(N)

	for _, altura := range alturas {
		if altura > media {
			fmt.Printf("%.2f\n", altura)
		}
	}
}
