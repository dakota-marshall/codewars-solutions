package kata

func Past(hours, minutes, seconds int) int {
	var millis_past_midnight int
	const millis_in_hour = 3600000
	const millis_in_minute = 60000
	const millis_in_second = 1000

	millis_past_midnight += hours * millis_in_hour
	millis_past_midnight += minutes * millis_in_minute
	millis_past_midnight += seconds * millis_in_second

	return millis_past_midnight

}
