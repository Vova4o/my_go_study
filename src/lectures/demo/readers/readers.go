package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader((os.Stdin))

	sum := 0

	for {
		// asking for input the numbers
		input, inputErr := r.ReadString(' ')
		// Triming the string all white spaces
		n := strings.TrimSpace(input)
		// check if n = nothing then run again
		if n == "" {
			continue
		}
		// when we done, convert the numbers
		// and add them together
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		// checking for input error
		// if we reached the end of file
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin", inputErr)
		}

	}
	fmt.Printf("sum: %v\n", sum)
}
