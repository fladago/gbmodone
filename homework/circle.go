// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.

package homework

import (
	"fmt"
	"math"
	"strconv"
)

func CircleArea() (dlina, diametr float64) {
	var input string
	var sq float64
	var err error
	for {
		fmt.Scan(&input)
		// nolint
		sq, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("Enter a number: ")
			continue
		}
		break
	}
	// Каков диаметр окружности (d) если её площадь S?
	// длина l = 2π * √s4π
	dl := 2 * math.Pi * math.Cbrt(sq/math.Pi)
	// диаметр	d = √4S/π
	di := math.Cbrt(sq * 4 * math.Pi)
	return dl, di
}
