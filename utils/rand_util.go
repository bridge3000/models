package utils

import (
	"math/rand"
)

type RandUtil struct {
}

func (this RandUtil) RandInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
