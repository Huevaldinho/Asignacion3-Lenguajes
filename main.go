/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion:

*/

package main

import (
	//Para usar las cosas del arbol hay que ejecutar el comando: go mod init asignacion3-lenguajes
	"asignacion3-lenguajes/arbolBinario"
	"asignacion3-lenguajes/generador"
	"fmt"
	//	"asignacion3-lenguajes/generador"
)

func main() {

	var aleatorios []int = generador.GeneraAleatorioAux(13, 10)

	//Puntero a arbol vacio.
	var arbol *arbolBinario.Arbol = &arbolBinario.Arbol{}

	//Inserta valores generados aleatoriamente
	for i := 0; i < len(aleatorios); i++ {
		arbol.Insertar(aleatorios[i])
	}

	arbolBinario.Balancear(arbol)
	arbolBinario.EnOrden(arbolBinario.ObtenerRaiz(arbol))
	arbolBinario.Print(arbolBinario.ObtenerRaiz(arbol), 0, "R")

	fmt.Println("Altura arbol 1: ", arbolBinario.Altura(arbolBinario.ObtenerRaiz(arbol)))
	fmt.Println("Cantidad de nodos: ", arbolBinario.ContarNodos(arbolBinario.ObtenerRaiz(arbol)))
	fmt.Println("Densidad: ", arbolBinario.Densidad(arbolBinario.ObtenerRaiz(arbol)))

}
