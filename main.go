package main

import "fmt"

func main() {
	const N int = 5
	students := [N]string{"John", "Jane", "Eric", "Louise", "Alice"}
	ratings := [N]float32{7.25, 9, 7, 8.5, 9.75}

	for i := 0; i < N; i++ {
		fmt.Printf("%s : %.2f\n", students[i], ratings[i])
	}

	for i, student := range students {
		fmt.Printf("Student %d : %s\n", i, student)
	}

	for _, rate := range ratings {
		fmt.Println(rate)
	}
}
