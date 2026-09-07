package main

import (
	"fmt"
	"math/rand"
)

func busquedaLineal(arr [10]int, objetivo int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == objetivo {
			return i
		}
	}
	return -1
}

func main() {
	var arreglo [10]int
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = rand.Intn(100) + 1
	}
	fmt.Println("Arreglo inicial:", arreglo)

	fmt.Println("\n Recorrido con for clasico")
	for i := 0; i < len(arreglo); i++ {
		fmt.Printf("Posicion %d: %d\n", i, arreglo[i])
	}

	fmt.Println("\n Recorrido con for-each")
	for _, valor := range arreglo {
		fmt.Println(valor)
	}

	fmt.Println("\n Cambiar valores impares por cero")
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i]%2 != 0 {
			arreglo[i] = 0
		}
	}
	fmt.Println("Arreglo con impares en cero:", arreglo)

	fmt.Println("\n Multiplicar cada valor por su Indice")
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = arreglo[i] * i
	}
	fmt.Println("Arreglo multiplicado por Indice:", arreglo)

	valorBuscado := arreglo[3]
	posicion := busquedaLineal(arreglo, valorBuscado)
	fmt.Printf("\nBusqueda lineal del valor %d: encontrado en posicion %d\n", valorBuscado, posicion)

}
