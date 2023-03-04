package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1; i++ {
		hello("Rank : ", i)
		hello2("Paris", "France")
	}
	sum("Silva", 1, 2, 3, 4, 5)
	s := sum2(1, 2, 3, 4, 5)
	t := sum3(1, 2, 3, 4, 5)
	u := sum4(1, 2, 3, 4, 5)
	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("the sum is here with return", s)
	fmt.Println("the sum is here with return using pointer", *t)
	fmt.Println("the sum is here with return using pointer", u)
	fmt.Println("Retunr divide", d)

	// anonymus function
	func() {
		fmt.Println("im anonymus")
	}() // <== to invoke this anonymus function

	f := func() {
		fmt.Println("in the inside of variable")
	}
	f()

	var g func() = func() {
		fmt.Println("in the inside of variable too")
	}
	g()
}

func hello(ranking string, idx int) {
	fmt.Println(ranking, idx)
}

func hello2(place, perfect string) {
	fmt.Println(place, perfect)
}

func sum(masg string, values ...int) {
	fmt.Println(values)
	dataResult := 0
	for _, v := range values {
		dataResult += v
	}
	fmt.Println(masg)
	fmt.Println("the sum is ", dataResult)
}

func sum2(values ...int) int {
	fmt.Println(values)
	dataResult := 0
	for _, v := range values {
		dataResult += v
	}
	return dataResult
}

func sum3(values ...int) *int {
	fmt.Println(values)
	dataResult := 0
	for _, v := range values {
		dataResult += v
	}
	return &dataResult
}

func sum4(values ...int) (dataResult int) {
	fmt.Println(values)
	for _, v := range values {
		dataResult += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
