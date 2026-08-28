package main

import (
	"fmt"
	"time"
)

func main() {
	var timestamp int64 = 1734998400
	var check rune = '\u2705'

	fecha := time.Unix(timestamp, 0)
	fechaISO := fecha.Format("2006-01-02T15:04:05Z")

	fmt.Printf("%c Conversión exitosa: %s\n", check, fechaISO)
}
