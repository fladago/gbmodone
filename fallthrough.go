package main

import "fmt"

// nolint
func Fallthrough(value int) {
	switch value {
	case 1:
		// Для 1, обрабатываем как 1
		// и проваливаемся (fallthrough) к 0
		fmt.Println("one", value)
		fallthrough
	case 0:
		// Для 0б просто печатаем ноль.
		fmt.Println("zero", value)
		// break

	// nolint
	case 3:
		fmt.Println("33", value)
	default:
	}
}
