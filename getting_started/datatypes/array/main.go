package main

import "fmt"

func main() {
	var surnames [5]string
	var names [5]string
	names[0] = "Hugo"
	surnames[0] = "Leinz"
	names[2] = "Melc"
	surnames[2] = "Backa"

	fmt.Println("My full name is: " + names[0] + " " + surnames[0])
	fmt.Println("i dont know you,  " + names[1])

	// Range
	for i := 0; i < len(names); i++ {
		fmt.Println("Names ", i, ": ", names[i])
	}

	//Sub-parts
	fmt.Println("All names:", names)
	fmt.Println("Some names:", names[1:])
	fmt.Println("Some names:", names[:2])

	var namePointer [3]*string
	namePointer[0] = &names[0]
	namePointer[1] = &names[1]
	namePointer[2] = &names[2]

	for i, aPointer := range namePointer {
		fmt.Println(i, " : ", aPointer)
	}

	for i, aPointer := range namePointer {
		if aPointer != nil {
			fmt.Println(i, " : ", *aPointer)
		}
	}

}
