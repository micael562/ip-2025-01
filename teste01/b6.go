package main

import ("fmt")

func main() {
	var numeroConta int
	var limiteCredito, saldoInicial, totalDepositos, totalRetiradas float64

	_, err := fmt.Scan(&numeroConta, &limiteCredito, &saldoInicial, &totalDepositos, &totalRetiradas)
	if err != nil {
		return
	}

	saldoAtual := saldoInicial + totalDepositos - totalRetiradas

	if saldoAtual > limiteCredito {
		fmt.Printf("Credito excedido: %.2f", saldoAtual)
	} else {
		fmt.Printf("Saldo: %.2f", saldoAtual)
	}
}
