package main

import (
	"fmt"
	"math"
)

func main() {

	var precio1, precio2, precio3 float64 = 19.99, 5.50, 12.00
	var cantidad1, cantidad2, cantidad3 int = 2, 1, 3

	precio1Centavos := int64(math.Round(precio1 * 100))
	precio2Centavos := int64(math.Round(precio2 * 100))
	precio3Centavos := int64(math.Round(precio3 * 100))

	totalConCentavos := precio1Centavos*int64(cantidad1) +
		precio2Centavos*int64(cantidad2) +
		precio3Centavos*int64(cantidad3)

	total := float64(totalConCentavos) / 100
	fmt.Printf("Total: $%.2f\n", total)

	fmt.Printf("Total con IVA (16%%): $%.2f\n", total*1.16)

}
