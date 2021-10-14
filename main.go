package main

import (
	"fmt"

	"github.com/fladago/gbmodone/homework"
)

func main() {
	// nolint
	// culculator
	// homework.Calculator()
	// BubbleSort
	// arr := []int{256, 1, 2, 5, 35, 17, 84, 1}
	// fmt.Println(arr)
	//  fmt.Println(homework.BubbleSort(arr))

	pn := 200
	primeNumber := homework.PrimeNumber(pn)
	fmt.Println(primeNumber)
}
