package main

import "fmt"

func rectPerimeter(a int, b int) int {
	return (a + b)*2
}

func rectArea(a, b int) int {
	return a * b;
}

func main() {
	a := 2
	b := 3
	rectPeri := rectPerimeter(a, b);
	fmt.Printf("Perimeter of rectangle is %d\n", rectPeri)
	rectAr := rectArea(a, b)
	fmt.Printf("Area of rectangle is %d\n", rectAr)
	fmt.Scanf("Pause")
}