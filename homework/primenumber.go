package homework

func PrimeNumber(num int) []int {
	var arrPrNum = []int{}
	// var pn = true

	for i := 2; i <= num; i++ {
		var pn = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				pn = false
			}
		}
		if pn {
			arrPrNum = append(arrPrNum, i)
		}
	}
	return arrPrNum
}
