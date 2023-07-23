package main

import "fmt"

func main(){
	fmt.Println("Go does not have While loop we should use For instead:")
	i:=5
	for i > 0 {
		fmt.Printf("%d\t", i);
		i--
	}
	fmt.Printf("\nBreak out of the loop:\n")

	for {
		fmt.Printf("loop \t")
		break;
	}

	fmt.Println("\nContinue to the next iteration of the loop:")
	for n:= 0; n <=5; n++ {
		if n%2 == 0 {
			continue;
		}
		fmt.Printf("%d\t", n)
	}

	fmt.Scanf("Pause")
}