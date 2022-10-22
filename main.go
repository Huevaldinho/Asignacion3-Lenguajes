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
	"fmt"
	//	"asignacion3-lenguajes/generador"
)

func main() {

	var llave = 10 //Llave a insertar
	//Puntero a arbol vacio.
	var arbol *arbolBinario.Arbol = &arbolBinario.Arbol{}

	arbol.Insertar(llave)
	llave = 5
	arbol.Insertar(llave)
	llave = 7
	arbol.Insertar(llave)
	llave = 20
	arbol.Insertar(llave)
	llave = 15
	arbol.Insertar(llave)
	llave = 30
	arbol.Insertar(llave)
	llave = 25
	arbol.Insertar(llave)
	llave = 40
	arbol.Insertar(llave)
	llave = 23
	arbol.Insertar(llave)

	arbolBinario.Balancear(arbol)
	arbolBinario.Print(arbolBinario.ObtenerRaiz(arbol), 0, "R")

	fmt.Println("Buscar iterativo:")
	fmt.Println(arbol.BuscarIterativo(1))
	fmt.Println("\nBuscar recursivo:")
	fmt.Println(arbolBinario.BuscarRecursivo(1, arbolBinario.ObtenerRaiz(arbol)))

}
