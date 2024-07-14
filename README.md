# Drawing Tool - Go Implementation

Este proyecto implementa una herramienta de dibujo simple utilizando Go.

## Requisitos previos

- Go 1.16 o superior

## Instalación

1. Clona este repositorio
2. Navega al directorio del proyecto

## Uso

1. Crea un archivo `input.txt` con los comandos de dibujo dentro del directorio raiz
2. Ejecuta el programa: 

```sh
git https://github.com/angeledugo/drawing_tool_go
cd drawing_tool_go
go run main.go
go test ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html