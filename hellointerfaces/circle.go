package hellointerfaces

import (
	// "errors" no need to use errors?
	"fmt"
	"math"
)

// Circle struct defines a circle with radius
type Circle struct {
	radius float64
}

// CreateNewCircle creates a new circle with a given radius
// Returns a pointer to the newly created Circle or nil in case of error
func CreateNewCircle(radius float64) (*Circle, error) {
	if radius < 0.0 {
		// errors.New(fmt.Sprintf("Radius can't be negative %f", radius))
		return nil, fmt.Errorf("Radius can't be negative %f", radius)
	}
	return &Circle{radius: radius}, nil
}

// Area method of Geometry interface
func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

// Perim method of Geometry interface
func (c Circle) Perim() float64 {
	return (math.Pi + math.Pi) * c.radius
}

// GetRadius returns the Circle radius
// can be just Circle, since this is a read only method
func (c Circle) GetRadius() float64 {
	return c.radius
}

// SetRadius changes the Circle radius
// must be *Circle, since it changes the subject struct
// Radius must be >= 0.0
// Returns error in case of invalid radius
func (c *Circle) SetRadius(newRadius float64) error {
	if newRadius < 0.0 {
		return fmt.Errorf("Radius can't be negative %f", newRadius)
	}
	c.radius = newRadius
	return nil
}
