package main

import (
	"errors"
	"fmt"
)

var ErrDivZero = errors.New("división entre cero")
var ErrOverflow = errors.New("overflow")

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

func main() {

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

	err := fmt.Errorf("validación: %w", ErrorValidacion{
		Campo:  "edad",
		Motivo: "inválida",
	})

	var e ErrorValidacion

	if errors.As(err, &e) {
		fmt.Println(e.Campo)
	}
}
