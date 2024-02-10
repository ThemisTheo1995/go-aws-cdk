package main

import "fmt"

// func main() {
// 	manyKnifes := make([]string, 3, 5)

// 	fmt.Println("How many:", len(manyKnifes))

// 	for _, slicer := range manyKnifes {
// 		fmt.Println((slicer))
// 	}

// 	manyKnifes = append(manyKnifes, "katana", "bowie")
// 	for _, slicer := range manyKnifes {
// 		fmt.Println((slicer))
// 	}
// }

func main() {
	beverages := make([]string, 5)

	fmt.Println("How many:", len(beverages))
	beverages[0] = "IRN Brue"
	beverages[1] = "Coca Colla"
	beverages[2] = "Limoncelo"
	beverages[3] = "Red Bull"
	beverages = append(beverages, "Coffee", "Tea", "Mocha", "Chai Latte")

	for i, v := range beverages {
		fmt.Printf("Pointer: %d -- ", i)
		fmt.Println("Value: " + v)
	}
}
