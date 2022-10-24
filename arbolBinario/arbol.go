package arbolBinario

import "fmt"

type Nodo struct {
	llave         int
	contador      int
	hijoIzquierdo *Nodo
	hijoDerecho   *Nodo
}
type Arbol struct {
	raiz *Nodo
}

func ObtenerRaiz(arbol *Arbol) *Nodo {
	return arbol.raiz
}

/*
Funcion para hacer recorrido inOrden de un arbol.
Parametro:

	-raiz *Nodo: Raiz del arbol a recorrer.

Retorna: No retorna nada
*/
func EnOrden(raiz *Nodo) {
	if raiz != nil {
		EnOrden(raiz.hijoIzquierdo)
		fmt.Print("Llave: ", raiz.llave, " | Contador: ", raiz.contador, "\n")
		EnOrden(raiz.hijoDerecho)
	}
}

/*
	Funcion para obtener altura maxima de un arbol.
	Parametros:
		-nodo *Nodo: Nodo raiz del arbol a partir del cual se calcula la altura.
	Retorna:
		-int: Altura del arbol.
*/
func Altura(nodo *Nodo) int {
	//La altura de un arbol vacio es -1.
	if nodo == nil {
		return -1
	} else { //El arbol si tiene nodo raiz.
		//Cuenta el nivel actual + el maximo de nivel entre los hijos.
		return 1 + MaximaAltura(Altura(nodo.hijoIzquierdo), Altura(nodo.hijoDerecho))
	}
}

/*
	Funcion para obtener el mayor el numero mayor entre dos numeros enteros.
	Parametros:
		-a int: Primer numero entero en la compracion.
		-b int: Segundo numero entero en la comparacion.
	Retorna:
		-int: Mayor entre a y b.
*/
func MaximaAltura(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
	Funcion para contar la cantidad de nodos de un arbol.
	Parametros:
		-nodo * Nodo: Nodo inicial del arbol.
	Retorna:
		-int: Cantidad de nodos del arbol.
*/
func ContarNodos(nodo *Nodo) int {
	//Si el nodo es nil entonces no cuenta nada
	//porque es una hoja o el arbol esta vacio.
	if nodo == nil {
		return 0
	} else { //El arbol si tiene nodos
		//Cuenta el nodo actual + nodos hijos
		//bajando por los hijos izquierdo y derecho
		return 1 + ContarNodos(nodo.hijoIzquierdo) + ContarNodos(nodo.hijoDerecho)
	}
}
func Densidad(nodo *Nodo) float64 {
	if nodo == nil {
		return 0
	}
	//Retorna la densidad del arbol.
	//esta esta dada por la cantidad de nodos del arbol/ altura del arbol.
	//utiliza float64 para convertir el int que retorna ContarNodos y Altura
	return float64(ContarNodos(nodo)) / float64(Altura(nodo))
}
