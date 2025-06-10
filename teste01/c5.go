package main

import ("fmt")

func isTriangular(n int) bool {
	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			return true
		}
	}
	return false
}

func main() {
	var num int
	fmt.Scan(&num)

	if isTriangular(num) {
		fmt.Printf("%d eh triangular\n", num)
	} else {
		fmt.Printf("%d nao eh triangular\n", num)
	}
}
