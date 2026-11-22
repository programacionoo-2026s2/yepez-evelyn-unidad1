## 03-tarea3 - Conversor de unidades de almacenamiento en Go

### Sesión 1
- **Avance:** Creé la carpeta `03-tarea3`, inicié el módulo con `go mod init tarea3` y redacté la base de `main.go`.
- **Problema y solución:** Al correr `go build -o convertir .` al principio salió `go.mod file not found in current directory`, porque aún no había inicializado el módulo; lo resolví ejecutando `go mod init tarea3`.

### Sesión 2
- **Avance:** Completé la validación de argumentos, `strconv.ParseFloat`, normalización con `strings.ToUpper`, el `switch` de unidades y los códigos de salida con `const`.
- **Problema y solución:** En PowerShell, al ejecutar `./convertir <valor> <unidad_origen> <unidad_destino>` con los símbolos `<` y `>`, la consola marcó error porque el operador `<` está reservado; lo resolví pasando valores reales sin esos signos y ejecutando `.\convertir.exe`.

### Sesión 3
- **Avance:** Probé los 8 casos de prueba, revisé los códigos con `$LASTEXITCODE` y ejecuté `gofmt -l .` y `go vet ./...`.
- **Problema y solución:** `gofmt -l .` listó `main.go` por detalles de formato; lo corregí con `gofmt -w main.go` y volví a comprobar que `gofmt -l .` y `go vet ./...` quedaran limpios.

### Sesión 4
- **Avance:** Documenté las cuatro secciones en `03-tarea3/README.md` (incluyendo las 4 preguntas), actualicé el índice en el `README.md` raíz y preparé el cierre en GitHub.