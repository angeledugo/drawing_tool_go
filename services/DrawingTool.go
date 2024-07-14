package services

import (
	"fmt"
	"strings"
)

// Slice Bidimencional
type DrawingTool struct {
	canvas [][]rune
}

func (dt *DrawingTool) ExecuteCommands(input string) string {
	commands := strings.Split(input, "\n")
	for _, command := range commands {
		fmt.Println(command)
	}
	return "canvas"

}
