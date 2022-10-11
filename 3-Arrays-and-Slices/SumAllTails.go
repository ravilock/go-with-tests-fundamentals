package Sum

func SumAllTails(numbersToSum ...[]int) []int {
	numberOfSlices := len(numbersToSum)
	tailsSum := make([]int, numberOfSlices)

	for index, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailsSum[index] = 0
			continue
		}
		tail := numbers[1:]
		tailsSum[index] = Sum(tail)
	}

	return tailsSum
}
