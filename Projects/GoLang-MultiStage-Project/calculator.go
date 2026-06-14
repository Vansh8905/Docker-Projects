package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculator App")
	fmt.Println("Type 'exit' to quit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter calculation (e.g. 1 + 2): ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input closed:", err)
			break
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(text)

		if len(parts) != 3 {
			fmt.Println("Invalid format. Example: 1 + 2")
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid first number")
			continue
		}

		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid second number")
			continue
		}

		var result int

		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Cannot divide by zero")
				continue
			}
			result = left / right
		default:
			fmt.Println("Invalid operator")
			continue
		}

		fmt.Printf("Result: %d\n", result)
	}
}
