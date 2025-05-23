package main

import ("fmt")

func main() {
	var N int

	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}

	if N%2 == 0 {
		fmt.Printf("%d Par", N)
	} else {
		fmt.Printf("%d Impar", N)
	}
}
