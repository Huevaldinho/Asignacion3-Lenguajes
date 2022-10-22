package main

import (
	"fmt"
	"math"
)

// import "main/main.go" as tree

//======================================Generador Aleatorio=========================================================

func isPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				return false
			}
		}
		return true
	}
}

// Funcion para optener la x
func getX(sem int) int {
	var candidato int
	candidato = sem
	//Se recorren posibles candidatos
	for true {
		//Se encontro un candidatos
		if (isPrime(candidato)) && candidato > sem {
			return candidato
		}
		candidato = candidato + 1
	}
	return -1
}

func generaAleatorio(semilla int, n int, list *[]int) {
	var x, m, a, b int
	x = getX(semilla)
	m = 4096
	a = 109
	b = 853
	for i := 0; i < n; i++ {
		(*list)[i] = x
		x = ((a*(x) + b) % m)
	}
}

func convertidor(list *[]int) {
	for i := 0; i < len(*list); i++ {
		(*list)[i] = (*list)[i] % 255
	}
}

func generaAleatorioAux(semilla, n int) []int {
	var array [5000]int                 //Valor maximo del array.
	var slice = array[0:n]              //Se devuelve un slice con el valor solicitado
	generaAleatorio(semilla, n, &slice) //slice se convierte en la lista de valores generados
	convertidor(&slice)                 //Los valores en slice se traducen al intervalo 0-255
	return slice
}

func imprimir(list []int) {
	//Para imprimir una lista
	for i := 0; i < len(list); i++ {
		fmt.Println("Elemento: ", i, " -> ", (list)[i])
	}

}

//======================================Arbol Binario=========================================================

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
Metodo auxiliar para buscar una llave en un arbol.
Parametros:

	-llave int: Llave del nodo que se esta buscando en el arbol.

Retorna:

	-bool: True si encuentra el nodo con la llave buscada
	| False si no la encuentra
	-int: Cantidad de comparaciones realizadas en la busqueda de la llave.
*/
func (arbol *Arbol) buscarConPointerReceiver(llaveIn int) (bool, int) {
	return arbol.buscarAlternativo(llaveIn, 0)
}

/*
Metodo para buscar una llave en un arbol.
Parametros:

	-llave int: Llave del nodo que se esta buscando en el arbol.
	-comparaciones int: Contador de cuantas comparaciones hace hasta encontrar o terminar de recorrer el
	arbol en la busqueda de la llaveIn.

Retorna:

	-bool: True si encuentra el nodo con la llave buscada
	| False si no la encuentra
	-int: Cantidad de comparaciones realizadas en la busqueda de la llave.
*/
func (arbol *Arbol) buscarAlternativo(llave int, comparaciones int) (bool, int) {
	//La raiz del arbol es nula
	if arbol.raiz == nil {
		return false, comparaciones
	} else if arbol.raiz.llave == llave { //La el nodo raiz tiene la llave buscada.
		return true, comparaciones
	} else if llave > arbol.raiz.llave { // Si la llave es mayor que el nodo raiz, llama recursivamente al hijo derecho.
		return buscar(llave, arbol.raiz.hijoDerecho, comparaciones+1)
	} else { //Si la llave es menor que el nodo raiz, llama recursivamente al hijo izquierdo
		return buscar(llave, arbol.raiz.hijoIzquierdo, comparaciones+1)
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

//======================================Balance de Arbol=========================================================

//http://www.smunlisted.com/day-stout-warren-dsw-algorithm.html
// https://csactor.blogspot.com/2018/08/dsw-day-stout-warren-algorithm-dsw.html

func rotateR(rama *Nodo) bool {
	if &rama == nil {
		return false
	}
	if rama.hijoIzquierdo == nil {
		return false
	}
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

func rotateL(rama *Nodo) bool {
	if rama == nil {
		return false
	}
	if rama.hijoDerecho == nil {
		return false
	}
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

func balanceTree(arbol *Arbol) {
	var tmp *Nodo = arbol.raiz
	var cant int = 0
	//Pasa todos los nodos al lado derecho
	for tmp != nil {
		for tmp.hijoIzquierdo != nil {
			rotateR(tmp)
		}
		cant += 1
		tmp = tmp.hijoDerecho
	}
	var nodosPerf = (int(math.Pow(2.0, math.Ceil(math.Log2(float64(cant)))-1.0)) - 1) //Numero de nodos del arbol perfectamente balanceado mas cercano
	var ultimos int = cant - nodosPerf                                                //Numero de hijos esperados en el ultimo nivel

	fmt.Println(nodosPerf)
	fmt.Println(cant)
	fmt.Println(ultimos)
	tmp = arbol.raiz
	for i := 0; i < ultimos; i++ {
		if i == 0 {
			rotateL(tmp)
			tmp = arbol.raiz
		} else {
			rotateL(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}
	// print(arbol.raiz, 0, "R")
	// var veces int = int(nodosPerf / 2) //Numero de veces en que se rota
	fmt.Println(nodosPerf)
	for nodosPerf > 1 {
		nodosPerf /= 2
		fmt.Println(nodosPerf)
		rotateL(arbol.raiz)
		tmp = arbol.raiz
		for i := 0; i < nodosPerf; i++ {
			rotateL(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}
}

func print(nodo *Nodo, cant int, nomb string) {
	if nodo == nil {
		return
	}
	var espacios string = ""
	//Espacios por estetica
	for i := 0; i < cant; i++ {
		espacios += " "
	}
	fmt.Println(espacios, nomb, ": Llave =>", nodo.llave, " Contador =>", nodo.contador)
	print(nodo.hijoIzquierdo, cant+2, "I")
	print(nodo.hijoDerecho, cant+2, "D")
}

func main() {
	var llave = 10 //Llave a insertar
	//Puntero a arbol vacio.
	var arbol *Arbol = &Arbol{}
	//En algunas funciones se usa el (puntero) nombreFuncion (parametros)
	//Eso se llama Pointer Receivers, sirve para modificar el valor del puntero

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

}
