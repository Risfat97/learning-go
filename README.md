# Learning The Go language

Structure of a Go program: 

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

>**Note**: Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.


## Variables

There are two ways to declare variables:  
1. With the `var` keyword  
`var variableName type = value` or `var variableName = value`  

2. With the `:=` sign  
`variableName := value`  

>**Note**: `var` can be used inside and outside functions while `:=` coan only be used inside functions.  

Look at the `main.go` file to see an example.


## String and numbers  

Strings are immutable. The predeclared string type is `string`. The length of a string `s` can be discovered using the built-in function `len`.  

Numbers have different types:  
- `uint8`, `uint16`, `uint32`, `uint64`
- `int8`, `int16`, `int32`, `int64`
- `float32`, `float64`
- `complex64`, `complex128`
- `byte` alias for `uint8`
- `rune` alias for `int32`


## Array and Slices  

In Go, an aray is a numbered sequence of elements of a single type, called the element type. The length of array `a` can be discovered using the built-in function `len`. The elements can be addressed by integer indices 0 through `len(a)-1`.  
Once you create an array you cannot change its length.  
Examples:  

```go
[10]int
[2*N]float64
[3][3]int
```

>**Note**: An array type T may not have an element of type T, or of a type containing T as a component, directly or indirectly, if those containing types are only array or struct types.  


A slice is a reference to a contiguous segment of an underlying array. It consists of three components:
- A pointer to the first element of the segment in the array.
- A length (the number of elements in the slice). You can access it using len(slice)
- A capacity (the maximum number of elements the slice can hold without reallocating memory). You can access it using cap(slice)

The main difference of a slice to an array is that you can change the length of a slice by adding items to it or taking away items from it.  
The syntax for creating a slice is the same as for an array, except that you don't specify the length. 
Examples:  

```go
[]int
[]T
[][3]int
```

>**Note**: Capacity determines how much a slice can grow without needing to allocate a new array. If you append elements to a slice and it exceeds its capacity, Go will automatically create a new array with a larger capacity and copy the elements over.  

A new, initialized slice value for a given element type T may be made using the built-in function `make`, which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing  

>`make([]T, length, capacity)`

produces the same slice as allocating an array and slicing it, so these two expressions are equivalent:  

```go
make([]int, 50, 100)
new([100]int)[0:50]
```


## Loops  

`for` is Go’s only looping construct. Here is a classic initial/condition/after for loop:  

```go
for i := start; i < end; i++ {
    // ...
}
```

If you skip the init and post statements, you get a `while` loop.  

```go
for n < end {
    // ...
    n++
}
```

If you skip the condition as well, you get an infinite loop.  

```go
for {
    // ...
}
```

#### For-each range loop  
Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.  
The `range` form of the for loop iterates over a slice or map.  
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```go
names := []string{"john", "mike", "alice", "bob"}

for i, s := range names {
    fmt.Println(i, s)
}
```
