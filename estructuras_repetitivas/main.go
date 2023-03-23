package main

import "fmt"

func main() {

	fmt.Println(factorial(6))

	fmt.Println(producto(5, 5))

	fmt.Println(es_primo())
}

/*
Escribir una función que recibe un número entero n, no negativo, y devuelve su factorial.
*/
func factorial(n int) int {

	//Guardo el valor de n en la variable resultado como valor inicial
	resultado := n

	for i := n - 1; i > 0; i-- { //Creo un ciclo que vaya desde el parámetro n a 0

		if i != 0 { //Si el valor de i es distinto de 0

			resultado *= i //Multiplico el valor que haya en resultado por i
		}

	}

	return resultado
}

/*
Escribir una función que recibe dos enteros: a y b y devuelve el producto.
Para el cálculo, utilizar sumas sucesivas.
*/
func producto(a, b int) int {

	//Inicializo la variable resultado en 0
	resultado := 0

	for i := 0; i < b; i++ { //Creo un ciclo que vaya de 0 a b

		resultado += a //Sumo el valor actual de resultado con el valor de a
	}

	return resultado
}

/*
Escribir una función que indique si un número entero, ingresado por un usuario, es primo
 (devuelve verdadero o falso).
*/
func es_primo() bool {

	//Inicializo la variable n para guardar lo que entre por Scan
	var n int

	fmt.Scanln(&n) //Guardo en la variable n lo que entre por Scan

	if n == 1 {

		return false //Si n es 1, devuelve false.
	}

	for i := 2; i < n; i++ { //Creo un ciclo que vaya de 2 a n
	
		if n % i == 0 { //Si el resto de la división entre n e i es 0, devuelve false.

			return false
		}
	}

	return true

}
