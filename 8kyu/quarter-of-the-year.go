package kata

import (
	"math"
)

func QuarterOf(month int) int {

	var quarter float64
	var rounded_quarter int

	// Divide by 3 month interval, then round up
	quarter = float64(month) / 3
	rounded_quarter = int(math.Ceil(quarter))

	return rounded_quarter

}
