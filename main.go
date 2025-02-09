package main

import "fmt"

func main() {
	// Arrays
	students := [5]string{"John", "Jane", "Eric", "Louise", "Alice"}
	ratings := [5]float32{7.25, 9, 7, 8.5, 9.75}
	students[2] = "Bob"
	ratings[len(ratings)-1] = 8.75

	// You can't do this
	// students[6] = "Milo"   // index 6 out of bounds [0:5]
	// students[0] = 123   // 123 is not a string

	fmt.Println(students, ratings)

	// Slices
	employees := []string{"Bob", "Alice", "John", "Louise"}
	employees = append(employees, "Milo")
	fmt.Println(employees)
	fmt.Println("Employees for the morning are:", employees[:2])
	fmt.Println("Employees for the evening are:", employees[2:])

	// Arrays and slices
	students1 := students[:2]
	fmt.Println(students, students1)
	students1 = append(students1, "Marc", "Toby")
	fmt.Println(students, students1) // Output: [John Jane Marc Toby Alice] [John Jane Marc Toby]
	fmt.Println("cap(students1) =", cap(students1))

	// Creating slice with make
	score := make([]int, 3, 10) // score = [0 0 0]
	score = append(score, 25)   // score [0 0 0 25]
	fmt.Println(score, len(score), cap(score))
}
