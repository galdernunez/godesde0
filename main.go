package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1596)
	fmt.Println(estado)
	fmt.Println(texto)
}
