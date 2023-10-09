package kata

import (
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {

	var character_list [14]string

	character_list[0] = "("
	character_list[4] = ")"
	character_list[5] = " "
	character_list[9] = "-"

	position := 1

	for _, number := range numbers {
	reeval:
		if character_list[position] == "" {
			character_list[position] = strconv.FormatUint(uint64(number), 10)
			position++
		} else {
			position++
			goto reeval
		}
	}

	return strings.Join(character_list[:], "")

}
