package main

import "fmt"

// nolint
func Iota() {
	const (
		North = iota
		East
		South
		West
	)
	fmt.Println(North, West)
}
