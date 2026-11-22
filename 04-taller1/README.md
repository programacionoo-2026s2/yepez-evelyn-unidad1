# 04-taller1 — FizzBuzz con switch sin condición

## Descripción
Programa de consola en Go (module taller1) que recorre los números del 1 al 30 con un bucle for y evalúa cada número usando un switch sin condición (sin sentencias if). Imprime FizzBuzz para múltiplos de 15, Fizz para múltiplos de 3, Buzz para múltiplos de 5 o el número en caso contrario, y al final muestra los contadores de cada palabra.

## Cómo ejecutar
Desde la carpeta 04-taller1, ejecuta el programa con:

go run .

## Resultado

### Salida original (case de 15 antes que 3 y 5)
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
Fizz=8 Buzz=4 FizzBuzz=2

### Experimento del orden (case de 3 antes que case de 15)
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 Fizz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 Fizz
Fizz=10 Buzz=4 FizzBuzz=0

## Preguntas

a) ¿Qué salida y qué conteos obtuviste al poner el case de 3 antes que el de 15, y por qué FizzBuzz dejó de aparecer?
Al poner case i%3 == 0 antes que case i%15 == 0, la salida cambió en los múltiplos de 15 (15 y 30) mostrando Fizz en vez de FizzBuzz, y los conteos quedaron en Fizz=10 Buzz=4 FizzBuzz=0. FizzBuzz dejó de aparecer porque todo múltiplo de 15 también es múltiplo de 3, entonces i%3 == 0 se cumple primero y el switch ejecuta ese caso en lugar de llegar al case de 15.

b) Cuando i vale 15, ¿cuántos case son verdaderos y cuál ejecuta Go? ¿Por qué no hace falta escribir break al final de cada case?
Cuando i vale 15, son verdaderos dos case: i%15 == 0 y i%3 == 0 (i%5 == 0 también da resto 0). En el orden original, Go evalúa de arriba hacia abajo y ejecuta solo el primero verdadero, que es case i%15 == 0. No hace falta escribir break al final de cada case porque en Go el break es implícito en cada rama del switch.