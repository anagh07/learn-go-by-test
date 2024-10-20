package arrays_and_slices

func ArraySum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(lists ...[]int) []int {
	sums := make([]int, len(lists))
	for i, list := range lists {
		sums[i] = ArraySum(list)
	}
	return sums
}

func SumAllTails(lists ...[]int) []int {
	var sums []int
	for _, list := range lists {
		if len(list) == 0 {
			sums = append(sums, 0)
		} else {
			tail := list[1:]
			sums = append(sums, ArraySum(tail))
		}
	}
	return sums
}
