package main

import ("fmt")

func spiralOrder(matrix [][]int, n, m int) {
	top, bottom := 0, n-1
	left, right := 0, m-1

	for top <= bottom && left <= right {
	
		for i := left; i <= right; i++ {
			fmt.Printf("%d ", matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			fmt.Printf("%d ", matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Printf("%d ", matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Printf("%d ", matrix[i][left])
			}
			left++
		}
	}
	fmt.Println()
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	spiralOrder(matrix, n, m)
}
