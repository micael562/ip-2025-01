package main

import ("fmt")

func isQuadradoPerfeito(n int) bool {
	if n < 1 {
		return false
	}
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func main() {
	var num int
	for {
		fmt.Scan(&num)
		if num <= 0 {
			break
		}
		if isQuadradoPerfeito(num) {
			fmt.Printf("%d eh quadrado perfeito\n", num)
		} else {
			fmt.Printf("%d nao eh quadrado perfeito\n", num)
		}
	}
}
