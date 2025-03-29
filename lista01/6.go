package main

import (
	"fmt"
)

func main() {
	var n int
	var fahrenheit float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Scan(&fahrenheit)

		celsius := 5 * (fahrenheit - 32) / 9

		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", fahrenheit, celsius)
	}
}
