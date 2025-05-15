func main() {
	var n int

	fmt.Print("Digite um número entre 1 e 20: ")
	fmt.Scan(&n)

	if n < 1 || n > 20 {
		fmt.Println("Erro: o número deve estar entre 1 e 20.")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", i*j)
		}
	}
}
