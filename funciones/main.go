package main

import (
	"fmt"
	"strings"
)

func main() {
	pol := []float64{3.0, 2.0, 1.0}
	lista := []int{1, 6, 4, 9, -2, 6}
	polinomio(pol)

	fmt.Println(enteros(lista))
}

/*
Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes),
muestre por pantalla el polinomio completo.
Por ejemplo, si recibe los coeficientes 3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2.
*/

func polinomio(coeficiente []float64) {

	//Declaro un arreglo del tipo string para guardar el polinomio
	terminos := []string{}

	//Recorro el arreglo desde el primer elemento hasta el último
	for i := 0; i < len(coeficiente); i++ {

		//Cambio el tipo de dato de float a string, formateando el float con 2 decimales ("%.2f")

		coef := fmt.Sprintf("%.2f", coeficiente[i])

		//Si es el primer elemento, lo guardo sin cambios en el arreglo de términos
		if i == 0 {

			terminos = append(terminos, coef)
		} else if i == 1 { //Si es el segundo término, lo guardo anteponiendo un "+" y agregándole una X al final

			coef = "+ " + coef + " X"

			terminos = append(terminos, coef)
		} else { //Si i es mayor a 1, lo guardo anteponiendo un "+" y agregándole "X**i" en donde i es el valor de la potencia.

			coef = fmt.Sprintf("+ "+coef+" X**%v", i)

			terminos = append(terminos, coef)
		}

	}

	fmt.Println("El polinomio completo es: ")
	fmt.Println(strings.Join(terminos, " ")) //Los muestro en la terminal poniendo un espacio entre elementos.
}

/*
Escribir una función que reciba una lista de enteros y devuelva el menor y el mayor de la lista
*/
func enteros(lista []int) (int, int) {

	//Inicializo las variables min y max con el primer valor de la lista, ya que éste es el menor  y el mayor al momento de iterarla.
	min := lista[0]
	max := lista[0]

	//Recorro el arreglo
	for i := 1; i < len(lista); i++ {
		
		//Si lista en la posición i es menor que el mínimo actual, se actualiza el valor de la variable con el valor actual.
		if lista[i] < min {

			min = lista[i]
		}

		//Si lista en la posición i es mayor que el máximo actual, se actualiza el valor de la variable con el valor actual.
		if lista[i] > max {

			max = lista[i]
		}
	}

	return min, max
}
