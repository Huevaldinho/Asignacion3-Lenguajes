/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 18/10/2022
Ultima Modificacion: 25/10/2022
*/

/*Algoritmo DSW
Fuente: GRB Dynamics. (2016). SimpleBST [Software]. En GitHub (v0.2).
https://github.com/GRB-Dynamics/SimpleBST/blob/master/src/IntBST.c
Modificado y adaptado para Golang por Maria Fernanda Sanabria Solano
*/

package arbolBinario

import (
	"fmt"
	"math"
)

/*
Funcion rotar la rama hacia la derecha.
Parametro:

	-rama *Nodo: Rama a rotar hacia la derecha.

Retorna:

	-bool: True si se logró rotar, False si no
*/
func RotarD(rama *Nodo) bool {
	if &rama == nil {
		return false //No hay nada que rotar
	}
	if rama.hijoIzquierdo == nil {
		return false //No se puede rotar con nulo
	}
	// Se guarda temporalmente la informacion de la rama
	var tempLlave int = rama.llave
	var tempContador int = rama.contador

	//Intercambiamos los valores de la rama y su hijo izquierdo
	rama.llave = rama.hijoIzquierdo.llave
	rama.contador = rama.hijoIzquierdo.contador
	rama.hijoIzquierdo.llave = tempLlave
	rama.hijoIzquierdo.contador = tempContador

	//Cambio de punteros de la rama y el hijo izquierdo
	var izq *Nodo = rama.hijoIzquierdo
	rama.hijoIzquierdo = izq.hijoIzquierdo
	izq.hijoIzquierdo = izq.hijoDerecho
	izq.hijoDerecho = rama.hijoDerecho
	rama.hijoDerecho = izq
	return true
}

/*
Funcion rotar la rama hacia la izquierda.
Parametro:

	-rama *Nodo: Rama a rotar hacia la izquierda.

Retorna:

	-bool: True si se logró rotar, False si no
*/
func RotarI(rama *Nodo) bool {
	if rama == nil {
		return false //No hay nada que rotar
	}
	if rama.hijoDerecho == nil {
		return false //No se puede rotar con nulo
	}
	// Se guarda temporalmente la informacion de la rama
	var tempLlave int = rama.llave
	var tempContador int = rama.contador

	//Intercambiamos los valores de la rama y su hijo Derecho
	rama.llave = rama.hijoDerecho.llave
	rama.contador = rama.hijoDerecho.contador
	rama.hijoDerecho.llave = tempLlave
	rama.hijoDerecho.contador = tempContador

	//Cambio de punteros de la rama y el hijo Derecho
	var der *Nodo = rama.hijoDerecho
	rama.hijoDerecho = der.hijoDerecho
	der.hijoDerecho = der.hijoIzquierdo
	der.hijoIzquierdo = rama.hijoIzquierdo
	rama.hijoIzquierdo = der
	return true
}

/*
Funcion para balancear un arbol siguiendo el algoritmo DSW.
Parametro:

	-arbol *Arbol: Arbol a balancear.

Retorna:
*/
func Balancear(arbol *Arbol) {
	var tmp *Nodo = arbol.raiz
	var cant int = 0
	//Pasa todos los nodos al lado derecho
	for tmp != nil {
		for tmp.hijoIzquierdo != nil {
			RotarD(tmp)
		}
		cant += 1
		tmp = tmp.hijoDerecho
	}
	var nodosPerf = (int(math.Pow(2.0, math.Ceil(math.Log2(float64(cant)))-1.0)) - 1) //Numero de nodos del arbol perfectamente balanceado mas cercano
	var ultimos int = cant - nodosPerf                                                //Numero de hijos esperados en el ultimo nivel

	//Realiza "ultimos" cantidad de rotaciones a la izquierda
	tmp = arbol.raiz
	for i := 0; i < ultimos; i++ {
		if i == 0 {
			RotarI(tmp)
			tmp = arbol.raiz
		} else {
			RotarI(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}

	//Realiza la cantidad de rotaciones a la izquierda necesaria para terminar con el balance
	for nodosPerf > 1 {
		nodosPerf /= 2

		RotarI(arbol.raiz)
		tmp = arbol.raiz
		for i := 0; i < nodosPerf; i++ {
			RotarI(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}
}

/*
Funcion para imprimir los nodos con sus hijos.
Parametro:

	-Nodo *Nodo: Nodo a imprimir.
	-cant int: Cantidad de espacios que deber'ian colocarse antes
	-nomb string: Posicion del nodo

Retorna:
*/
func Print(nodo *Nodo, cant int, nomb string) {
	if nodo == nil {
		return //No hay nada que imprimir
	}
	var espacios string = ""
	//Espacios por estetica
	for i := 0; i < cant; i++ {
		espacios += " "
	}
	fmt.Println(espacios, nomb, ": Llave =>", nodo.llave, " Contador =>", nodo.contador)
	Print(nodo.hijoIzquierdo, cant+2, "I")
	Print(nodo.hijoDerecho, cant+2, "D")
}
