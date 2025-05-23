package main

import ("fmt")

func main() {
	var X, Y, Z int
  
	_, err := fmt.Scan(&X, &Y, &Z)
	if err != nil {
		return
	}

	if X+Y > Z && X+Z > Y && Y+Z > X {
		if X == Y && Y == Z {
			fmt.Println("Equilatero")
		} else if X == Y || X == Z || Y == Z {
			fmt.Println("Isosceles")
		} else {
			fmt.Println("Escaleno")
		}
	} else {
		fmt.Println("Nao forma triangulo")
	}
}
