package main

import "fmt"

func main() {

	status := 200

	switch status / 100 {
	case 2:
		fmt.Println("Éxito")
	case 3:
		fmt.Println("Redirección")
	case 4:
		fmt.Println("Error del cliente")
	case 5:
		fmt.Println("Error del servidor")
	default:
		fmt.Println("No valido")
	}
}
