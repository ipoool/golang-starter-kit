package helpers

import "strconv"

// StringToInt - Convert string to integer
func StringToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}
