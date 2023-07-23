package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var (
	PASSWORD = "123456"
)

func getRandNumber(n int) int {
	return rand.Intn(n)
}

func main() {
	var password string;

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter password \n");

	for scanner.Scan() {
		password = scanner.Text()
		break;
	}

	if password == PASSWORD {
		fmt.Printf("Your pasword is correct\n");
	} else {
		fmt.Printf("Your pasword is incorrect\n");
	}

	if num:= getRandNumber(20); num < 19 {
		fmt.Printf("number %d is greater than 10", num)
	} else if num < 10 {
		fmt.Printf("number %d is smaller than 10", num)
	} else {
		fmt.Printf("number %d is equal than 10", num)
	}

    fmt.Scanf("Pause")
}