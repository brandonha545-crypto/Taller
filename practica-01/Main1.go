package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	const descripcion = "localhost"
	const identificadorString = "5432"
	const edadString = "25"
	const inscritoString = "true"
	const timeOutDurationString = "5000"

	identificador, _ := strconv.Atoi(identificadorString)
	edad, _ := strconv.Atoi(edadString)
	inscrito, _ := strconv.ParseBool(inscritoString)
	timeOutMs, _ := strconv.Atoi(timeOutDurationString)
	timeOutDuration := time.Duration(timeOutMs) * time.Millisecond

	fmt.Printf("descripcion, tipo de dato: %T Valor: %v\n", descripcion, descripcion)
	fmt.Printf("identificador, tipo de dato: %T Valor: %d\n", identificador, identificador)
	fmt.Printf("edad, tipo de dato: %T Valor: %d\n", edad, edad)
	fmt.Printf("inscrito, tipo de dato: %T Valor: %v\n", inscrito, inscrito)
	fmt.Printf("timeOutDuration, tipo de dato: %T Valor: %v\n", timeOutDuration, timeOutDuration)
}
