package main

import (
	"fmt"
	"go-string"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	fmt.Print("adding numbers : ")
	fmt.Print(addNumbers(5, 6))

}

func addNumbers(a int, b int) int {
	return a + b
}
