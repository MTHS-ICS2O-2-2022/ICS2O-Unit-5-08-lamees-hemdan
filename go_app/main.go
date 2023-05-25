// Created by: Lamees Hemdan
// Created on: May 2023
//
// This file contains the main function for the go_app application

package main

import (
	"fmt"
)

func main() {
	// this function Divides two numbers and shows the answer

	// input
	var number1 int
	var number2 int

	fmt.Println("Enter two numbers in the spaces below:")
	fmt.Println("Divided")
	fmt.Print("Number1... ")
	fmt.Scanln(&number1)
	fmt.Println("Divided by")
	fmt.Print("Number2... ")
	fmt.Scanln(&number2)

	// process

	var counter int = 0

	for number1 >= number2 {
		number1 = number1 - number2
		counter++
	}

	// output
	if number1 != 0 {
		fmt.Println("The answer is", counter, "with a remainder of", number1)
	} else {
		fmt.Println("The answer is", counter)
	}
	
	fmt.Println("\nDone.")
}
