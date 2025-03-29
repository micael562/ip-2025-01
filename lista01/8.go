package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159

	var raio, altura float64

	fmt.Scan(&raio)
	fmt.Scan(&altura)

	areaBase := pi * raio * raio            
	areaLateral := 2 * pi * raio * altura  
	areaTotal := 2*areaBase + areaLateral   

	custo := areaTotal * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
