package main

import "fmt"

var i = 100

func main() {
	i := 200
	for j := 1; j < 2; j++ {
		i := 43
		fmt.Println("Scope for:", i)
		i = 44
		fmt.Println("Scope for:", i)
	}
	fmt.Println("Scope main:", i)
	outer()
}

func outer() {
	fmt.Println("Function: ", i)
}
