package hellointerfaces

// Rect struct defines a rectangle with width and height
type Rect struct {
	width float64
	height float64
}

// CreateNewRect creates a new Rect with a given width and height
func CreateNewRect(width float64, height float64) Rect {
	return Rect{width: width, height: height}
}

// Area implements Geometry interface method
func (r Rect) Area() float64 {
	return r.width * r.height
}

// Perim implements Geometry interface method
func (r Rect) Perim() float64 {
	return r.width + r.width + r.height + r.height
}

// GetWidth returns the rectangle width
func (r Rect) GetWidth() float64 {
	return r.width
}

// SetWidth changes the Rect width
// must be *Rect, since it changes the subject struct
func (r *Rect) SetWidth(newWidth float64) {
	r.width = newWidth
}

// GetHeight returns the rectangle height
func (r Rect) GetHeight() float64 {
	return r.height
}

// SetHeight changes the Rect height
// must be *Rect, since it changes the subject struct
func (r *Rect) SetHeight(newHeight float64) {
	r.height = newHeight
}
