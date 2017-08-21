package main

import (
	"examples"
	"fmt"
	"utils"
)

func main() {
	fmt.Println("Hello World!")
	utils.PrintHelloWorld()
	fmt.Println(utils.Sin())

	intPair := examples.IntPair{X: 20, Y: 45}
	intPair.Draw()

	point := examples.Point{IntPair: examples.IntPair{X: 20, Y: 45}}
	point.Draw()
	point.ChangeColor(1)
	point.Draw()
}
