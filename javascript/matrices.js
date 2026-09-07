let matriz = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
];
console.log("Matriz inicial:");
console.log(matriz);

console.log("\n Imprimir la matriz en forma de tabla");
for (let i = 0; i < matriz.length; i++) {
    let fila = "";
    for (let j = 0; j < matriz[i].length; j++) {
        fila += matriz[i][j] + "\t";
    }
    console.log(fila);
}

console.log("\n Recorrer por columnas");
const filas = matriz.length;
const columnas = matriz[0].length;
for (let col = 0; col < columnas; col++) {
    let columna = "";
    for (let fila = 0; fila < filas; fila++) {
        columna += matriz[fila][col] + "\t";
    }
    console.log(columna);
}


console.log("\n Sumar todos los elementos");
let suma = 0;
for (let i = 0; i < matriz.length; i++) {
    for (let j = 0; j < matriz[i].length; j++) {
        suma += matriz[i][j];
    }
}
console.log("Suma total:", suma);

console.log("\n Intercambiar la primera fila con la última");
const ultima = matriz.length - 1;
[matriz[0], matriz[ultima]] = [matriz[ultima], matriz[0]];
console.log("Matriz con filas intercambiadas:");
console.log(matriz);