package iteraciones

import (
	"fmt"
)

func IterarPasoaPaso() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func IterarInfinito() {
	for {
		fmt.Println("hola")
	}
}
func IterarInfinitobreak() {
	for {
		fmt.Println("hola")
		break
	}
}

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func IterarSaltos() {
	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}
}

func IterarInversa() {
	for i := 10; i > 1; i-- {
		fmt.Println(i)
	}
}

func Iterarbreak() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func Iterarcontinue() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
