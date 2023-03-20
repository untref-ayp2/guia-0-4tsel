package main

import (
	"fmt"
	"strings"
)

func main() {
	pol := []float64{1.0, 2.0, 3.0, 4.0}

	polinomio(pol)
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
