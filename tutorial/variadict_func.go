package main

import "fmt"

func min(numbers ...int) int {
	result := numbers[0]

	for _, num := range numbers {
		if result > num {
			result = num
		}
	}

	return result;
}

func main() {
	fmt.Printf("min of 5,3,2 is %d\n", min(5, 3, 2))
	fmt.Printf("min of 5,3,2,1 is %d\n", min(5, 3, 2, 1))
	numbers := []int{5, 3, 2, 1, 0}
	fmt.Printf("min of array [5,3,2,1,0] is %d", min(numbers...))

	fmt.Scanf("Pause")
}