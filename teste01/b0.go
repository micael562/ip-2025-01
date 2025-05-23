package main

import ("fmt")

func main() {
	var x int
	fmt.Scan(&x)

	if x > 1 {
		fmt.Printf("%d pratos de trigo pra %d tigres tristes", x, x)
	} else {
		return
	}
}
