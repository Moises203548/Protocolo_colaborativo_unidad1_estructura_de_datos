import random


arreglo = [random.randint(1, 100) for i in range(10)]
print("Arreglo inicial:", arreglo)


print("\n Recorrido con for clásico (por índice) ")
for i in range(len(arreglo)):
    print(f"Posición {i} = {arreglo[i]}")


print("\n Recorrido con for-each ")
for i in arreglo:
    print(i)


print("\n Cambiar valores impares por cero ")
for i in range(len(arreglo)):
    if arreglo[i] % 2 != 0:
        arreglo[i] = 0
print("Arreglo con impares en cero:", arreglo)

print("\n Multiplicar cada valor por su índice ")
for i in range(len(arreglo)):
    arreglo[i] = arreglo[i] * i
print("Arreglo multiplicado por índice:", arreglo)


def busqueda_lineal(arr, objetivo):
    for i in range(len(arr)):
        if arr[i] == objetivo:
            return i
    return -1

valor_buscado = arreglo[7] if len(arreglo) > 7 else 0
posicion = busqueda_lineal(arreglo, valor_buscado)
print(f"\nBúsqueda lineal del valor {valor_buscado}: encontrado en posición {posicion}")
