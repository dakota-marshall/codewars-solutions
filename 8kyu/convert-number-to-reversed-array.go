package kata

func Digitize(n int) []int {

	rev_list := []int{}

	if n == 0 {
		return []int{0}
	}

	for n > 0 {
		rev_list = append(rev_list, n%10)
		n /= 10
	}

	return rev_list
}
