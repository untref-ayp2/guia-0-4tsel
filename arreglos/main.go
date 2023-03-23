package main

import "fmt"

func main(){

	arreglo:=[]int{1, 15, 8, 16, 66, 18}

	fmt.Println(suma(arreglo))
}

/* 
Escribir una función que reciba un arreglo de enteros como parámetros y devuelva la suma de todos sus elementos
*/
func suma(arreglo []int) int{

	resultado := 0

	for i := 0; i < len(arreglo); i++ {
		
		resultado += arreglo[i]
	}

	return resultado
}

