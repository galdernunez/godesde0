package arreglos_slice

import "fmt"

var tablaSlice []int = []int{2, 5, 7, 9}
var arreglo [10]int = [10]int{15, 0, 0, 59, 7, 9, 23, 19}

func MuestraSlice() {
	fmt.Println(tablaSlice)
	tablaSlice = append(tablaSlice, 25)
	fmt.Println(tablaSlice)
	//slice creado desde vector posición 2 al final
	porcion := arreglo[2:]
	//slice creado desde 0 a posición 5
	porcion2 := arreglo[:5]
	//slice creado desde 2 a posición 5
	porcion3 := arreglo[2:5]
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	// En GO crea a 0 el slices en ejecución ira reservando mas memoria tanto como 2 a la n
	nums := make([]int, 0)
	for i := 0; i < 132; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
