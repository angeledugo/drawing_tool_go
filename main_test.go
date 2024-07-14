package main

import (
	"testing"

	"github.com/angeledugo/drawing_tool_go/services"
)

func TestDrawingTool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Create Canvas",
			input:    "C 20 4",
			expected: "----------------------\n|                    |\n|                    |\n|                    |\n|                    |\n----------------------",
		},
		{
			name:     "Draw Line",
			input:    "C 20 4\nL 1 2 6 2",
			expected: "----------------------\n|                    |\n|xxxxxx              |\n|                    |\n|                    |\n----------------------",
		},
		{
			name:     "Draw Rectangle",
			input:    "C 20 4\nR 16 1 20 3",
			expected: "----------------------\n|               xxxxx|\n|               x   x|\n|               xxxxx|\n|                    |\n----------------------",
		},
		{
			name:     "Bucket Fill",
			input:    "C 20 4\nL 1 2 6 2\nL 6 3 6 4\nR 16 1 20 3\nB 10 3 o",
			expected: "----------------------\n|oooooooooooooooxxxxx|\n|xxxxxxooooooooox   x|\n|     xoooooooooxxxxx|\n|     xoooooooooooooo|\n----------------------",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &services.DrawingTool{}
			result := dt.ExecuteCommands(tt.input)
			if result != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, result)
			}
		})
	}
}
