package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"tarea4/internal/almacenamiento"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Error: se esperan exactamente 3 argumentos.")
		fmt.Fprintln(os.Stderr, "Uso: convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(1)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número válido.")
		os.Exit(2)
	}

	if valor < 0 {
		fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número mayor o igual a 0.")
		os.Exit(2)
	}

	unidadOrigen := strings.ToUpper(os.Args[2])
	bytesOrigen, okOrigen := almacenamiento.Bytes(unidadOrigen)
	if !okOrigen {
		fmt.Fprintf(os.Stderr, "Error: unidad de origen no reconocida (%s). Use B, KB, MB o GB.\n", os.Args[2])
		os.Exit(3)
	}

	unidadDestino := strings.ToUpper(os.Args[3])
	bytesDestino, okDestino := almacenamiento.Bytes(unidadDestino)
	if !okDestino {
		fmt.Fprintf(os.Stderr, "Error: unidad de destino no reconocida (%s). Use B, KB, MB o GB.\n", os.Args[3])
		os.Exit(3)
	}

	resultado := almacenamiento.Convertir(valor, bytesOrigen, bytesDestino)
	fmt.Printf("%s %s = %.2f %s\n", os.Args[1], unidadOrigen, resultado, unidadDestino)
}
