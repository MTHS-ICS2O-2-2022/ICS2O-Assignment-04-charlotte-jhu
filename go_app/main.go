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
	var subtotal float64
	var tax float64
	var total float64

	// input
	fmt.Println("Welcome to the bracelet store!")
	fmt.Println("Please enter the length of your bracelet (5, 6 or 7 inches): ")
	fmt.Scanln(&braceletLength)

	// process
	if braceletLength == 5 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
			subtotal = 4.50
		} else if numberOfColors == 2 {
			subtotal = 5.00
		} else if numberOfColors == 3 {
			subtotal = 6.00
		} else {
			fmt.Println("Invalid input")
		}
	}
	if braceletLength == 6 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
			subtotal = 5.00
		} else if numberOfColors == 2 {
			subtotal = 5.50
		} else if numberOfColors == 3 {
			subtotal = 6.50
		} else {
			fmt.Println("Invalid input")
		}
	}
	if braceletLength == 7 {
		fmt.Println("Please enter the number of colours (1-3): ")
		fmt.Scanln(&numberOfColors)
		if numberOfColors == 1 {
		subtotal = 5.50
		} else if numberOfColors == 2 {
			subtotal = 6.00
		} else if numberOfColors == 3 {
			subtotal = 7.00
		} else {
			fmt.Println("Invalid input")
		}
	}
	if subtotal >= 6.00 {
		fmt.Println("Thank you for your purchase, you get a 10% discount next time!")
	}

	// output
	roundedSubtotal := fmt.Sprintf("%.2f", subtotal)
	fmt.Println("Subtotal: $",roundedSubtotal)
	tax = subtotal * 0.13
	roundedTax := fmt.Sprintf("%.2f", tax)
	fmt.Println("Tax: $",roundedTax)
	total = subtotal + tax
	roundedTotal := fmt.Sprintf("%.2f", total)
	fmt.Println("Total: $",roundedTotal)
	fmt.Println("\nDone.")
}
