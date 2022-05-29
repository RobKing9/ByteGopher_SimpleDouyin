package utils

import "math/rand"

// RandNum returns a random number between min and max
func RandRangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)

}
