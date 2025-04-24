package main
import "fmt"

func main() {
  var ano int

	fmt.Scan(&ano)

	if ano <= 0 {
		return
	}

	animais := []string{
		"Macaco", "Galo", "Cao", "Porco", "Rato", "Boi", "Tigre", "Coelho", "Dragao", "Serpente", "Cavalo", "Cabra",
	}

	indice := (ano - 2004) % 12
	if indice < 0 {
		indice += 12 
	}
	fmt.Println(animais[indice])
}
