package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3

	bytesPorB  = 1.0
	bytesPorKB = 1024.0
	bytesPorMB = 1024.0 * 1024.0
	bytesPorGB = 1024.0 * 1024.0 * 1024.0
)

func bytesPorUnidad(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return bytesPorB, true
	case "KB":
		return bytesPorKB, true
	case "MB":
		return bytesPorMB, true
	case "GB":
		return bytesPorGB, true
	default:
		return 0, false
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || valor < 0 {
		fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número mayor o igual a 0.")
		os.Exit(ErrValor)
	}

	origen := strings.ToUpper(os.Args[2])
	destino := strings.ToUpper(os.Args[3])

	bytesOrigen, okOrigen := bytesPorUnidad(origen)
	if !okOrigen {
		fmt.Fprintf(os.Stderr, "Error: unidad de origen no reconocida (%s). Use B, KB, MB o GB.\n", os.Args[2])
		os.Exit(ErrUnidad)
	}

	bytesDestino, okDestino := bytesPorUnidad(destino)
	if !okDestino {
		fmt.Fprintf(os.Stderr, "Error: unidad de destino no reconocida (%s). Use B, KB, MB o GB.\n", os.Args[3])
		os.Exit(ErrUnidad)
	}

	resultado := valor * bytesOrigen / bytesDestino
	fmt.Printf("%g %s = %.2f %s\n", valor, origen, resultado, destino)
}
