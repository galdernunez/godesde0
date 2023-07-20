package ejerc_interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.Sexo())
}
