package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	fmt.Scan(&X)

	S := 0.0
	factorial := 1.0 

	for n := 0; n < 20; n++ {
		if n > 0 {
			factorial *= float64(n)
		}
		
		term := math.Pow(X, float64(n)) / factorial
		if n%2 == 0 {
			S += term
		} else {
			S -= term
		}
	}

	fmt.Println(S)
}
