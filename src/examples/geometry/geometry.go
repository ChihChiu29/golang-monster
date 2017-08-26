package geometry

import "fmt"

type IntPair struct {
	X, Y int
}

func (intPair IntPair) Draw() {
	fmt.Println("Draw at location:", intPair)
}

type Point struct {
	IntPair
	Color int
}

func (p *Point) ChangeColor(color int) {
	p.Color = color
}

func (p Point) Draw() {
	fmt.Println("Draw at point", p.IntPair, "with color", p.Color)
}
