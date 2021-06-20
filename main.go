package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	low := 1
	high := 100

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Plz think of a number of ", low, " to ", high, "")
	fmt.Println("Press Enter when ready")
	scanner.Scan()
	i := 0
	for {
		// Binary Search Strategy
		guess := (low + high) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that: ")
		fmt.Println("(a) too high? ")
		fmt.Println("(b) too low? ")
		fmt.Println("(c) correct? ")
		scanner.Scan()
		response := scanner.Text()
		// See's if you are lying
		if i < 6 {
			if response == "a" {
				high = guess - 1
				i++
			} else if response == "b" {
				low = guess + 1
				i++
			} else if response == "c" {
				fmt.Println("I won!")
				fmt.Println("the computer took ", i, "amount of trys. ")
				break
			} else {
				fmt.Println("invalid response, try again.")
			}
		} else {
			fmt.Println("You are lying to the program (So Good Bye)")
			break
		}
	}
}
