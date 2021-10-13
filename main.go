package main

import (
	"fmt"

	"github.com/fladago/gbmodone/homework"
)

func main() {
	// nolint
	// Task1
	// var length, width float64
	// fmt.Print("Add length of rectangle area: ")
	// length = homework.InputNumber()
	// fmt.Print("Add width of rectangle area: ")
	// width = homework.InputNumber()
	// area := homework.RectangleArea(length, width)
	// fmt.Println(reflect.TypeOf(area))
	// fmt.Printf("Rectangle area is: %0.2f", area)

	// Task 2
	// fmt.Print("Add area: ")
	// height, diametr := homework.CircleArea()
	// fmt.Println(height, diametr)

	// Task 3
	fmt.Print("Add digit: ")
	digsl := homework.Digits()
	for _, v := range digsl {
		fmt.Println(v)
	}
}
