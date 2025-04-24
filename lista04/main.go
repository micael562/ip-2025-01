package main
import "fmt"

func main() {
  var nota int
	fmt.Scan(&nota)
	if nota < 0 || nota > 10 {
		return
	}
	if nota < 3 {
	    fmt.Printf("E")
	}
	if nota >= 3 && nota < 5 {
	    fmt.Printf("D")
	}
	if nota >= 5 && nota < 7 {
	    fmt.Printf("C")
	}
	if nota >= 7 && nota < 9 {
	    fmt.Printf("B")
	}
	if nota >= 9{
	    fmt.Printf("A")
	}
}
