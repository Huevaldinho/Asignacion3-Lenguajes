/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion: 28/10/2022

*/

package main

import (
	//Para usar las cosas del arbol hay que ejecutar el comando: go mod init asignacion3-lenguajes
	"asignacion3-lenguajes/arbolBinario"
	"asignacion3-lenguajes/generador"
	"fmt"
)

func Pruebas() {
	var arbol *arbolBinario.Arbol = &arbolBinario.Arbol{}

	//Valores generados aleatoriamente
	var A []int = generador.GeneraAleatorioAux(11, 5000)
	for i := 0; i < len(A); i++ {
		fmt.Println("Inserta: ", A[i], " - Cantidad Comparaciones: ", arbol.Insertar(A[i]))
	}

	arbolBinario.Balancear(arbol)

	var secuencia []int = generador.GeneraAleatorioAux(11, 100000)

	for i := 0; i < len(secuencia); i++ {
		fmt.Print(arbol.Buscar(secuencia[i]))
		fmt.Print(" - ")
	}
	fmt.Println()
	fmt.Println(" ---------------------------------- ")
	fmt.Println("Altura arbol: ", arbolBinario.Altura(arbolBinario.ObtenerRaiz(arbol)))
	fmt.Println("Cantidad de nodos: ", arbolBinario.ContarNodos(arbolBinario.ObtenerRaiz(arbol)))
	fmt.Println("Densidad: ", arbolBinario.Densidad(arbolBinario.ObtenerRaiz(arbol)))
	fmt.Println("Profundidad promedio: ", arbolBinario.ProfundidadPromedio(arbolBinario.ObtenerRaiz(arbol)))
}

func main() {
	Pruebas()
}
