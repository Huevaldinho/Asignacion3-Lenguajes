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
func ObtenerRaiz(arbol * Arbol) * Nodo{
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
