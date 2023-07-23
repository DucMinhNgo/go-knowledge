package main

import "fmt" // fortmat

func main() {
	var a = true;
	fmt.Printf("a=%t\n", a);

	b:="good bye"
	fmt.Printf("b=%s\n", b);

	var c, d int = 1, 2;
	fmt.Printf("c=%d\nd=%d\n", c, d);

	var e = float32(10.5);
	fmt.Printf("e=%f\n", e);
	// fmt.Scanf("Pause");
}