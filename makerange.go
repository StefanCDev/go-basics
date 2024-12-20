package piscine

func MakeRange(min, max int) []int {
	// it's the condition
	if min >= max {
		return nil
	}
	/* I create a variable using make, and the variable will range from max and min, for example 10
	is max, and 5 is min, then the array range in 5 value in between
	*/
	resultArr := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		// the result will loop start at 0, so 0 + 5 = 5, then 1 + 5 = 6, until 4 + 5 = 9 is the last value
		resultArr[i] = i + min
	}
	return resultArr
}
