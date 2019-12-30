package hellointerfaces_test

import (
	"github.com/EmilianoFraga/helloGo/hellointerfaces"
	"testing"
)

func TestRectSetters(t *testing.T) {
	r := hellointerfaces.CreateNewRect(2, 3)

	expectedRectWidth := 2.0
	expectedRectHeight := 3.0

	assertRectDimensions(t, r, expectedRectWidth, expectedRectHeight)

	r.SetWidth(2.5)
	expectedRectWidth = 2.5
	assertRectDimensions(t, r, expectedRectWidth, expectedRectHeight)

	r.SetHeight(3.5)
	expectedRectHeight = 3.5
	assertRectDimensions(t, r, expectedRectWidth, expectedRectHeight)
}

func TestRectArea(t *testing.T) {
	r := hellointerfaces.CreateNewRect(2, 3)

	expectedRectArea := 3.0 * 2.0

	assertGeometryArea(t, r, expectedRectArea)
}

func TestRectPerim(t *testing.T) {
	r := hellointerfaces.CreateNewRect(2, 3)

	expectedRectPerim := (3.0 + 3.0) + (2.0 + 2.0)

	assertGeometryPerim(t, r, expectedRectPerim)
}

func assertRectDimensions(t *testing.T, r hellointerfaces.Rect, expectedWidth float64, expectedHeight float64) {
	if actualWidth := r.GetWidth(); actualWidth != expectedWidth {
		t.Errorf("Wrong rectangle width %f, expected %f", actualWidth, expectedWidth)
	}

	if actualHeight := r.GetHeight(); actualHeight != expectedHeight {
		t.Errorf("Wrong rectangle width %f, expected %f", actualHeight, expectedHeight)
	}
}