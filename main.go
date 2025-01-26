package main

import "fmt"

func main() {
	var fullname string = "John Doe"
	var age int = 27
	jobTitle := "Go Developer"
	annualSalary := 70000.00

	fmt.Printf("%s %d yo - %s\n", fullname, age, jobTitle)
	fmt.Printf("Annual Salary $%.0f\n", annualSalary)
	fmt.Printf("Monthly Salary $%.2f\n\n", annualSalary/12)

	// Using some printing formats
	fmt.Printf("%v %v yo - %v\n", fullname, age, jobTitle)
	fmt.Printf("%q %T\n", fullname, fullname)
	fmt.Printf("You have an annual salary greater than $100000. %t\n", (annualSalary > 100000))
	fmt.Printf("The value %q is stored in %p.\n", fullname, &fullname)
}
