/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion:

*/

package arbolBinario

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
func (arbol *Arbol) Insertar(llaveIn int) int {
	//Si la raiz del arbol es vacia, se inserta de una vez.
	var comparaciones = 0 //Variable para contar las comparaciones.
	if arbol.raiz == nil {
		arbol.raiz = &Nodo{llave: llaveIn, contador: 1}
		return comparaciones + 1
	}
	//El arbol si tiene nodos
	var nodoActual = arbol.raiz //Se crea nodo auxiliar para bajar por el arbol.
	for {                       //Ciclo se repite mientas no se llegue a una hoja o encuentre una llave igual.
		comparaciones = comparaciones + 1 //Si entra en el ciclo va a realizar una comparacion.
		if llaveIn == nodoActual.llave {
			nodoActual.contador = nodoActual.contador + 1 //Si el nodo actual tiene la llave buscada, aumenta el contador + 1 .
			break
		} else if llaveIn < nodoActual.llave && nodoActual.hijoIzquierdo == nil { // Llave buscada es menor que la llave del nodo actual.
			//La llave buscada es menor y el hijo izquierdo del nodo actual es nil
			nodoActual.hijoIzquierdo = &Nodo{llave: llaveIn, contador: 1} //Crea nuevo nodo
			break
		} else if llaveIn < nodoActual.llave && nodoActual.hijoIzquierdo != nil {
			//Hijo izquierdo del nodo actual no es nulo, se tiene que volver a iterar bajando por el hijo izquierdo del nodo actual.
			nodoActual = nodoActual.hijoIzquierdo
		} else if llaveIn > nodoActual.llave && nodoActual.hijoDerecho == nil { //Llave buscada es mayor que la llave del nodo actual
			//Si el hijo derecho del nodo actual es nil, se puede insertar de una vez.
			nodoActual.hijoDerecho = &Nodo{llave: llaveIn, contador: 1} //Crea nuevo nodo
			break
		} else {
			//Hijo derecho del nodo actual no es nulo, se tiene que volver a iterar bajando por el hijo derecho del nodo actual.
			nodoActual = nodoActual.hijoDerecho
		}
	}
	comparaciones = comparaciones + 1
	return comparaciones
}
