package kata

func Summation(n int) int {
	working_total := 0

	for num := 0; num <= n; num++ {
		working_total += num
	}

	return working_total

}
