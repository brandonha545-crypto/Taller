package validador

type Registro struct {
	Nombre string
	Edad   int
}

func Validar(r Registro) []error {
	var errores []error

	if r.Nombre == "" {
		errores = append(errores, errorString("nombre vacío"))
	}

	if r.Edad < 18 {
		errores = append(errores, errorString("edad inválida"))
	}

	return errores
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}

var limites = map[string]int{}

func init() {
	limites["nombre"] = 1
	limites["edad"] = 18
}

func Limite(campo string) (int, bool) {
	valor, existe := limites[campo]
	return valor, existe
}
