package main

// Note :  map and slice using internal pointer

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name, "nc")
}

// pointer receiver\
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name, "nc")
	g.name = "" // cause we use this poibter receiver so ist gonna change the
	// original data
}
