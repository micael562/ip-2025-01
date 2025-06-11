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

	linhasOK := true
	for i := 0; i < N; i++ {
		for j := 1; j < M; j++ {
			if matriz[i][j] <= matriz[i][j-1] {
				linhasOK = false
				break
			}
		}
		if !linhasOK {
			break
		}
	}

	colunasOK := true
	for j := 0; j < M; j++ {
		for i := 1; i < N; i++ {
			if matriz[i][j] <= matriz[i-1][j] {
				colunasOK = false
				break
			}
		}
		if !colunasOK {
			break
		}
	}

	if linhasOK && colunasOK {
		fmt.Println("SIM")
	} else {
		fmt.Println("NA0")
	}
}
