package homework

import (
	"fmt"
	"math"
	"os"
)

func NumValidation() float64 {
	var a float64
	for {
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Print("Вы должны ввести число: ")
			continue
		}
		return a
	}
}

func Calculator() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	a = NumValidation()
	fmt.Println(a)

	fmt.Print("Введите второе число: ")
	b = NumValidation()
	fmt.Println(b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)

	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
