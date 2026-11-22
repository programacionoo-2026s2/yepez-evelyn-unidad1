// Package almacenamiento provee constantes y funciones para convertir unidades de
// almacenamiento en base 1024 (B, KB, MB y GB).
package almacenamiento

import "math"

// Constantes de unidades de almacenamiento expresadas en bytes (base 1024).
const (
	B  = 1.0
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// Bytes devuelve la cantidad de bytes que representa la unidad indicada
// y un booleano que indica si la unidad es válida.
func Bytes(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return B, true
	case "KB":
		return KB, true
	case "MB":
		return MB, true
	case "GB":
		return GB, true
	default:
		return 0, false
	}
}

// Convertir convierte un valor desde bytesOrigen hacia bytesDestino y lo redondea
// a dos decimales.
func Convertir(valor, bytesOrigen, bytesDestino float64) float64 {
	enBytes := valor * bytesOrigen
	convertido := enBytes / bytesDestino
	return redondear(convertido)
}

// redondear redondea un número decimal a dos cifras decimales.
func redondear(n float64) float64 {
	return math.Round(n*100) / 100
}
