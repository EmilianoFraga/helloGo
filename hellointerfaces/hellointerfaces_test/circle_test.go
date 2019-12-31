package hellointerfaces_test

import (
	"github.com/EmilianoFraga/helloGo/hellointerfaces"
	"math"
	"testing"
)

func TestCreateNewCircle(t *testing.T) {
	assertCreateNewCircle(t, 2.0)

	assertCreateNewCircle(t, 3.0)
}

func TestCircleSetters(t *testing.T) {
	c, _ := hellointerfaces.CreateNewCircle(2.0)

	assertSetCircleRadius(t, c, 2.5)

	assertSetCircleRadius(t, c, 3.0)

	assertSetCircleRadiusError(t, c, -1.0)
}

func TestCircleArea(t *testing.T) {
	c, _ := hellointerfaces.CreateNewCircle(2)

	expectedCircleArea := math.Pi * (2.0 * 2.0)

	assertGeometryArea(t, c, expectedCircleArea)

}

func TestCirclePerim(t *testing.T) {
	c, _ := hellointerfaces.CreateNewCircle(2)

	expectedCirclePerim := (math.Pi + math.Pi) * 2.0

	assertGeometryPerim(t, c, expectedCirclePerim)
}

func assertCreateNewCircle(t *testing.T, expectedRadius float64) {
	c, ok := hellointerfaces.CreateNewCircle(expectedRadius)

	if ok != nil {
		t.Errorf("Error creating new Circle: %s", ok)
		return
	}

	if actualRadius := c.GetRadius(); actualRadius != expectedRadius {
		t.Errorf("Wrong circle radius %f, expected %f", actualRadius, expectedRadius)
	}
}

func assertCreateNewCircleError(t *testing.T, invalidRadius float64) {
	c, ok := hellointerfaces.CreateNewCircle(invalidRadius)

	if ok == nil {
		t.Errorf("Expected error while creating new circle with radius: %f", invalidRadius)
	}

	if c != nil {
		t.Errorf("After error returned circle should be nil")
	}
}

func assertSetCircleRadius(t *testing.T, c *hellointerfaces.Circle, expectedRadius float64) {
	ok := c.SetRadius(expectedRadius)

	if ok != nil {
		t.Errorf("Error setting radius: %s", ok)
		return
	}

	if actualRadius := c.GetRadius(); actualRadius != expectedRadius {
		t.Errorf("Wrong circle radius %f, expected %f", actualRadius, expectedRadius)
	}
}

func assertSetCircleRadiusError(t *testing.T, c *hellointerfaces.Circle, invalidRadius float64) {
	previousRadius := c.GetRadius()

	ok := c.SetRadius(invalidRadius)

	if ok == nil {
		t.Errorf("Expected error while setting radius: %f", invalidRadius)
	}

	if actualRadius := c.GetRadius(); actualRadius != previousRadius {
		t.Errorf("Circle radius shouldn't change after error. Actual %f, expected %f", actualRadius, previousRadius)
	}
}
