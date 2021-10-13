package main

import "fmt"

// nolint
func Continue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continuing loop...")
			continue
		}
		fmt.Println("The value of i is ", i)
	}
	fmt.Println("Exiting program")
}

// nolint
func ContinueToLabel() {
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			break OuterLoop
		}
	}
}
