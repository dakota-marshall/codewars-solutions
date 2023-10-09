package kata

func ExpressionMatter(a int, b int, c int) int {

	highest_value := 0

	plus_plus := a + b + c
	mult_mult := a * b * c
	plus_mult := a + b*c
	mult_plus := a*b + c
	parenplus_mult := (a + b) * c
	parenmult_mult := (a * b) * c
	mult_parenplus := a * (b + c)
	mult_parenmult := a * (b * c)
	plus_parenplus := a + (b + c)
	plus_parenmult := a + (b * c)

	values := []int{
		plus_plus, mult_mult, plus_mult, mult_plus, parenplus_mult,
		parenmult_mult, mult_parenplus, mult_parenmult, plus_parenplus, plus_parenmult,
	}

	for _, value := range values {
		if value > highest_value {
			highest_value = value
		}
	}

	return highest_value

}
