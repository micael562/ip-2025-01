package main

import ("fmt")

func main() {
	var nota1, nota2, nota3 float64
	var faltas int

	_, err := fmt.Scan(&nota1, &nota2, &nota3, &faltas)
	if err != nil {
		return
	}

	const totalAulas = 64
	const limiteFaltas = int(0.25 * float64(totalAulas))

	if faltas > limiteFaltas {
		fmt.Println("Reprovado por falta")
		return
	}

	media := (nota1 + nota2 + nota3) / 3

	switch {
	case media >= 7.0:
		fmt.Println("Aprovado")
	case media >= 5.0:
		fmt.Println("Prova final")
	default:
		fmt.Println("Reprovado")
	}
}
