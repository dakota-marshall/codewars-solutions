package kata

func Invert(arr []int) []int {

	var inverse_int []int

	for _, num := range arr {
		inverse_int = append(inverse_int, num*-1)
	}

	return inverse_int
}
