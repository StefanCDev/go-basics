package piscine

func AppendRange(min, max int) []int {
	// you create an array of int
	var resultArr []int
	// you create the loop to put the value inside the array, the minimum value is min, and go to max
	for i := min; i < max; i++ {
		// just condition
		if min >= max {
			return nil
		}
		// the result is included the resultArr and the value inside the loop
		resultArr = append(resultArr, i)
	}
	return resultArr
}
