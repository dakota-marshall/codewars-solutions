package kata

func Feast(beast string, dish string) bool {

	var result bool

	if beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1] {
		result = true
	} else {
		result = false
	}

	return result

}
