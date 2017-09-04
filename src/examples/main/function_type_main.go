package main

import "fmt"

type CustomFunc func(int) int
type A struct {
	X int
	Y int
}

func CallCustomFuct(customFunc CustomFunc) int {
	return customFunc(5)
}

func main() {
	fmt.Println(
		CallCustomFuct(func(x int) int {
			return x + 1
		}))

	fmt.Println(A{X: 2, Y: 3})
}
