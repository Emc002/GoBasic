package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	var inc Incremeter = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incremeter interface {
	Increment() int
}

type IntCounter int

// Methods
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
