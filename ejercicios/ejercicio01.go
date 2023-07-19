package ejercicios

import (
	"strconv"
)

func Mayor100(numero string) (int, string) {
	var mensaje string
	num, err := strconv.Atoi(numero)
	// _ sirve para despreciar un variable que devuelva la función
	//num, _ := strconv.Atoi(numero)
	if err == nil {
		if num > 100 {
			mensaje = "Es mayor a 100"
		} else {
			mensaje = "Es menor a 100"
		}
	} else {
		mensaje = "Error al convertir string" + err.Error()
	}

	return num, mensaje
}
