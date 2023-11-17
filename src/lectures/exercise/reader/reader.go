//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	cmdHello = "Hello"
	cmdBuy   = "buy"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	sumLines := 0
	sumCommands := 0

	for scanner.Scan() {
		//  Create an interactive command line application that supports arbitrary
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())

			switch text {
			case cmdHello:
				sumCommands += 1
				fmt.Println("command responce: hi")
			case cmdBuy:
				sumCommands += 1
				fmt.Println("command responce: buy")
			}
			if text != "" {
				sumLines += 1
			}
		}
	}
	fmt.Printf("You entered %v lines\n", sumLines)
	fmt.Printf("You entered %v commands\n", sumCommands)
}

// 	fmt.Println("Input", input)
// 	fmt.Println("Value", hello)

// 	if input == hello {
// 		fmt.Println("Hello stranger!")
// 		sum += 1
// 	}

// 	if input == buy {
// 		fmt.Println("Buy, thanks for coming!")
// 		sum++
// 	}

// 	if input == "Q" || input == "q" {
// 		fmt.Printf("You send a %v requests.", sum)
// 		fmt.Println("Sorry to see you leaving, buy for now.")
// 		break
// 	}

// 	// fmt.Println(input)

// 	if inputErr != nil {
// 		fmt.Println("Error reading Stdin", inputErr)
// 	}

// 	if inputErr == io.EOF {
// 		break
// 	}
// }

// }
