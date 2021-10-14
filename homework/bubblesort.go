package homework

import "fmt"

func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		// 256, 1, 2, 5, 35, 17, 84, 1
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				fmt.Println(arr)
				swapped = true
			}
		}
	}
	return arr
}
