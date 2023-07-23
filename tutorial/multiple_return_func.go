package main

import "fmt"

func rectPerimeter1(a int, b int) int {
	return (a + b) * 2
}

func rectArea1(a, b int) int {
	return a * b
}

func rectPeriAndArea1(a, b int) (int, int) {
	rectPeri := rectPerimeter1(a, b);
	rectArea := rectArea1(a, b);

	return rectPeri, rectArea;
}

func rectPeriAndArea2(a, b int) (rectPeri int, rectAr int) {
	rectPeri = rectPerimeter1(a, b);
	rectAr = rectArea1(a, b);
	return;
}

func rectPeriAndArea3(a, b int) (rectPeri, rectAr int) {
	rectPeri = rectPerimeter1(a, b);
	rectAr = rectArea1(a, b);
	return;	
}

func main() {
	a := 2
	b := 3
	rectPeri1, rectAr1 := rectPeriAndArea1(a, b)
	fmt.Printf("rectPeri1=%d, rectAr1=%d\n", rectPeri1, rectAr1)
	rectPeri2, rectAr2 := rectPeriAndArea2(a, b)
	fmt.Printf("rectPeri2=%d, rectAr2=%d\n", rectPeri2, rectAr2)
	rectPeri3, rectAr3 := rectPeriAndArea3(a, b)
	fmt.Printf("rectPeri3=%d, rectAr3=%d\n", rectPeri3, rectAr3)

	fmt.Scanf("Pause")
}