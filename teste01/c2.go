package main

import ("fmt")

func main() {
	var l, r int
	fmt.Scan(&l, &r) 

	soma := 0
	quantidade := 0

	for i := l; i <= r; i++ {
		if i%2 == 0 {
			soma += i
			quantidade++
		}
	}

	fmt.Println(soma)
	
	if quantidade > 0 {
		media := soma / quantidade 
		fmt.Println(media)
	} else {
		fmt.Println(0)
	}
}
