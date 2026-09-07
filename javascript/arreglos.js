let arreglo = Array.from({ length: 10 }, () => Math.floor(Math.random() * 100) + 1);
console.log("Arreglo inicial:", arreglo);

console.log("\n Recorrido con for clasico");
for (let i = 0; i < arreglo.length; i++) {
    console.log(`Posición ${i}: ${arreglo[i]}`);
}

console.log("\n Recorrido con for-each");
arreglo.forEach((valor) => {
    console.log(valor);
});

console.log("\n Cambiar valores impares por cero");
for (let i = 0; i < arreglo.length; i++) {
    if (arreglo[i] % 2 !== 0) {
        arreglo[i] = 0;
    }
}
console.log("Arreglo con impares en cero:", arreglo);

console.log("\n Multiplicar cada valor por su Indice");
for (let i = 0; i < arreglo.length; i++) {
    arreglo[i] = arreglo[i] * i;
}
console.log("Arreglo multiplicado por Indice:", arreglo);

function busquedaLineal(arr, objetivo) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === objetivo) {
            return i;
        }
    }
    return -1;
}

const valorBuscado = arreglo.length > 3 ? arreglo[3] : 0;
const posicion = busquedaLineal(arreglo, valorBuscado);
console.log(`\n Busqueda lineal del valor ${valorBuscado}: encontrado en posicion ${posicion}`);