package main

import ("fmt")

func main() {
	var B, E int
	fmt.Scan(&B, &E) 

	resultado := 1

	if E == 0 {
		fmt.Println(1)
		return
	}

	for i := 1; i <= E; i++ {
		resultado *= B
	}

	fmt.Println(resultado)
}
