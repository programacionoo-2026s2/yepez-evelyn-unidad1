package main

import "fmt"

type Permiso uint8

const (
	Leer     Permiso = 1 << iota // 1 (r)
	Escribir                     // 2 (w)
	Ejecutar                     // 4 (x)
)

// Describe los permisos en formato rwx
func describir(p Permiso) string {
	r := '-'
	w := '-'
	x := '-'
	if p&Leer != 0 {
		r = 'r'
	}
	if p&Escribir != 0 {
		w = 'w'
	}
	if p&Ejecutar != 0 {
		x = 'x'
	}
	return string([]rune{r, w, x})
}

func main() {
	// 2. Defina usuario = rwx, grupo = r-x, otros = r--
	usuario := Leer | Escribir | Ejecutar // rwx
	grupo := Leer | Ejecutar              // r-x
	otros := Leer                         // r--

	// Imprima los tres juntos
	fmt.Printf("%s%s%s\n", describir(usuario), describir(grupo), describir(otros))

	// 3. Imprima el valor numérico de cada uno con %d
	fmt.Printf("numérico: %d%d%d\n", usuario, grupo, otros)

	// 4. Conceda escritura a otros con |= y revoque la ejecución del grupo con &^=
	otros |= Escribir  // Añade escritura a otros (pasa de r-- a rw-)
	grupo &^= Ejecutar // Revoca ejecución al grupo (pasa de r-x a r--)

	// Imprimir el resultado final
	fmt.Printf("%s%s%s\n", describir(usuario), describir(grupo), describir(otros))
}
