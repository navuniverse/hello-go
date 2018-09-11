package main

import (
	"fmt"
)

func readUserChoice() int {
	fmt.Println("Please select the operation you wish to perform - ")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Modulation")

	fmt.Print("Enter your choice now - ")

	var choice int

	fmt.Scan(&choice)

	return choice
}

func readParams(a *int, b *int) {
	fmt.Print("Enter First Number - ")
	fmt.Scan(a)

	fmt.Print("Enter Second Number - ")
	fmt.Scan(b)
}

func addition() {
	var a = 0
	var b = 0

	readParams(&a, &b)

	//	fmt.Println("You Entered Numbers - One: ", a, ", Two: ", b)

	fmt.Print("Addition : ")
	fmt.Println(a + b)
}

func subtraction() {
	var a = 0
	var b = 0

	readParams(&a, &b)

	fmt.Print("Subtraction : ")
	fmt.Println(a - b)
}

func multiplication() {
	var a = 0
	var b = 0

	readParams(&a, &b)

	fmt.Print("Multiplication : ")
	fmt.Println(a * b)
}

func division() {
	var a = 0
	var b = 0

	readParams(&a, &b)

	fmt.Print("Division : ")
	fmt.Println(a / b)
}

func modulation() {
	var a = 0
	var b = 0

	readParams(&a, &b)

	fmt.Print("Modulation : ")
	fmt.Println(a % b)
}

func main() {

	fmt.Println("************ Welcome to Calculator Service ***********")

	var keepUsing = "y"

	for keepUsing == "y" || keepUsing == "Y" {

		var choice = readUserChoice()

		switch choice {

		case 1:
			addition()
			break

		case 2:
			subtraction()
			break

		case 3:
			multiplication()
			break

		case 4:
			division()
			break

		case 5:
			modulation()
			break

		default:
			fmt.Println("Invalid Choice Made!!! Please try again!!")
			break
		}

		fmt.Print("Do you wish to try again? (Y | N) - ")
		fmt.Scanln(&keepUsing)
	}

	fmt.Println("Thanks for using calculator.")
}
