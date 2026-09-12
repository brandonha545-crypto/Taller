package formatter

import "strings"

func Resumen(errs []error) string {

	if len(errs) == 0 {
		return "Registro exitoso"
	}

	var mensajes []string

	for _, err := range errs {
		mensajes = append(mensajes, err.Error())
	}

	return "Error: " + strings.Join(mensajes, ", ")
}
