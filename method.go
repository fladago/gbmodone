package main

import "fmt"

// Важно, для банковских вычислений использовать другие библиотеки
//
func Float() {
	var a float32 = 359.9
	fmt.Println(a) // 359.9

	var b = float64(a)
	fmt.Println(b) // 359.8999938964844

	var c float64 = 359.9
	fmt.Println(c) // 359.9
}

func Floatd() {
	var a float32 = 358.99999
	var b float32 = 359.00001
	fmt.Println(b - a) //0
}

func Iota() {
	const (
		// можно использовать например текстовые кастомные списки ошибок
		// некий enum
		North int64 = iota
		East
		South
		West
	)
	fmt.Println(North, West)
}
