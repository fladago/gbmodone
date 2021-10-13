package main

import "fmt"

// nolint
func Break() {
	// nolint
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking out of loop")
			break
		}
		fmt.Println("The value of i is", i)
	}
}

// nolint
func NestedBreak() {
	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("1Breaking out of outer loop")
			break
		}
		fmt.Println("2The value of outer is", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("3Breaking out of inner loop")
				break
			}
			fmt.Println("4The value of inner is", inner)
		}
	}
}
