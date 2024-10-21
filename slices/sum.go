package slice

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(a, b int) int { return a + b }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}

func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	result := initialValue
	for _, value := range collection {
		result = f(value, result)
	}

	return result
}
