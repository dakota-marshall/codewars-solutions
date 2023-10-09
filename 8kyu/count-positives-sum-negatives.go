package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	positive_count := 0
	negative_sum := 0

	for count := 0; count < len(numbers); count++ {
		number := numbers[count]

		if number > 0 {
			positive_count++
		} else if number < 0 {
			negative_sum += number
		}

	}

	res = append(res, positive_count)
	res = append(res, negative_sum)

	return res
}
