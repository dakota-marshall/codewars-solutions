package kata

func LoveFunc(flower1, flower2 int) bool {

	mod_flower1 := flower1 % 2
	mod_flower2 := flower2 % 2

	if (mod_flower1 == 0 || mod_flower2 == 0) && !(mod_flower1 == 0 && mod_flower2 == 0) {
		return true
	}

	return false
}
