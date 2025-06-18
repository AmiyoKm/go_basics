package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
type Abs interface {
	Abs() float64
}

func main() {
	var abs Abs
	fmt.Println(abs)
	f := MyFloat(-math.Sqrt2)
	v := Vertex{4, 5}

	abs = f
	abs = &v

	//abs = v
}