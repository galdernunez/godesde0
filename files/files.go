package files

import (
	"bufio"
	"fmt"
	"godesde0/ejercicios"
	"io/ioutil"
	"os"
)

var filename string = "./files/txt/archivo.txt"

func GrabaTabla() {
	texto := ejercicios.TablaMultiplicacionStr()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("error al crear el archivo: ", err.Error())
		return
	}
	fmt.Fprint(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.TablaMultiplicacionStr()
	//if append(filename, texto) == false {
	if !append(filename, texto) {
		fmt.Printf("Error al concatenar contenido")
	}
}

func append(filename string, texto string) bool {
	archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error al abrir el archivo: ", err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("error al escribir en el archivo: ", err.Error())
		return false
	}
	archivo.Close()
	return true
}

func LecturaArchivoIoUitls() {
	//ioutil.ReadFile deprecado
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error al leer el archivo: ", err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LecturaArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("error al leer el archivo: ", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
	archivo.Close()
}
