package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoATexto(1596)
	fmt.Println(estado)
	fmt.Println(texto)*/
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
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
	}
}
