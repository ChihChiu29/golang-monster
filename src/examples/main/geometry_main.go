package main

import (
	"examples/geometry"
	"fmt"
	"utils"
)

func main() {
	fmt.Println("Hello World!")
	utils.PrintHelloWorld()
	fmt.Println(utils.Sin())

	intPair := geometry.IntPair{X: 20, Y: 45}
	intPair.Draw()

	point := geometry.Point{IntPair: geometry.IntPair{X: 20, Y: 45}}
	point.Draw()
	point.ChangeColor(1)
	point.Draw()
}
