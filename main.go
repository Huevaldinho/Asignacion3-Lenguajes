/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion:

*/

package main

import "fmt"

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
		return false, comparaciones
	} else if raiz.llave == llave { //La el nodo raiz tiene la llave buscada.
		return true, comparaciones
	} else if llave > raiz.llave { // Si la llave es mayor que el nodo raiz, llama recursivamente al hijo derecho.
		return buscar(llave, raiz.hijoDerecho, comparaciones+1)
	} else { //Si la llave es menor que el nodo raiz, llama recursivamente al hijo izquierdo
		return buscar(llave, raiz.hijoIzquierdo, comparaciones+1)
	}
}

/*
	4. INSERTAR
	Funcion para insertar un nuevo nodo con la llave ingresada.
	Debe ser ejecutado con la instancia del arbol a la que se quiere insertar el nuevo nodo.
	Si la llave esta repetida, aumenta el contador del nodo.

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
func inOrder(raiz *Nodo) {
	if raiz != nil {
		inOrder(raiz.hijoIzquierdo)
		fmt.Print("Llave: ", raiz.llave, " | Contador: ", raiz.contador, "\n")
		inOrder(raiz.hijoDerecho)
	}
}

func main() {
	var llave = 10 //Llave a insertar
	//Puntero a arbol vacio.
	var arbol *Arbol = &Arbol{}
	//En algunas funciones se usa el (puntero) nombreFuncion (parametros)
	//Eso se llama Pointer Receivers, sirve para modificar el valor del puntero

	var cantidadComparaciones = arbol.insertar(llave)
	fmt.Print("Cantidad de comparaciones la llave ", llave, "= ", cantidadComparaciones, " \n")
	llave = 5
	cantidadComparaciones = arbol.insertar(llave)
	fmt.Print("Cantidad de comparaciones la llave ", llave, "= ", cantidadComparaciones, " \n")
	llave = 15
	cantidadComparaciones = arbol.insertar(llave)
	fmt.Print("Cantidad de comparaciones la llave ", llave, "= ", cantidadComparaciones, " \n")
	llave = 5
	cantidadComparaciones = arbol.insertar(llave)
	fmt.Print("Cantidad de comparaciones la llave ", llave, "= ", cantidadComparaciones, " \n")

	fmt.Print("\n Inorden arbol: \n")
	inOrder(arbol.raiz)

	fmt.Print("\nBuscar nodo\n")
	fmt.Print(buscar(5, arbol.raiz, 0))
	fmt.Print("\n")
	fmt.Print(buscar(50, arbol.raiz, 0))

	//La funcion buscar debe incluir la compracion del nodo actual? O sea, si llego a una hoja tambien se incluye?

}
