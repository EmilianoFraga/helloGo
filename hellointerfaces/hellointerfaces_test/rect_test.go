package hellointerfaces_test

import (
	"github.com/EmilianoFraga/helloGo/hellointerfaces"
	"testing"
)

func TestCreateNewRect(t *testing.T) {
	assertCreateNewRect(t, 2.0, 3.0)

	assertCreateNewRect(t, 0.0, 3.0)

	assertCreateNewRect(t, 2.0, 0.0)

	assertCreateNewRect(t, 0.0, 0.0)

	assertCreateNewRectError(t, 2.0, -0.0000001)

	assertCreateNewRectError(t, -0.0000001, 3.0)
}

func TestRectSetters(t *testing.T) {
	r, _ := hellointerfaces.CreateNewRect(2, 3)

	// this also works... !!!
	// assertRectSetDimension(t, r.SetWidth, func() func() float64 { return r.GetWidth }, 3.5)

	// getterGetterFunc is a function that returns a function that returns a float64
	getterGetterFunc := func() func() float64 { return r.GetWidth }
	assertRectSetDimension(t, r.SetWidth, getterGetterFunc, 3.5)

	assertRectSetDimension(t, r.SetWidth, getterGetterFunc, 0.0)

	assertRectSetDimensionError(t, r.SetWidth, getterGetterFunc, -0.000001)

	getterGetterFunc = func() func() float64 { return r.GetHeight }
	assertRectSetDimension(t, r.SetHeight, getterGetterFunc, 3.5)

	assertRectSetDimension(t, r.SetHeight, getterGetterFunc, 0.0)

	assertRectSetDimensionError(t, r.SetHeight, getterGetterFunc, -0.000001)
}

func TestRectArea(t *testing.T) {
	r, _ := hellointerfaces.CreateNewRect(2, 3)

	expectedRectArea := 3.0 * 2.0

	assertGeometryArea(t, r, expectedRectArea)
}

func TestRectPerim(t *testing.T) {
	r, _ := hellointerfaces.CreateNewRect(2, 3)

	expectedRectPerim := (3.0 + 3.0) + (2.0 + 2.0)

	assertGeometryPerim(t, r, expectedRectPerim)
}

func assertCreateNewRect(t *testing.T, expectedWidth float64, expectedHeight float64) {
	r, ok := hellointerfaces.CreateNewRect(expectedWidth, expectedHeight)

	if ok != nil {
		t.Errorf("Error creating new Rect: %s", ok)
		return
	}

	if actualWidth := r.GetWidth(); actualWidth != expectedWidth {
		t.Errorf("Wrong rectangle width %f, expected %f", actualWidth, expectedWidth)
	}

	if actualHeight := r.GetHeight(); actualHeight != expectedHeight {
		t.Errorf("Wrong rectangle width %f, expected %f", actualHeight, expectedHeight)
	}
}

func assertCreateNewRectError(t *testing.T, expectedWidth float64, expectedHeight float64) {
	r, ok := hellointerfaces.CreateNewRect(expectedWidth, expectedHeight)

	if ok == nil {
		t.Errorf("Expected error while creating new Rect with dimensions: w:%f  h:%f", expectedWidth, expectedHeight)
	}

	if r != nil {
		t.Errorf("After error returned Rect should be nil")
	}
}

func assertRectSetDimension(t *testing.T, setterFunc func(value float64) error, getterGetterFunc func() func() float64, expectedValue float64) {
	ok := setterFunc(expectedValue)

	if ok != nil {
		t.Errorf("Error setting dimension: %s", ok)
		return
	}

	getterFunc := getterGetterFunc()

	// this also works.... ()() !!!!
	// if actualValue := getterGetterFunc()(); actualValue != expectedValue {

	if actualValue := getterFunc(); actualValue != expectedValue {
		t.Errorf("Wrong rectangle dimention %f, expected %f", actualValue, expectedValue)
	}
}

func assertRectSetDimensionError(t *testing.T, setterFunc func(value float64) error, getterGetterFunc func() func() float64, invalidValue float64) {
	getterFunc := getterGetterFunc()

	previousValue := getterFunc()

	ok := setterFunc(invalidValue)

	if ok == nil {
		t.Errorf("Expected error while setting one Rect dimension with value: %f", invalidValue)
	}

	// this also works.... ()() !!!!
	// if actualValue := getterGetterFunc()(); actualValue != expectedValue {

	if actualValue := getterFunc(); actualValue != previousValue {
		t.Errorf("Rect dimension shouldn't change after error. Actual %f, expected %f", actualValue, previousValue)
	}
}
