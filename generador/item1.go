/*
Creadores:
	Obando Arrieta Felipe de Jesús - 2021035489
    Sanabria Solano María Fernanda - 2021005572
Fecha Creación: 17/10/2022
Ultima Modificacion: 24/10/2022
*/

/*Generador de números pseudoaleatorios
Fuente: Dromey, 'How to solve it by computer'
Algorithm 3.6, 'Generation of pseudo-random numbers'
Modificado por Alonso Garita Granados
Adaptado para Golang por Maria Fernanda Sanabria Solano
*/

package generador

import "fmt"

/*
Funcion para comprobar si un numero es o no primo.
Parametro:

	-num int: Numero a comprobar si es o no primo.

Retorna:

	-bool: True si es primo, False si no
*/
func esPrimo(num int) bool {
	if num == 0 || num == 1 { //No son primos
		return false
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				return false //Se encontro un divisor, no es primo
			}
		}
		return true //No se encontro divisor, si es primo
	}
}

/*
Funcion para obtener la x a partir de una semilla.
Parametro:

	-sem int: Semilla a partir de la cual se obtendra x

Retorna:

	-int: x obtenida a partir de sem
*/
func GetX(sem int) int {
	var candidato int
	candidato = sem
	//Se recorren posibles candidatos
	for true {
		if (esPrimo(candidato)) && candidato > sem {
			return candidato //Se encontro un candidato
		}
		candidato = candidato + 1
	}
	return -1
}

/*
Funcion generador de aleatorios.
Parametro:

	-semilla int: Semilla a partir de la cual se obtendran los numeros
	-n int: Cantidad de numeros a generar
	-list *[]int: Lista en donde se almacenaran los numeros generados

Retorna:
*/
func GeneraAleatorio(semilla int, n int, list *[]int) {
	var x, m, a, b int
	x = GetX(semilla)
	// Parámetros tomados de la implementación de Dromey (1982)
	m = 4096
	a = 109
	b = 853
	for i := 0; i < n; i++ {
		(*list)[i] = x        //Inserta el numero generado en el slice
		x = ((a*(x) + b) % m) //Genera el siguiente numero para insertar
	}
}

/*
Funcion convertidor de numeros a intervalo [0,255].
Parametro:

	-list *[]int: Lista en donde se obtendran los numeros para convertirloa

Retorna:
*/
func Convertidor(list *[]int) {
	for i := 0; i < len(*list); i++ { //Recorre list para ir convirtiendo los numeros
		(*list)[i] = (*list)[i] % 255
	}
}

/*
Funcion generador de aleatorios.
Parametro:

	-semilla int: Semilla a partir de la cual se obtendran los numeros
	-n int: Cantidad de numeros a generar

Retorna:

	-n []int: Retorna la lista de numeros solicitados
*/
func GeneraAleatorioAux(semilla, n int) []int {
	if semilla < 11 || semilla > 257 || !esPrimo(semilla) {
		semilla = 11
	}
	if n < 500 || n > 5000 {
		n = 500
	}
	var array [5000]int                 //Valor maximo del array.
	var slice = array[0:n]              //Se devuelve un slice con el valor solicitado
	GeneraAleatorio(semilla, n, &slice) //slice se convierte en la lista de valores generados
	Convertidor(&slice)                 //Los valores en slice se traducen al intervalo 0-255
	return slice
}

/*
Funcion que imprime una lista
Parametro:

	-n []int: Lista a imprimir

Retorna:
*/
func Imprimir(list []int) {
	//Para imprimir una lista
	fmt.Print("[")
	for i := 0; i < len(list)-1; i++ {
		fmt.Print((list)[i], ", ")
	}
	fmt.Print((list)[len(list)-1], " ]")
}

/*
// Main function
func main() {
	fmt.Println("Prueba n = 500, semilla = 11")
	Imprimir(GeneraAleatorioAux(257, 5000))
}
*/
