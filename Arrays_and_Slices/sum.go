package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// length := len(numbersToSum)
	// allSum := make([]int, length)

	var allSum []int
	for _, numbers := range numbersToSum {
		allSum = append(allSum, Sum(numbers))
	}
	return allSum
}

func SumAllTails(numbersToTailSum ...[]int) []int {
	// length := len(numbersToSum)
	// allSum := make([]int, length)

	var allTailSum []int
	for _, numbers := range numbersToTailSum {
		if len(numbers) == 0 {
			allTailSum = append(allTailSum, 0)
		} else {
			allTailSum = append(allTailSum, Sum(numbers[1:]))
		}
	}
	return allTailSum
}
