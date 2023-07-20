package main

import (
	"fmt"
	"godesde0/goroutines"
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
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
	*/
	canal := make(chan bool)
	canal2 := make(chan bool)

	go goroutines.MiNombreLento("Galder", canal)
	go goroutines.MiNombreLento("Galder Nuñez", canal2)
	fmt.Println("Estoy aqui")
	defer func() {
		<-canal
		<-canal2

	}()

}
