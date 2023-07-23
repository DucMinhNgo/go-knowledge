package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var say string;

	scanner := bufio.NewScanner(os.Stdin);

	for scanner.Scan() {
		say = scanner.Text();
		break;
	}

	switch say {
	case "hello", "hi":
		fmt.Printf("Greetings")
	case "good bye":
		fmt.Printf("Bye")
	default:
		fmt.Printf("Unknown")
	}

	whatIsMyType := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("String kkk")
		case int:
			fmt.Println("integer kkk")
		default:
			fmt.Printf("Your are %T", t)
		}
	}

	whatIsMyType("Hello")
	whatIsMyType(1222)
	whatIsMyType(true)

	fmt.Scanf("Pause")
}