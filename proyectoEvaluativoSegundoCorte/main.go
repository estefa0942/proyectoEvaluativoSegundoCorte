package main

import (
	"fmt"

	"github.com/estefa0942/proyectoEvaluativoSegundoCorte/insercionMasCercana"
	"github.com/estefa0942/proyectoEvaluativoSegundoCorte/vecinoMasCercano"
)

// 2-opt Swap para mejorar la ruta
func twoOptSwap(tour []int, i, k int) []int {
	newTour := make([]int, len(tour))
	copy(newTour, tour[:i])
	for j := 0; j <= k-i; j++ {
		newTour[i+j] = tour[k-j]
	}
	copy(newTour[k+1:], tour[k+1:])
	return newTour
}

// Función de costo para calcular el costo total de un tour
func calcularCostoTotal(tour []int, matrizAdyacencia [][]int) int {
	costoTotal := 0
	for i := 0; i < len(tour)-1; i++ {
		costoTotal += matrizAdyacencia[tour[i]][tour[i+1]]
	}
	costoTotal += matrizAdyacencia[tour[len(tour)-1]][tour[0]]
	return costoTotal
}

// Implementación del algoritmo de búsqueda de vecindario
func busquedaVecindario(tour []int, matrizAdyacencia [][]int) ([]int, int) {
	mejorTour := make([]int, len(tour))
	copy(mejorTour, tour)
	mejorCosto := calcularCostoTotal(mejorTour, matrizAdyacencia)

	improvement := true
	for improvement {
		improvement = false
		for i := 1; i < len(tour)-1; i++ {
			for k := i + 1; k < len(tour); k++ {
				nuevoTour := twoOptSwap(mejorTour, i, k)
				nuevoCosto := calcularCostoTotal(nuevoTour, matrizAdyacencia)
				if nuevoCosto < mejorCosto {
					copy(mejorTour, nuevoTour)
					mejorCosto = nuevoCosto
					improvement = true
				}
			}
		}
	}
	return mejorTour, mejorCosto
}

func main() {
	// Datos para el método del vecino más cercano
	nodos := []string{"Pereira", "Armenia", "Medellín", "Cartago"}
	costos := [][]int{
		{0, 45, 60, 25},
		{45, 0, 90, 50},
		{60, 90, 0, 70},
		{25, 50, 70, 0},
	}

	rutas := vecinoMasCercano.MetodoDelCartero(nodos, costos)
	vecinoMasCercano.ImprimirRutas(nodos, rutas)

	// Datos para el método de inserción más cercana
	coordenadasPuntos := [][]int{
		{0, 1, 1},
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
		// Agrega más coordenadas según sea necesario
	}
	matrizAdyacencia := insercionMasCercana.ConstruirMatrizAdyacencia(coordenadasPuntos)
	tour, fo := insercionMasCercana.InsercionMasCercana(coordenadasPuntos, matrizAdyacencia)
	fmt.Println("Ruta optimizada usando Inserción Más Cercana:")
	insercionMasCercana.ImprimirTour(coordenadasPuntos, tour, fo)

	// Buscar la mejor ruta combinando ambos resultados
	fmt.Println("Busqueda Vecindario:")
	mejorRuta, mejorCosto := busquedaVecindario(tour, matrizAdyacencia)
	insercionMasCercana.ImprimirTour(coordenadasPuntos, mejorRuta, mejorCosto)
}
