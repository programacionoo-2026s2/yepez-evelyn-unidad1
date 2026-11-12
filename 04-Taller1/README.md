# Taller 04 - Analizador de Cadenas en Go

Este proyecto contiene un script en Go (`analizador.go`) desarrollado para cumplir con los requerimientos del taller de análisis y manipulación de cadenas de texto utilizando características nativas del lenguaje.

## Descripción de la Actividad
El programa analiza una cadena de texto predefinida y calcula las siguientes métricas:
1. **Bytes y Runas:** Uso de `len()` y `utf8.RuneCountInString()`.
2. **Palabras:** Extracción y conteo mediante `strings.Fields()`.
3. **Mayúsculas y Dígitos:** Recorrido con bucle `range` evaluando con el paquete `unicode`.
4. **Caracteres especiales:** Conteo de letras con tilde y la letra 'ñ' utilizando `strings.ContainsRune()`.
5. **Diferencia técnica:** Explicación de por qué la cantidad de bytes y runas difiere en una codificación UTF-8.

## Ejecución
Para ejecutar la aplicación, sitúate en la carpeta del taller y utiliza el siguiente comando en tu terminal:

```bash
go run analizador.go