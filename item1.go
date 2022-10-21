/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572

Fecha Creación: 17/10/2022
Ultima Modificacion:

*/

package main

import "fmt"

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

// Main function
func main() {
	imprimir(generaAleatorioAux(13, 10))
	// fmt.Println("!... Hello World ...!")
	// fmt.Println(getX(20))
	// fmt.Println(getX(13))
}
