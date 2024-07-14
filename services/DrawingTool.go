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
		dt.executeCommand(command)
	}
	return dt.renderCanvas()

}

func (dt *DrawingTool) executeCommand(command string) {
	parts := strings.Fields(command)
	switch parts[0] {
	case "C":
		width := atoi(parts[1])
		height := atoi(parts[2])
		dt.createCanvas(width, height)
	case "L":
		x1, y1, x2, y2 := atoi(parts[1]), atoi(parts[2]), atoi(parts[3]), atoi(parts[4])
		dt.drawLine(x1, y1, x2, y2)
	case "R":
		x1, y1, x2, y2 := atoi(parts[1]), atoi(parts[2]), atoi(parts[3]), atoi(parts[4])
		dt.drawRectangle(x1, y1, x2, y2)
	}
}

func (dt *DrawingTool) createCanvas(width, height int) {
	dt.canvas = make([][]rune, height+2)
	for i := range dt.canvas {
		dt.canvas[i] = make([]rune, width+2)
		for j := range dt.canvas[i] {
			if i == 0 || i == height+1 {
				dt.canvas[i][j] = '-'
			} else if j == 0 || j == width+1 {
				dt.canvas[i][j] = '|'
			} else {
				dt.canvas[i][j] = ' '
			}
		}
	}
}

func (dt *DrawingTool) drawLine(x1, y1, x2, y2 int) {
	if x1 == x2 {
		for y := min(y1, y2); y <= max(y1, y2); y++ {
			dt.canvas[y][x1] = 'x'
		}
	} else if y1 == y2 {
		for x := min(x1, x2); x <= max(x1, x2); x++ {
			dt.canvas[y1][x] = 'x'
		}
	}
}

func (dt *DrawingTool) drawRectangle(x1, y1, x2, y2 int) {
	dt.drawLine(x1, y1, x2, y1)
	dt.drawLine(x1, y2, x2, y2)
	dt.drawLine(x1, y1, x1, y2)
	dt.drawLine(x2, y1, x2, y2)
}

func (dt *DrawingTool) renderCanvas() string {
	var sb strings.Builder
	for _, row := range dt.canvas {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return strings.TrimSpace(sb.String())
}

func atoi(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
