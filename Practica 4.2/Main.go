package main

import (
	"fmt"
)

var datos = []string{"Juan", "Pedro", "Ana"}

func obtener(id int) string {
	if id < 0 || id >= len(datos) {
		panic("id fuera de rango")
	}

	return datos[id]
}

func seguroObtener(id int) (valor string, err error) {

	defer func() {
		r := recover()

		if r != nil {
			err = fmt.Errorf("recuperado: %v", r)
		}
	}()

	valor = obtener(id)

	return valor, nil
}

func main() {

	valor, err := seguroObtener(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

	valor, err = seguroObtener(5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
}
