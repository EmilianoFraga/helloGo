package hellointerfaces_test

import (
	"github.com/EmilianoFraga/helloGo/hellointerfaces"
	"math"
	"testing"
)

func TestCircleSetters(t *testing.T) {
	c := hellointerfaces.CreateNewCircle(2)

	expectedCircleRadius := 2.0
	assertCircleRadius(t, c, expectedCircleRadius)

	c.SetRadius(3.0)

	expectedCircleRadius = 3.0
	assertCircleRadius(t, c, expectedCircleRadius)
}

func TestCircleArea(t *testing.T) {
	c := hellointerfaces.CreateNewCircle(2)

	expectedCircleArea := math.Pi * (2.0 * 2.0)

	assertGeometryArea(t, c, expectedCircleArea)

}

func TestCirclePerim(t *testing.T) {
	c := hellointerfaces.CreateNewCircle(2)

	expectedCirclePerim := (math.Pi + math.Pi) * 2.0

	assertGeometryPerim(t, c, expectedCirclePerim)
}

func assertCircleRadius(t *testing.T, c hellointerfaces.Circle, expectedRadius float64) {
	if actualRadius := c.GetRadius(); actualRadius != expectedRadius {
		t.Errorf("Wrong circle radius %f, expected %f", actualRadius, expectedRadius)
	}
}