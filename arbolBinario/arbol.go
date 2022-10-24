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
		return 1 + Max(Altura(nodo.hijoIzquierdo), Altura(nodo.hijoDerecho))
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
func Max(a int, b int) int {
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
	altura := float64(Altura(nodo))
	if altura != 0 {
		return float64(ContarNodos(nodo)) / altura //Retorna la densidad del arbol.
	} else {
		fmt.Println("Error al calcular la densiadad del arbol, altura del arbol es 0, por lo tanto hay una division entre 0.")
		return 0
	}
}

/*
La altura de un nodo en un arbol se define como la longitud del camino
más largo que comienza en el nodo y termina en una hoja. La altura de un
nodo hoja será de cero, y la altura de un nodo se puede calcular sumando
uno a la mayor altura de sus hijos.
La altura de un árbol se define como la altura de su raiz.

La profundidad de un nodo se define como la longitud del camino (único)
que comienza en la raiz y termina en el nodo. La profundidad de la raiz es
cero, y la profundidad de un nodo se puede calcular como la profundidad
de su padre mas uno.



The average depth is the sum of all the node depths divided
by n.

*/
