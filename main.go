package main

import (
	"fmt"
	"io/ioutil"

	"github.com/angeledugo/drawing_tool_go/services"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo de entrada:", err)
		return
	}

	dt := &services.DrawingTool{}
	output := dt.ExecuteCommands(string(input))

	err = ioutil.WriteFile("output.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo de salida:", err)
		return
	}

	fmt.Println("Dibujo completado. Revisa output.txt para ver el resultado.")
}
