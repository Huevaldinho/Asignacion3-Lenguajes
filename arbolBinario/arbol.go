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
	if nodo == nil {
		return 0
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

/*
	Funcion para calcular la densidad de un arbol.
	Esta esta dada por: cantidad nodos arbol / altura del arbol.
	Parametros:
		-nodo *Nodo: Raiz del arbol.
	Retorna:
		-float64: Densidad del arbol.
*/
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
	Funcion para calcular la profundidad promedio de un arbol.
	Profundidad promedio es la suma de las profundidades de todos
	los nodos del arbol dividida por el total de nodos del arbol.
	Se utiliza en vez de ProfundidadPromedioAux para esconder el parametro
	de la profundidad.

	Parametros:
		-nodo *Nodo: Nodo raiz del arbol.
		-p_promedio float32: Profundidad promedio del arbol, inicia en 0.
	Returna:
		-float32: Profundidad promedio del arbol.
*/
func ProfundidadPromedio(nodo *Nodo) float32 {
	return ProfundidadPromedioAux(nodo, 0) / float32(ContarNodos(nodo))
}

/*
	Idea tomada de: https://iq.opengenus.org/average-height-of-nodes-in-binary-tree/
	Modificado y adaptado a Golang por Felipe Obando Arrieta
	Funcion para calcular la profundidad promedio de un arbol.
	Profundidad promedio es la suma de las profundidades de todos
	los nodos del arbol dividida por el total de nodos del arbol.

	Parametros:
		-nodo *Nodo: Nodo raiz del arbol.
		-p_promedio float32: Profundidad promedio del arbol, inicia en 0.
	Returna:
		-float32: Profundidad promedio del arbol.
*/
func ProfundidadPromedioAux(nodo *Nodo, p_promedio float32) float32 {
	//Si el nodo es una hoja, su profundidad es 0+p_promedio.
	var p_promedio_hi float32 = 0 //Profundidad promedio hijo izquierdo. Cambia si no es hoja.
	var p_promedio_hd float32 = 0 //Profundidad promedio hijo derecho. Cambia si no es hoja.
	if nodo.hijoIzquierdo != nil {
		p_promedio_hi = ProfundidadPromedioAux(nodo.hijoIzquierdo, p_promedio+1) //Suma p_promedio+1 y baja por la izquierda.
	}
	if nodo.hijoDerecho != nil {
		p_promedio_hd = ProfundidadPromedioAux(nodo.hijoDerecho, p_promedio+1) //Suma p_promedio+1 y baja por la derecha.
	}
	return float32(p_promedio + p_promedio_hi + p_promedio_hd)
}
