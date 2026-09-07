matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
]
print("Matriz inicial:")
print(matriz)


print("\n Imprimir la matriz en forma de tabla")
for fila in matriz:
    for valor in fila:
        print(valor, end="\t")
    print()

print("\n Recorrer por columnas")
filas = len(matriz)
columnas = len(matriz[0])
for col in range(columnas):
    for fila in range(filas):
        print(matriz[fila][col], end="\t")
    print()


print("\n Sumar todos los elementos")
suma = 0
for fila in matriz:
    for valor in fila:
        suma += valor
print("Suma total:", suma)

print("\n Intercambiar la primera fila con la última")
matriz[0], matriz[-1] = matriz[-1], matriz[0]
print("Matriz con filas intercambiadas:")
for fila in matriz:
    print(fila)