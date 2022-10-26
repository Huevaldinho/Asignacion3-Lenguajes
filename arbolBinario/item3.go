/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 15/10/2022
Ultima Modificacion:

*/

package arbolBinario

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
func (arbol *Arbol) Buscar(llaveBuscada int) (bool, int) {
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
func Buscar2(llaveBuscada int, raiz *Nodo) (bool, int) {
	return BuscarRecursivo(llaveBuscada, raiz, 0)
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
func BuscarRecursivo(llave int, raiz *Nodo, comparaciones int) (bool, int) {
	//La raiz del arbol es nula
	if raiz == nil {
		return false, comparaciones + 1
	} else if raiz.llave == llave { //La el nodo raiz tiene la llave buscada.
		return true, comparaciones + 1
	} else if llave > raiz.llave { // Si la llave es mayor que el nodo raiz, llama recursivamente al hijo derecho.
		return BuscarRecursivo(llave, raiz.hijoDerecho, comparaciones+1)
	} else { //Si la llave es menor que el nodo raiz, llama recursivamente al hijo izquierdo
		return BuscarRecursivo(llave, raiz.hijoIzquierdo, comparaciones+1)
	}
}
