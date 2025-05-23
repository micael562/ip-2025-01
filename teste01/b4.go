package main

import ("fmt")

func main() {
	var idade int

	_, err := fmt.Scan(&idade)
	if err != nil {
		return
	}

	switch {
	case idade >= 5 && idade <= 7:
		fmt.Println("Infantil A")
	case idade >= 8 && idade <= 10:
		fmt.Println("Infantil B")
	case idade >= 11 && idade <= 13:
		fmt.Println("Juvenil A")
	case idade >= 14 && idade <= 17:
		fmt.Println("Juvenil B")
	case idade >= 18:
		fmt.Println("Adulto")
	default:
		return
	}
}
