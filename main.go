package main

import (
	"godesde0/goroutines"
	"godesde0/teclado"
)

func main() {
	/*estado, texto := variables.ConviertoATexto(1596)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("esto no es Win")
	} else {
		fmt.Println("esto es Win")
	}

	switch os := runtime.GOARCH; os {
	case "intel":
		fmt.Println("Esto es Intel")
	case "amd":
		fmt.Println("Esto es amd")
	default:
		fmt.Printf("%s \n", os)
	}*/

	/*
			num, text := ejercios.Mayor100("fff")
			fmt.Println(num)
			fmt.Println(text)



		//teclado.IngresoNumeros()

		iteraciones.Iterarbreak()
		iteraciones.Iterarcontinue()
	*/
	//ejercicios.TablaMultiplicacion()
	//files.LecturaArchivo()

	//arreglos_slice.Capacidad()
	//arreglos_slice.MuestroArreglos
	//maps.MostrasMapas()

	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)

	ejerc_interfaces.HumanoRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejerc_interfaces.HumanoRespirando(Maria)
	*/

	go goroutines.MiNombreLento("Galder")
	teclado.IngresoNumeros()
}
