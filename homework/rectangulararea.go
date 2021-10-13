package homework

import (
	"fmt"
	"strconv"
)

func InputNumber() float64 {
	var input string
	for {
		fmt.Scan(&input)
		// nolint
		length, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("Enter a number: ")
			continue
		}
		return length
	}
}

func RectangleArea(length, width float64) float64 {
	area := length * width
	return area
}
