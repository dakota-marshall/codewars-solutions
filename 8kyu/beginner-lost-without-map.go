package kata

func Maps(input_list []int) []int {

	var doubled_list []int

	for count := 0; count < len(input_list); count++ {

		doubled_list = append(doubled_list, input_list[count]*2)

	}

	return doubled_list

}
