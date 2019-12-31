package hellointerfaces

// Rect struct defines a rectangle with width and height
type Rect struct {
	width  float64
	height float64
}

// CreateNewRect creates a new Rect with a given width and height
// width and height must be >= 0.0
// Returns a pointer to newly created Rect or RectError in case of error
func CreateNewRect(width float64, height float64) (*Rect, error) {
	if width < 0.0 {
		return nil, &RectError{"Invalid width", width}
	}

	if height < 0.0 {
		return nil, &RectError{"Invalid height", height}
	}

	return &Rect{width: width, height: height}, nil
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
// width must be >= 0
// Returns RectError in case of error
func (r *Rect) SetWidth(newWidth float64) error {
	if newWidth < 0.0 {
		return &RectError{"Invalid width", newWidth}
	}

	r.width = newWidth

	return nil
}

// GetHeight returns the rectangle height
func (r Rect) GetHeight() float64 {
	return r.height
}

// SetHeight changes the Rect height
// must be *Rect, since it changes the subject struct
func (r *Rect) SetHeight(newHeight float64) error {
	if newHeight < 0.0 {
		return &RectError{"Invalid height", newHeight}
	}

	r.height = newHeight

	return nil
}
