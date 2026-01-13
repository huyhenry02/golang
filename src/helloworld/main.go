package main

import "fmt"

func rectangleArea(length, width int) (area int) {
	area = length * width
	return
}

func main() {
	fmt.Printf("Area: %d\n", rectangleArea(5, 3))
}
