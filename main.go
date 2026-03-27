package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	guide := ""
	isOver := false
	yORn := ""

	fmt.Println("Welcome to the CLI calculator. Type 'help' for guide or 'go' to continue")
	fmt.Scan(&guide)

	if guide == "help" {
		fmt.Println(".......................")
		fmt.Println(" ")
		fmt.Println("Available operations are: 'add', 'mul', 'sub', 'div'")
		fmt.Println("Example Usage: add <num> <num> then click enter")
		fmt.Println(".......................")
		fmt.Println(" ")
		main()
	} else {
		fmt.Println("Type your operation")

		for isOver != true {
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			parts := strings.Fields(line)
			if len(parts) != 3 {
				fmt.Println("Wrong number of arguments. Usage example: add <num> <num> (eg, add 1 2)")
				continue
			}
			op := parts[0]
			left, err1 := strconv.Atoi(parts[1])
			right, err2 := strconv.Atoi(parts[2])

			if err1 != nil || err2 != nil {
				fmt.Println("Arguments must be numbers. Usage example: add 1 5")
				continue
			}

			switch op {
			case "add" :
				fmt.Println(add(left, right));
			case "sub" :
				fmt.Println(sub(left, right));
			case "mul" :
				fmt.Println(mul(left, right));
			case "div" :
				if right == 0 {
		        fmt.Println("Cannot divide by zero")
				continue
	        }
				fmt.Println(div(left, right));
			default :
			    fmt.Println("Unknown command. Type 'help' for available operations")
			}

			fmt.Println("Do you wish to continue [yes/quit]")
			fmt.Println(".......................")
		    fmt.Println(" ")
			fmt.Scan(&yORn)

			if yORn == "quit" {
				isOver = true
				continue
			} else {
				fmt.Println("Input your operation")
			}
		}
	}
}

func add(left, right int) int {
	return left + right
}

func sub(left, right int) int {
	return left - right
}

func mul(left, right int) int {
	return left * right
}

func div(left, right int) int {
	return left / right
}