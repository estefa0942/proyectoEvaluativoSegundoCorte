package vecinoMasCercano

import (
	"fmt"
	"math"
)

type Ruta struct {
	Ruta   [][]int
	Costo  int
}

// vecinoMasCercano encuentra el vecino más cercano no visitado al nodo actual
func vecinoMasCercano(nodoActual int, visitados []bool, costos [][]int) int {
	distanciaMinima := math.MaxInt64
	vecinoCercano := -1

	for nodo := 0; nodo < len(costos); nodo++ {
		if nodo != nodoActual && !visitados[nodo] {
			dist := costos[nodoActual][nodo]
			if dist < distanciaMinima {
				distanciaMinima = dist
				vecinoCercano = nodo
			}
		}
	}
	return vecinoCercano
}

// MetodoDelCartero genera la ruta del vecino más cercano comenzando desde cada nodo
func MetodoDelCartero(nodos []string, costos [][]int) []Ruta {
	var rutas []Ruta

	for inicio := 0; inicio < len(nodos); inicio++ {
		visitados := make([]bool, len(nodos))
		ruta := [][]int{}
		costoTotal := 0
		nodoActual := inicio
		visitados[nodoActual] = true

		for len(ruta) < len(nodos)-1 {
			vecino := vecinoMasCercano(nodoActual, visitados, costos)
			if vecino == -1 {
				break
			}
			costo := costos[nodoActual][vecino]
			ruta = append(ruta, []int{nodoActual, vecino})
			visitados[vecino] = true
			nodoActual = vecino
			costoTotal += costo
		}

		// Agrega el último enlace para volver al punto de partida
		ruta = append(ruta, []int{nodoActual, inicio})
		costoTotal += costos[nodoActual][inicio]

		rutas = append(rutas, Ruta{Ruta: ruta, Costo: costoTotal})
	}

	return rutas
}

// ImprimirRutas imprime todas las rutas generadas y sus costos
func ImprimirRutas(nodos []string, rutas []Ruta) {
	for i, ruta := range rutas {
		fmt.Printf("Ruta %d comenzando desde %s:\n", i+1, nodos[i])
		for _, enlace := range ruta.Ruta {
			fmt.Printf("De %s a %s\n", nodos[enlace[0]], nodos[enlace[1]])
		}
		fmt.Printf("Costo total: %d\n\n", ruta.Costo)
	}
}
