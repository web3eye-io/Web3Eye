package utils

import "strconv"

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Str2uint64(str string) (uint64, error) {
	baseNum := 16
	targetNum := 64
	i, err := strconv.ParseUint(str, baseNum, targetNum)
	return i, err
}
