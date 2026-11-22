## Descripción

Conversor de unidades de almacenamiento (`B`, `KB`, `MB` y `GB` en base 1024) desarrollado en Go para línea de comandos, pensado para scripts o procesos automatizados que necesitan validar el éxito o el motivo del fallo mediante códigos de salida.

## Cómo ejecutar

1. Inicializa el módulo y compila dentro de la carpeta `03-tarea3`:
   ```bash
   go mod init tarea3
   go build -o convertir.exe .

   ## Preguntas

### a) ¿Por qué el valor está en os.Args[1] y no en os.Args[0]?
Porque en Go, `os.Args` es un *slice* donde la posición `0` (`os.Args[0]`) guarda el nombre o ruta del ejecutable que se corrió, y los argumentos que uno escribe en la terminal empiezan desde el índice `1`. En mi `main.go`, `os.Args[1]` recibe el número a convertir, `os.Args[2]` la unidad de origen y `os.Args[3]` la unidad de destino.

### b) ¿Por qué los mensajes de error se escriben en stderr y no con fmt.Println?
Porque `fmt.Println` manda el texto a `stdout` (la salida estándar), donde mi programa imprime el resultado de la conversión. Escribir los fallos en `os.Stderr` (con `fmt.Fprintln` o `fmt.Fprintf`) separa los errores de la salida normal, lo cual ayuda si otro programa o script llama al conversor y solo quiere leer el resultado válido sin que se mezclen mensajes de error.

### c) ¿Qué pasaría en tu programa si escribieras valor, _ := strconv.ParseFloat(...) e ignoraras el error al ejecutar convertir abc KB MB?
Si ignoro el error con `_`, al pasar `"abc"`, `strconv.ParseFloat` devolverá `0` como valor por defecto y `nil` en el error. Como `0` no es menor que cero, la condición `if err != nil || valor < 0` de mi `main.go` no se cumple, así que el programa no detendría la ejecución con `ErrValor` (`2`) y haría la conversión usando `0`.

### d) ¿Qué ventaja tiene declarar los códigos de salida con const en lugar de escribir 1, 2 y 3 dentro de os.Exit?
Usar constantes (`ErrUso`, `ErrValor`, `ErrUnidad`) en un bloque `const` hace que el código sea más legible y fácil de entender al revisarlo, en vez de dejar números sueltos. Además, si necesito cambiar algún código o reutilizarlo en más de un `os.Exit`, evito equivocarme al escribir el número en distintos lugares.