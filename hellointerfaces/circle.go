package hellointerfaces

import "math"

// Circle struct defines a circle with radius
type Circle struct {
	radius float64
}

// CreateNewCircle creates a new circle with a given radius
func CreateNewCircle(radius float64) Circle {
	return Circle{radius: radius}
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
func (c *Circle) SetRadius(newRadius float64) {
	c.radius = newRadius
}
