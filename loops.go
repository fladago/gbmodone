package main

import "fmt"

// nolint
func until() {
	// in golang only for loop
	confition := false
	for ok := true; ok; ok = !confition {
		fmt.Println("until")
	}
}

// nolint
func until2() {
	condition := false
	for {
		if condition {
			break
		}
	}
}
