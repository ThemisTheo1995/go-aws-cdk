package main

import "fmt"

func main() {
	// Explicit variables declaration
	var varInt = 42
	var varString = "mega"
	var varFloat = 3.14
	fmt.Printf("Value Int : %v\n", varInt)
	fmt.Printf("Type Int  : %T\n", varInt)
	fmt.Printf("Value String : %v\n", varString)
	fmt.Printf("Type String  : %T\n", varString)
	fmt.Printf("Value Float : %v\n", varFloat)
	fmt.Printf("Type Float  : %T\n", varFloat)

	fmt.Println("")

	// Implicit variables declaration
	implVarInt := 42
	implVarString := "mega"
	implVarFloat := 3.14
	fmt.Printf("Value Int implicit var : %v \n", implVarInt)
	fmt.Printf("Type Int implicit var  : %T \n", implVarInt)
	fmt.Printf("Value String implicit var : %v \n", implVarString)
	fmt.Printf("Type String implicit var  : %T \n", implVarString)
	fmt.Printf("Value Float implicit var : %v \n", implVarFloat)
	fmt.Printf("Type Float implicit var  : %T \n", implVarFloat)
}
