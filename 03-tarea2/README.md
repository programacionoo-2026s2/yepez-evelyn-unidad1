# Tarea 2: Gestión de Permisos con Operadores Bit a Bit en Go

## Estrategia de Desarrollo

Para resolver el problema de asignación y consulta de permisos de archivos (estilo `chmod` de Unix/Linux), se implementó una solución basada en **máscaras de bits (bitmasks)** sobre un tipo de dato entero sin signo de 8 bits (`uint8`).

### 1. Definición de Permisos y Constantes
Se utilizó el identificador `iota` junto con el operador de desplazamiento de bits a la izquierda (`<<`) para asignar a cada permiso una potencia de 2 única, evitando colisiones en binario:
- **Leer (`r`)**: `1 << 0` (Valor binario `001`, decimal `1`)
- **Escribir (`w`)**: `1 << 1` (Valor binario `010`, decimal `2`)
- **Ejecutar (`x`)**: `1 << 2` (Valor binario `100`, decimal `4`)

*(Nota: Para el cálculo numérico tipo `chmod`, se mapean los pesos estándar donde Lectura = 4, Escritura = 2 y Ejecución = 1).*

### 2. Operaciones Bit a Bit Implementadas
* **Asignación combinada (`|` - OR):** Se empleó para agrupar múltiples permisos en una sola variable (por ejemplo, `Leer | Escribir | Ejecutar` para formar `rwx`).
* **Verificación de permisos (`&` - AND):** En la función `describir()`, se evalúa si un bit específico está activo comparando el permiso recibido contra cada constante (`p & Leer != 0`). Si el resultado es distinto de cero, se asigna el carácter correspondiente (`r`, `w`, `x`); de lo contrario, se mantiene `-`.
* **Concesión de permisos (`|=` - OR asignación):** Se utilizó sobre la variable `otros` para añadir el permiso de escritura sin alterar los permisos previos.
* **Revocación de permisos (`&^=` - AND NOT asignación):** Se aplicó sobre la variable `grupo` para limpiar (apagar) el bit correspondiente a ejecución sin afectar los demás permisos.

## Ejecución del Programa

Desde la terminal, ubicado dentro del directorio `03-tarea2`, ejecutar:

```bash
go run permisos.go