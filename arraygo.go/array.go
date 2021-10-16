package arraygo

import (
	"fmt"
	"reflect"
)

func ArrayExample() {
	var newArr [2]int
	var arr = [5]int{1, 3, 75, 14}
	fmt.Println("Start arr: ", arr)
	newSlice := arr[:2]
	fmt.Println("Start slice: ", newSlice)
	for i := 0; i < 2; i++ {
		newArr[i] = arr[i]
	}
	arr[0] += 100
	fmt.Println("End arr: ", arr)
	fmt.Println("End newSlice: ", newSlice)
	fmt.Println("Type newSlice: ", reflect.TypeOf(newSlice))
	fmt.Println("newArr: ", newArr)
	fmt.Println("Type newArr: ", reflect.TypeOf(newArr))

	var myArr = [5]int{1, 3, 75, 14}
	var myArr2 = myArr
	fmt.Printf("myArr: %v, myArr2: %v\n", myArr, myArr2)
	myArr[0] += 20
	fmt.Printf("myArr: %v, myArr2: %v\n", myArr, myArr2)
}
