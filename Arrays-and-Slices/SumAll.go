package Sum

func SumAll(numbersToSum ...[]int) []int {
	numberOfSlices := len(numbersToSum)
	sums := make([]int, numberOfSlices)

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}

	return sums
}
