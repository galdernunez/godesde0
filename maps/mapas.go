package maps

import "fmt"

func MostrasMapas() {
	// la
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])
	campeonato := map[string]int{
		"Barcelona":     39,
		"Athletic Club": 42,
		"Chivas":        25,
		"Boca Junior":   30,
	}
	fmt.Println(campeonato)
	for equipo, puntuacion := range campeonato {
		fmt.Printf("Equipo: %s, Putuacion %d \n", equipo, puntuacion)
	}

	delete(campeonato, "Chivas")
	for equipo, puntuacion := range campeonato {
		fmt.Printf("Equipo: %s, Putuacion %d \n", equipo, puntuacion)
	}
	puntuaje, existe := campeonato["Juventus"]
	fmt.Printf("La puntuacios capturada es %d y el equipo existe = %t \n", puntuaje, existe)
	puntuaje, existe = campeonato["Athletic Club"]
	fmt.Printf("La puntuacios capturada es %d y el equipo existe = %t \n", puntuaje, existe)
}
