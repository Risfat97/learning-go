package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings
	quote := "Socrate: All I know is that I don't know anything."
	fmt.Println(strings.Contains(quote, "Socrate")) // true
	fmt.Println(strings.Replace(quote, "Socrate", "Sirtaf", 1))
	fmt.Println(strings.Split(quote, " "))

	// sort
	ratings := []int{12, 17, 9, 15, 11, 8, 7, 19}
	names := []string{"john", "mike", "alice", "bob"}
	sort.Ints(ratings)
	sort.Strings(names)
	index1 := sort.SearchInts(ratings, 11)      // Reasearch the value in the sorted slice
	index2 := sort.SearchStrings(names, "jane") // Returns 2 because it's the position to insert 'jane' to keep the slice sorted
	fmt.Println(index1, index2)
}
