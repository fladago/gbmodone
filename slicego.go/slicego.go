package slicego

import "fmt"

// nolint
func sliceExample() {
	slice := []int64{}
	fmt.Println(slice, len(slice), cap(slice))
	slice2 := []int64{1, 2, 3}
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 11, 12, 13)
	fmt.Println(slice2, len(slice2), cap(slice2), "\n ---------------")

	arr := [5]int64{1, 2, 3, 4, 5}
	slice = arr[3:]
	fmt.Println(arr, slice, cap(arr))
	arr[3] += 100
	fmt.Println(arr, slice)
	slice[1] += 100
	fmt.Println(arr, slice)

	newSlise := make([]int64, 10)
	fmt.Println(newSlise, cap(newSlise))
}
