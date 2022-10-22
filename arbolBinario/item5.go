package arbolBinario

import (
	"fmt"
	"math"
)

//======================================Balance de Arbol=========================================================

//http://www.smunlisted.com/day-stout-warren-dsw-algorithm.html
// https://csactor.blogspot.com/2018/08/dsw-day-stout-warren-algorithm-dsw.html

func RotarD(rama *Nodo) bool {
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

func RotarI(rama *Nodo) bool {
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

func Balancear(arbol *Arbol) {
	var tmp *Nodo = arbol.raiz
	var cant int = 0
	//Pasa todos los nodos al lado derecho
	for tmp != nil {
		for tmp.hijoIzquierdo != nil {
			RotarD(tmp)
		}
		cant += 1
		tmp = tmp.hijoDerecho
	}
	var nodosPerf = (int(math.Pow(2.0, math.Ceil(math.Log2(float64(cant)))-1.0)) - 1) //Numero de nodos del arbol perfectamente balanceado mas cercano
	var ultimos int = cant - nodosPerf                                                //Numero de hijos esperados en el ultimo nivel

	tmp = arbol.raiz
	for i := 0; i < ultimos; i++ {
		if i == 0 {
			RotarI(tmp)
			tmp = arbol.raiz
		} else {
			RotarI(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}

	for nodosPerf > 1 {
		nodosPerf /= 2
		fmt.Println(nodosPerf)
		RotarI(arbol.raiz)
		tmp = arbol.raiz
		for i := 0; i < nodosPerf; i++ {
			RotarI(tmp.hijoDerecho)
			tmp = tmp.hijoDerecho
		}
	}
}

func Print(nodo *Nodo, cant int, nomb string) {
	if nodo == nil {
		return
	}
	var espacios string = ""
	//Espacios por estetica
	for i := 0; i < cant; i++ {
		espacios += " "
	}
	fmt.Println(espacios, nomb, ": Llave =>", nodo.llave, " Contador =>", nodo.contador)
	Print(nodo.hijoIzquierdo, cant+2, "I")
	Print(nodo.hijoDerecho, cant+2, "D")
}
