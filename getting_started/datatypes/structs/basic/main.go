package main

import (
	"fmt"
)

type HeroChildren struct {
	name string
	str  int
	hp   int
}
type Hero struct {
	name     string
	str      int
	hp       int
	agility  int
	children HeroChildren
}

func main() {
	spiderman := new(Hero)
	fmt.Println(*spiderman)

	spiderman.name = "Peter Parker"
	fmt.Println(*spiderman)
	spiderman.str = 6
	fmt.Println(*spiderman)
	spiderman.str = 120
	fmt.Println(*spiderman)
	spiderman.agility = 9
	fmt.Println(*spiderman)

	spidermanSon := HeroChildren{
		hp:   25,
		str:  2,
		name: "Peter Parker's son",
	}
	spidermanSon.hp = 20
	spiderman.children = spidermanSon

	fmt.Println(*spiderman)
	fmt.Println(spidermanSon)
}
