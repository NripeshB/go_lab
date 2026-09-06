package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	height float64
	width  float64
}

func (t Triangle) Area() float64 {
	return (t.height * t.width) / 2

}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width

}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi

}

func Perimeter(Rec Rectangle) float64 {
	return (Rec.height + Rec.width) * 2
}

// func Area(Rec Rectangle) float64 {
// 	return Rec.width * Rec.height
// }
