# 03-tarea4 — Conversor modular de unidades de almacenamiento

## Descripción y estructura
Herramienta de línea de comandos en Go (módulo `tarea4`) que convierte valores numéricos positivos entre unidades de almacenamiento en base 1024 (`B`, `KB`, `MB` y `GB`), sin distinción entre mayúsculas y minúsculas en las unidades. Separa la interfaz CLI y validación de argumentos (`main.go`) de la lógica de dominio y conversión (`internal/almacenamiento/almacenamiento.go`).

- `go.mod`: define el módulo `tarea4`.
- `main.go`: lee los argumentos de entrada, valida cantidad y unidades, invoca el paquete `almacenamiento` y devuelve códigos de salida (`0` éxito, `1` error de argumentos/sintaxis, `2` valor inválido o negativo, `3` unidad desconocida).
- `internal/almacenamiento/almacenamiento.go`: exporta las constantes `B`, `KB`, `MB`, `GB` y las funciones `Bytes` y `Convertir`, manteniendo `redondear` como función auxiliar no exportada.

## Cómo ejecutar
Desde la carpeta `03-tarea4`:

```powershell
go build -o convertir.exe .
.\convertir.exe <valor> <unidad_origen> <unidad_destino>

## Preguntas

1. **¿Por qué usar `internal/almacenamiento` en lugar de dejar todo en `main` o en un paquete público?**
   Porque colocar el paquete dentro de `internal/` restringe su importación únicamente a paquetes del mismo módulo (`tarea4`), evitando que código externo dependa de detalles internos y manteniendo la lógica de conversión bien separada de la interfaz CLI.

2. **¿Qué diferencia hay entre `Bytes` y `Convertir`?**
   `Bytes` recibe el nombre de una unidad (`B`, `KB`, `MB` o `GB`) y devuelve su equivalencia en bytes en base 1024 junto con un booleano que indica si es válida; `Convertir` toma el valor numérico y los factores de origen y destino para calcular la conversión final redondeada a dos decimales.

3. **¿Cómo se manejan los errores y códigos de salida en la CLI?**
   Los mensajes de error se envían a `stderr` y el programa finaliza con códigos diferenciados según el tipo de fallo (`1` para argumentos faltantes o sobrantes, `2` para cantidad no numérica o negativa y `3` para unidad desconocida), devolviendo `0` cuando la conversión es exitosa.