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
