package ejercicios

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// usando recursividad
func TablaMultiplicacionRecursivoLllama() {
	var err error
	var num int
	err = errors.New("asdasd")
	for err != nil {
		num, err = recogidaDato()
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(num, " * ", i, " = ", num*i)
	}
}

func recogidaDato() (int, error) {
	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero de la tabla : ")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
	}
	return numero, err
}

func TablaMultiplicacion() {
	var numero int
	var err error
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese numero de la tabla : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d \n", numero, i, numero*i)
	}
}

func TablaMultiplicacionStr() string {
	var numero int
	var err error
	var texto string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese numero de la tabla : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d * %d = %d \n", numero, i, numero*i)
	}
	return texto
}
