package utils

import "math/rand"

func RandRangeIn(low, hi int) int {

	return low + rand.Intn(hi-low)

}
