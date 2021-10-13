package homework

import (
	"fmt"
	"strconv"
)

func Digits() []int { // (hundred, dozen, unit int)
	var input string
	for {
		fmt.Scan(&input)

		_, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Enter a number: ")
			continue
		}
		break
	}

	mnoj := 1
	for i := 1; i < len(input); i++ {
		mnoj *= 10
	}
	fmt.Println("Множитель", mnoj)
	var digsl []int
	for _, v := range input {
		srtDigit, _ := strconv.Atoi(string(v))
		dig := srtDigit * mnoj
		if dig == 0 {
			mnoj /= 10
			continue
		}
		digsl = append(digsl, dig)
		mnoj /= 10
	}
	return digsl
}
