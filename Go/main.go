package main

import "fmt"

func main() {
	fmt.Println("Hello World!") // Print "Hello World!" to the console
	result := sumNumbers(5, 5)  // Call the sumNumbers function with arguments 5 and 5, and assign the result to the variable "result"
	fmt.Println("Sum:", result) // Print the sum of the numbers to the console

	power := 1000
	fmt.Printf("default power is %d\n", power)

	name, power := "Goku", 9000
	fmt.Printf("%s's power is %d\n", name, power)
}

func sumNumbers(x, y int) int {
	return x + y // Return the sum of the two numbers
}

// Output:
// Hello World!
// Sum: 10
