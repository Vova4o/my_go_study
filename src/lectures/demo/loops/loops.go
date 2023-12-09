package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Counter starts at", sum)

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Adding", i, "to", sum)
	}

	for sum > 15 {
		sum -= 5
		fmt.Println("decrementing at 5", sum)
	}
	fmt.Println(sum)
}
