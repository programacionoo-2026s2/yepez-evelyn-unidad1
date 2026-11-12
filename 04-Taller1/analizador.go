package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {

	texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

	// 1. Conteo de bytes y runas
	totalBytes := len(texto)
	totalRunes := utf8.RuneCountInString(texto)

	// 2. Conteo de palabras usando strings.Fields
	palabras := strings.Fields(texto)
	totalPalabras := len(palabras)
	mayusculas := 0
	digitos := 0
	for _, r := range texto {
		if unicode.IsUpper(r) {
			mayusculas++
		}
		if unicode.IsDigit(r) {
			digitos++
		}
	}

	letrasEspeciales := "áéíóúñÁÉÍÓÚÑ"
	conTildeOEnie := 0
	for _, r := range texto {
		if strings.ContainsRune(letrasEspeciales, r) {
			conTildeOEnie++
		}
	}

	// Mostrar resultados formateados tal como se espera en la ejecución
	fmt.Printf("Bytes: %d\n", totalBytes)
	fmt.Printf("Runes: %d\n", totalRunes)
	fmt.Printf("Palabras: %d\n", totalPalabras)
	fmt.Printf("Mayúsculas: %d\n", mayusculas)
	fmt.Printf("Dígitos: %d\n", digitos)
	fmt.Printf("Con tilde/ñ: %d\n", conTildeOEnie)
}

