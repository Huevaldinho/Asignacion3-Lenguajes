/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion:

*/

package main

import (
	"fmt"
)

// Estructura para el nodo de los arboles
type Nodo struct {
	llave         int
	contador      int
	hijoIzquierdo *Nodo
	hijoDerecho   *Nodo
}
type Arbol struct {
	raiz *Nodo
}

/*
3. BUSCAR
Metodo para buscar iterativamente una llave en un arbol.
Parametros:

	-llave int: Llave del nodo que se esta buscando en el arbol.

Retorna:

	-bool: True si encuentra el nodo con la llave buscada
	| False si no la encuentra
	-int: Cantidad de comparaciones realizadas en la busqueda de la llave.
*/
func (arbol *Arbol) buscarIterativo(llaveBuscada int) (bool, int) {
	var comparaciones = 0       //Variable para contar las comparaciones
	var nodoActual = arbol.raiz //Se crea nodo para bajar por el arbol
	for {                       //Ciclo se repite mientas no se llegue a una hoja o encuentre el nodo.
		comparaciones = comparaciones + 1
		if nodoActual == nil { //Llego a una hoja, la llaveBuscada no existe
			return false, comparaciones
		} else if llaveBuscada == nodoActual.llave {
			break
		} else if llaveBuscada < nodoActual.llave { // Baja por la izquierda
			nodoActual = nodoActual.hijoIzquierdo
		} else { //Baja por la derecha
			nodoActual = nodoActual.hijoDerecho
		}
	}
	//Si encontro la llave buscada
	return true, comparaciones
}

/*
Funcion auxiliar para esconder el contador de comparaciones en la funcion buscar.

Parametros:

	-llave int: Llave del nodo que se esta buscando en el arbol.
	-raiz * Nodo: Nodo raiz para buscar en el arbol (cambia conforme se hace la recursion).

Retorna:

	-bool: True si encuentra el nodo con la llave buscada
	| False si no la encuentra
	-int: Cantidad de comparaciones realizadas en la busqueda de la llave.
*/
func buscarRecursivo(llaveBuscada int, raiz *Nodo) (bool, int) {
	return buscar(llaveBuscada, raiz, 0)
}

/*
3. BUSCAR alternativo usando recursion
Funcion para buscar una llave en un arbol.
Parametros:

	-llave int: Llave del nodo que se esta buscando en el arbol.
	-raiz * Nodo: Nodo raiz para buscar en el arbol (cambia conforme se hace la recursion).
	-comparaciones: Contador para las comparaciones que se hacen hasta encontrar el nodo.

Retorna:

	-bool: True si encuentra el nodo con la llave buscada
	| False si no la encuentra
	-int: Cantidad de comparaciones realizadas en la busqueda de la llave.
*/
func buscar(llave int, raiz *Nodo, comparaciones int) (bool, int) {
	//La raiz del arbol es nula
	if raiz == nil {
		return false, comparaciones + 1
	} else if raiz.llave == llave { //La el nodo raiz tiene la llave buscada.
		return true, comparaciones + 1
	} else if llave > raiz.llave { // Si la llave es mayor que el nodo raiz, llama recursivamente al hijo derecho.
		return buscar(llave, raiz.hijoDerecho, comparaciones+1)
	} else { //Si la llave es menor que el nodo raiz, llama recursivamente al hijo izquierdo
		return buscar(llave, raiz.hijoIzquierdo, comparaciones+1)
	}
}

/*
4. INSERTAR
Funcion para insertar un nuevo nodo con la llave ingresada.
Si la llave esta repetida, aumenta el contador del nodo con esa llave.

Parametros:

	-llave int: Llave que se quiere insertar
	-nodo *Nodo: Nodo para recorrer el arbol.
	-comparaciones int: Contador de comparaciones que se realizan hasta insertar el nodo.

Retorna: Número entero, que será la cantidad de comparaciones realizadas, incluida la que

	llevó a la inserción
*/
func (arbol *Arbol) insertar(llaveIn int) int {
	//Si la raiz del arbol es vacia, se inserta de una vez.
	var comparaciones = 0 //Variable para contar las comparaciones
	if arbol.raiz == nil {
		arbol.raiz = &Nodo{llave: llaveIn, contador: 1}
		return comparaciones + 1
	}
	//El arbol si tiene nodos
	var nodoActual = arbol.raiz //Se crea nodo para bajar por el arbol
	for {                       //Ciclo se repite mientas no se llegue a una hoja.
		comparaciones = comparaciones + 1
		if llaveIn == nodoActual.llave {
			nodoActual.contador = nodoActual.contador + 1
			break
		} else if llaveIn < nodoActual.llave { // Llave buscada es menor que la llave del nodo actual

			if nodoActual.hijoIzquierdo == nil { //Si el hijo izquierdo del nodo actual es nulo entonces se puede insertar de una vez
				nodoActual.hijoIzquierdo = &Nodo{llave: llaveIn, contador: 1}
				break
			}
			//Hijo izquierdo del nodo actual no es nulo, se tiene que volver a iterar.
			nodoActual = nodoActual.hijoIzquierdo
		} else { //Llave buscada es mayor que la llave del nodo actual
			//Si el hijo derecho del nodo actual es nulo entonces se puede insertar de una vez
			if nodoActual.hijoDerecho == nil {
				nodoActual.hijoDerecho = &Nodo{llave: llaveIn, contador: 1}
				break
			}
			//Hijo derecho del nodo actual no es nulo, se tiene que volver a iterar.
			nodoActual = nodoActual.hijoDerecho
		}
	}
	return comparaciones
}

/*
Funcion para hacer recorrido inOrden de un arbol.
Parametro:

	-raiz *Nodo: Raiz del arbol a recorrer.

Retorna: No retorna nada
*/
func enOrden(raiz *Nodo) {
	if raiz != nil {
		enOrden(raiz.hijoIzquierdo)
		fmt.Print("Llave: ", raiz.llave, " | Contador: ", raiz.contador, "\n")
		enOrden(raiz.hijoDerecho)
	}
}

func main() {
	var llave = 10 //Llave a insertar
	//Puntero a arbol vacio.
	var arbol *Arbol = &Arbol{}

	arbol.insertar(llave)
	llave = 5
	arbol.insertar(llave)
	llave = 7
	arbol.insertar(llave)
	llave = 20
	arbol.insertar(llave)
	llave = 15
	arbol.insertar(llave)
	llave = 30
	arbol.insertar(llave)
	llave = 25
	arbol.insertar(llave)
	llave = 40
	arbol.insertar(llave)
	llave = 23
	arbol.insertar(llave)

	balanceTree(arbol)
	print(arbol.raiz, 0, "R")

	fmt.Println("Buscar iterativo:")
	fmt.Println(arbol.buscarIterativo(1))
	fmt.Println("\nBuscar recursivo:")
	fmt.Println(buscarRecursivo(1, arbol.raiz))

}
