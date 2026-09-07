package main

import "fmt"

func main() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matriz inicial:", matriz)

	fmt.Println("\n Imprimir la matriz en forma de tabla")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%d\t", matriz[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\n Recorrer por columnas")
	filas := len(matriz)
	columnas := len(matriz[0])
	for col := 0; col < columnas; col++ {
		for fila := 0; fila < filas; fila++ {
			fmt.Printf("%d\t", matriz[fila][col])
		}
		fmt.Println()
	}

	fmt.Println("\n Sumar todos los elementos")
	suma := 0
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			suma += matriz[i][j]
		}
	}
	fmt.Println("Suma total:", suma)

	fmt.Println("\n Intercambiar la primera fila con la última")
	ultima := len(matriz) - 1
	matriz[0], matriz[ultima] = matriz[ultima], matriz[0]
	fmt.Println("Matriz con filas intercambiadas:", matriz)
}
