package services

import (
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	dt := &DrawingTool{}
	dt.createCanvas(20, 4)
	if len(dt.canvas) != 6 || len(dt.canvas[0]) != 22 {
		t.Errorf("Expected canvas size 22x6, but got %dx%d", len(dt.canvas[0]), len(dt.canvas))
	}
}

func TestDrawLine(t *testing.T) {
	dt := &DrawingTool{}
	dt.createCanvas(20, 4)
	dt.drawLine(1, 2, 6, 2)
	if dt.canvas[2][1] != 'x' || dt.canvas[2][6] != 'x' {
		t.Errorf("Expected line from (1,2) to (6,2) with 'x'")
	}
}

func TestDrawRectangle(t *testing.T) {
	dt := &DrawingTool{}
	dt.createCanvas(20, 4)
	dt.drawRectangle(14, 1, 18, 3)
	if dt.canvas[1][14] != 'x' || dt.canvas[3][18] != 'x' {
		t.Errorf("Expected rectangle from (14,1) to (18,3) with 'x'")
	}
}

func TestBucketFill(t *testing.T) {
	dt := &DrawingTool{}
	dt.createCanvas(20, 4)
	dt.bucketFill(10, 3, 'o')
	if dt.canvas[3][10] != 'o' {
		t.Errorf("Expected filled area with 'o' at (10,3)")
	}
}
