// Created by: Charlotte Jhu
// Created on: April 2023
//
// This program calcualtes the price of your order from a bracelet store

package main

import (
	"fmt"
)

func main() {
	// variables
	var braceletLength float64
	var numberOfColors float64

	// input
	fmt.Println("Welcome to the bracelet store!")
	fmt.Println("Please enter the length of your bracelet (5, 6 or 7 inches): ")
	fmt.Scanln(&braceletLength)

	// process and output
	if braceletLength == 5 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
			fmt.Println("Your bracelet costs $4.50")
		} else if numberOfColors == 2 {
			fmt.Println("Your bracelet costs $5.00")
		} else if numberOfColors == 3 {
			fmt.Println("Your bracelet costs $6.00")
		} else {
			fmt.Println("Invalid input")
		}
	}
	if braceletLength == 6 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
			fmt.Println("Your bracelet costs $5.00")
		} else if numberOfColors == 2 {
			fmt.Println("Your bracelet costs $5.50")
		} else if numberOfColors == 3 {
			fmt.Println("Your bracelet costs $6.50")
		} else {
			fmt.Println("Invalid input")
		}
	}
	if braceletLength == 7 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
			fmt.Println("Your bracelet costs $5.50")
		} else if numberOfColors == 2 {
			fmt.Println("Your bracelet costs $6.00")
		} else if numberOfColors == 3 {
			fmt.Println("Your bracelet costs $7.00")
		} else {
			fmt.Println("Invalid input")
		}
	}

	fmt.Println("\nDone.")
}
