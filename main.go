package main

import (
	"errors"
	"fmt"
)

// Возвращаемая ошибка
// nolint
func Foo(s string) (string, error) {
	if s == "" {
		return s, errors.New("empty string s")
	}
	return "ok", nil
}

func main() {
	// nolint
	// Break()
	// NestedBreak()
	// Continue()
	// ContinueToLabel()
	arr := []int{3, 520, 20, 3, 2, 13, 7}
	fmt.Println(BubbleSort(arr))
	arr2 := []int{3, 520, 20, 3, 2, 13, 7}
	fmt.Println(BubbleSort2(arr2))
}
