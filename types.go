package main

import "fmt"

func main() {
	var x int = 1 // Integer Data Type
	var y int     //  Integer Data Type
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	var a, b, c = 5.25, 25.25, 14.15 // Multiple float32 variable declaration
	fmt.Println("a : ", a, ", b : ", b, ", c : ", c)

	city := "Delhi"    // String variable declaration
	Country := "India" // Variable names are case sensitive
	fmt.Println("City: ", city)
	fmt.Println("Country: ", Country) // Variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 300 // Multiple type of variable declaration in same line
	fmt.Println("Food: ", food, ", Drink: ", drink, ", Price: ", price)
	m, n := 1, 2
	fmt.Println("Var Initialization - m:  ", m, ", n : ", n)
}
