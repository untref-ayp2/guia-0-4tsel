package main

import "fmt"

func main() {

	arreglo := []int{7, 5, 9, 1, 6, 4, 8, 3, 1}
	arreglo2 := []int{2, 1, 6, 4, 5, 0, 3, 9}

	fmt.Println(suma(arreglo))

	fmt.Println(union(arreglo, arreglo2))

	fmt.Println(interseccion(arreglo, arreglo2))

}

/*
Escribir una función que reciba un arreglo de enteros como parámetros y devuelva la suma de todos sus elementos
*/
func suma(arreglo []int) int {

	resultado := 0

	for _, v := range arreglo {

		resultado += v
	}

	return resultado
}

/*
Escribir una función que reciba dos arreglos A y B, de N y M elementos respectivamente
que representan conjuntos de enteros y devuelva una arreglo con la unión y otro con la intersección de A y B.
*/
func union(a, b []int) []int {

	unido := a[:]

	unido = append(unido, b...)

	unido = eliminar_repetidos(unido)

	return unido
}

func interseccion(a, b []int) []int {

	intersec := []int{}

	for _, v := range a {

		aux, encontrado := 0, false

		for !encontrado && aux < len(b) {

			if v == b[aux] {

				encontrado = true
			}

			aux++
		}

		if encontrado {

			intersec = append(intersec, v)
		}
	}

	intersec = eliminar_repetidos(intersec)

	return intersec
}

func eliminar_repetidos(a []int) []int {

	var filtrado []int

	filtrado = append(filtrado, a[0])

	for _, v := range a {

		encontrado, aux := false, 0

		for !encontrado && aux < len(filtrado) {

			if filtrado[aux] == v {

				encontrado = true
			}

			aux++
		}

		if !encontrado {

			filtrado = append(filtrado, v)
		}
	}

	return filtrado
}
