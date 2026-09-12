package main

import (
	"errors"
	"fmt"
	"os"
	"practica4/formatter"
	"practica4/validador"
)

var ErrDivZero = errors.New("división entre cero")
var ErrOverflow = errors.New("overflow")

func cargar() error {
	f, err := os.Open("config.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, createErr := os.Create("config.txt")
			return createErr
		}
		return err
	}
	defer f.Close()
	return nil
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivZero
	}

	r := a / b
	if r > 1e308 {
		return 0, ErrOverflow
	}

	return r, nil
}

type ErrorValidacion struct {
	Campo  string
	Motivo string
}

func (e ErrorValidacion) Error() string {
	return e.Campo + ": " + e.Motivo
}

var datos = []string{"Juan", "Pedro", "Ana"}

func obtener(id int) string {
	if id < 0 || id >= len(datos) {
		panic("id fuera de rango")
	}

	return datos[id]
}

func seguroObtener(id int) (valor string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recuperado: %v", r)
		}
	}()

	valor = obtener(id)
	return valor, nil
}

func ejercicio41() {
	fmt.Println("Ejercicio 4.1")
	casos := []struct {
		a, b float64
	}{
		{10, 2},
		{5, 0},
		{1e308, 0.0001},
	}

	for _, c := range casos {
		r, err := Dividir(c.a, c.b)
		switch {
		case errors.Is(err, ErrDivZero):
			fmt.Println("cero")
		case errors.Is(err, ErrOverflow):
			fmt.Println("overflow")
		case err != nil:
			fmt.Println("inesperado")
		default:
			fmt.Println(r)
		}
	}

	err := fmt.Errorf("validación: %w", ErrorValidacion{Campo: "edad", Motivo: "inválida"})
	var e ErrorValidacion
	if errors.As(err, &e) {
		fmt.Println(e.Campo)
	}
}

func ejercicio42() {
	fmt.Println("Ejercicio 4.2")
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

func ejercicio43() {
	fmt.Println("Ejercicio 4.3")
	registro := validador.Registro{Nombre: "Brandon", Edad: 20}
	fmt.Println(formatter.Resumen(validador.Validar(registro)))
}

func ejercicio44() {
	fmt.Println("Ejercicio 4.4")
	registroInvalido := validador.Registro{Nombre: "", Edad: 17}
	fmt.Println(formatter.Resumen(validador.Validar(registroInvalido)))
}

func main() {
	if err := cargar(); err != nil {
		fmt.Println(err)
		return
	}

	ejercicio41()
	ejercicio42()
	ejercicio43()
	ejercicio44()
}
