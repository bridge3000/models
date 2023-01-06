package utils

import (
	"fmt"
	"strconv"
)

type MathUtil struct {
}

func (this *MathUtil) F2i(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}
