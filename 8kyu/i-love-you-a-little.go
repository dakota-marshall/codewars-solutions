package kata

func HowMuchILoveYou(num_petals int) string {

	response := map[int]string{
		1: "I love you",
		2: "a little",
		3: "a lot",
		4: "passionately",
		5: "madly",
		6: "not at all",
	}

	for num_petals > 6 {
		num_petals -= 6
	}

	return response[num_petals]
}
