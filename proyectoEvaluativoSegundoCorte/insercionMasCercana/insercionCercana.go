
package insercionMasCercana

import (
	"fmt"
	"math"
	"sync"
)

// Message estructura para la comunicación de inserciones
type Message struct {
	Node      int
	Position  int
	Difference int
}

// Funcion para calcular la distancia euclidiana entre dos puntos
func distanciaEuclidiana(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

// Funcion para construir la matriz de adyacencia
func ConstruirMatrizAdyacencia(coordenadasPuntos [][]int) [][]int {
	numeroNodos := len(coordenadasPuntos)
	matrizAdyacencia := make([][]int, numeroNodos)
	for i := range matrizAdyacencia {
		matrizAdyacencia[i] = make([]int, numeroNodos)
		for j := range matrizAdyacencia[i] {
			if i != j {
				matrizAdyacencia[i][j] = distanciaEuclidiana(coordenadasPuntos[i][1], coordenadasPuntos[i][2], coordenadasPuntos[j][1], coordenadasPuntos[j][2])
			}
		}
	}
	return matrizAdyacencia
}

// InsercionMasCercana realiza el algoritmo de inserción más cercana
func InsercionMasCercana(coordenadasPuntos [][]int, matrizAdyacencia [][]int) ([]int, int) {
	numeroNodos := len(coordenadasPuntos)
	nodosCubiertos := make(map[int]bool)
	nodosSinCubrir := make(map[int]bool)
	for i := 0; i < numeroNodos; i++ {
		nodosSinCubrir[i] = true
	}
	tour := []int{0}
	nodosCubiertos[0] = true
	delete(nodosSinCubrir, 0)

	// Inicializar la envolvente con los dos puntos más lejanos del nodo inicial
	posiblesVerticesEnvolvente := []struct {
		Distancia int
		Nodo      int
	}{}
	for i := 0; i < numeroNodos; i++ {
		if matrizAdyacencia[0][i] != 0 {
			posiblesVerticesEnvolvente = append(posiblesVerticesEnvolvente, struct {
				Distancia int
				Nodo      int
			}{matrizAdyacencia[0][i], i})
		}
	}
	posiblesVerticesEnvolvente = append(posiblesVerticesEnvolvente, struct {
		Distancia int
		Nodo      int
	}{matrizAdyacencia[0][1], 1})
	posiblesVerticesEnvolvente = append(posiblesVerticesEnvolvente, struct {
		Distancia int
		Nodo      int
	}{matrizAdyacencia[0][2], 2})
	tour = append(tour, posiblesVerticesEnvolvente[0].Nodo)
	tour = append(tour, posiblesVerticesEnvolvente[1].Nodo)
	nodosCubiertos[posiblesVerticesEnvolvente[0].Nodo] = true
	nodosCubiertos[posiblesVerticesEnvolvente[1].Nodo] = true
	delete(nodosSinCubrir, posiblesVerticesEnvolvente[0].Nodo)
	delete(nodosSinCubrir, posiblesVerticesEnvolvente[1].Nodo)

	for len(nodosSinCubrir) > 0 {
		ch := make(chan Message)
		var wg sync.WaitGroup

		for i := 0; i < len(tour)-1; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				costoOriginalArista := matrizAdyacencia[tour[i]][tour[i+1]]
				for k := range nodosSinCubrir {
					costoInsercionK := matrizAdyacencia[tour[i]][k] + matrizAdyacencia[k][tour[i+1]]
					diferenciaInsercion := costoInsercionK - costoOriginalArista
					ch <- Message{Node: k, Position: i + 1, Difference: diferenciaInsercion}
				}
			}(i)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			costoOriginalArista := matrizAdyacencia[tour[len(tour)-1]][tour[0]]
			for k := range nodosSinCubrir {
				costoInsercionK := matrizAdyacencia[tour[len(tour)-1]][k] + matrizAdyacencia[k][tour[0]]
				diferenciaInsercion := costoInsercionK - costoOriginalArista
				ch <- Message{Node: k, Position: 0, Difference: diferenciaInsercion}
			}
		}()

		go func() {
			wg.Wait()
			close(ch)
		}()

		mejorInsercion := Message{Difference: math.MaxInt32}
		for mensaje := range ch {
			if mensaje.Difference < mejorInsercion.Difference {
				mejorInsercion = mensaje
			}
		}

		tour = append(tour[:mejorInsercion.Position], append([]int{mejorInsercion.Node}, tour[mejorInsercion.Position:]...)...)
		nodosCubiertos[mejorInsercion.Node] = true
		delete(nodosSinCubrir, mejorInsercion.Node)
	}

	// Calcular la función objetivo
	fo := 0
	for i := 0; i < len(tour)-1; i++ {
		fo += matrizAdyacencia[tour[i]][tour[i+1]]
	}
	fo += matrizAdyacencia[tour[len(tour)-1]][tour[0]]

	return tour, fo
}

// ImprimirTour imprime el tour y el costo total
func ImprimirTour(coordenadasPuntos [][]int, tour []int, fo int) {
	fmt.Println("Tour optimizado:")
	for _, nodo := range tour {
		fmt.Printf("Nodo %d -> ", nodo)
	}
	fmt.Printf("Nodo %d\n", tour[0])
	fmt.Printf("Costo total: %d\n", fo)
}
