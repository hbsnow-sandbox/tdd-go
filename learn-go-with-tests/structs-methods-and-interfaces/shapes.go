package shapes

import "math"

// Shape 形
type Shape interface {
	Area() float64
}

// Rectangle 長方形
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 面積を求める
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 円
type Circle struct {
	Radius float64
}

// Area 面積を求める
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle 三角形
type Triangle struct {
	Base   float64
	Height float64
}

// Area 面積を求める
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter 周を求める
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
