package kata

import "strconv"

func CountBits(number uint) int {

	binary_string := strconv.FormatInt(int64(number), 2)

	var binary_sum int

	for _, value := range binary_string {
		if value == 49 {
			binary_sum += 1
		}
	}

	return binary_sum
}
