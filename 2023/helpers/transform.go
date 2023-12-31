package helpers

import (
	"strconv"
)

func StringToNumberString(s string) string {
	numberMap := make(map[string]string)
	numberMap["one"] = "1"
	numberMap["two"] = "2"
	numberMap["three"] = "3"
	numberMap["four"] = "4"
	numberMap["five"] = "5"
	numberMap["six"] = "6"
	numberMap["seven"] = "7"
	numberMap["eight"] = "8"
	numberMap["nine"] = "9"
	_, err := strconv.Atoi(s)
	if err != nil {
		return numberMap[s]
	}
	return s
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
