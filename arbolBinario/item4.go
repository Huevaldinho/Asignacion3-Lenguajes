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
