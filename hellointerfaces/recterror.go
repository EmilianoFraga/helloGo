package hellointerfaces

import "fmt"

// RectError custom Rect Error struct
type RectError struct {
	msg          string
	invalidValue float64
}

func (e *RectError) Error() string {
	return fmt.Sprintf("Error creating / using rectangle: %s %f", e.msg, e.invalidValue)
}
