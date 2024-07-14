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

	fmt.Printf(output)
}
