package perimeter

import (
	"math"
)

//Rectangle struct 
type Rectangle struct {
    Width float64
    Height float64
}

// Area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
// Circle struct
type Circle struct {
    Radius float64
}
// Area of a circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
// Triangle struct
type Triangle struct {
    Base   float64
    Height float64
}
// Area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
