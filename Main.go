package main

import "fmt"

type Cuenta struct {
	titular     string
	saldo       float64
	limite      float64
	movimientos int
}

func (c *Cuenta) Depositar(monto float64) {
	if monto > 0 {
		c.saldo += monto
		c.movimientos++
		fmt.Println("Saldo despues de deposito:", c.saldo)
	} else {
		fmt.Println("No se puede añadir esta cantidad")
	}
}

func (c *Cuenta) Retirar(monto float64) bool {
	if monto > 0 && monto <= c.saldo && monto <= c.limite {
		c.saldo -= monto
		c.movimientos++
		fmt.Println("Saldo despues de retiro:", c.saldo)
		return true
	} else {
		fmt.Println("No es posible retirar esta cantidad")
		return false
	}
}

func (c Cuenta) Informar() string {
	return fmt.Sprintf("Titular: %s\nSaldo: %.2f\nTotal Movimientos: %d",
		c.titular, c.saldo, c.movimientos)
}

type Item struct {
	nombre   string
	precio   float64
	cantidad int
}

func (i Item) Subtotal() float64 {
	return i.precio * float64(i.cantidad)
}

type Carrito struct {
	items     []Item
	descuento float64
}

func (c *Carrito) Agregar(item Item) {
	for i := range c.items {
		if c.items[i].nombre == item.nombre {
			c.items[i].cantidad += item.cantidad
			fmt.Println("Cantidad nueva:", c.items[i].cantidad)
			return
		}
	}

	c.items = append(c.items, item)
	fmt.Println("Item agregado:", item.nombre)
}

func (c *Carrito) Eliminar(nombre string) bool {
	for i := 0; i < len(c.items); i++ {
		if c.items[i].nombre == nombre {
			c.items = append(c.items[:i], c.items[i+1:]...)
			fmt.Println("Item eliminado:", nombre)
			return true
		}
	}

	fmt.Println("Item no encontrado:", nombre)
	return false
}

func (c *Carrito) AplicarDescuento(porcentaje float64) {
	if porcentaje >= 0 && porcentaje <= 100 {
		c.descuento = porcentaje / 100
	}
}

func (c *Carrito) Total() float64 {
	total := 0.0

	for _, item := range c.items {
		total += item.Subtotal()
	}

	total = total - (total * c.descuento)

	return total
}

func (c *Carrito) Resumen() string {
	resumen := ""

	for _, item := range c.items {
		resumen += fmt.Sprintf("%s - $%.2f x %d\n",
			item.nombre, item.precio, item.cantidad)
	}

	resumen += fmt.Sprintf("Total: $%.2f", c.Total())

	return resumen
}

type Config struct {
	host   string
	puerto int
	debug  bool
}

func (c Config) ConDebug() Config {
	c.debug = true
	return c
}

func (c *Config) ActivarDebug() {
	c.debug = true
}

func activarSlice(slice []Config) {
	for i := range slice {
		(&slice[i]).ActivarDebug()
	}
}

func main() {
	cuenta := Cuenta{
		titular: "Ana Torres",
		saldo:   150.00,
		limite:  500,
	}

	cuenta.Depositar(0)
	cuenta.Depositar(-50)
	cuenta.Retirar(1500)

	fmt.Println(cuenta.Informar())

	carrito := Carrito{}

	carrito.Agregar(Item{
		nombre:   "Teclado",
		precio:   45.50,
		cantidad: 1,
	})

	carrito.Agregar(Item{
		nombre:   "Mouse",
		precio:   20.00,
		cantidad: 2,
	})

	carrito.Agregar(Item{
		nombre:   "Teclado",
		precio:   45.50,
		cantidad: 1,
	})

	carrito.AplicarDescuento(10)

	fmt.Println(carrito.Resumen())

	cfg := Config{"localhost", 8080, false}
	cfg2 := cfg.ConDebug()

	fmt.Println("cfg.debug:", cfg.debug)
	fmt.Println("cfg2.debug:", cfg2.debug)

	cfg.ActivarDebug()

	fmt.Println("cfg.debug después de ActivarDebug:", cfg.debug)

	configs := []Config{
		{"localhost", 8080, false},
		{"servidor", 3000, false},
	}

	configs[0].ConDebug()

	fmt.Println("configs[0].debug:", configs[0].debug)

	activarSlice(configs)

	fmt.Println("configs[0].debug despues de ActivarDebug:", configs[0].debug)
}
