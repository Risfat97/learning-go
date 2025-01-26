package main

import "fmt"

func main() {
	var firstname string = "John"
	lastname := "Doe" // Shortcut for assigning a value to a new variable, the type is inferred.
	var age int = 20

	fmt.Println("my name is", firstname, lastname, "and my age is ", age)
	fmt.Println(firstname, "has a length of", len(firstname))

	// You can't do this
	// 1. := is allowed once on a variable
	// lastname := "Carter" 	// error, lastname = "Carter" will work

	// 2. Go is statically typed
	// age = "Hello" 	// error

	// Using other numeric types
	var age1 uint8 = 60
	var monthSalary uint16 = 22700
	var yearSalary uint32 = 12 * 22700

	fmt.Println("Now my age is", age1, "and my monthly salary is", monthSalary, ", my annual salary is", yearSalary)

	// It doesn't work if I try to do:
	// age1 = 257   // uint8 can only store value between 0 and 255
	// monthSalary = 12 * 22700   // uint16 can only store value between 0 and 65535
}
