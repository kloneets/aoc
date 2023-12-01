package helpers

import (
	"log"
	"strconv"
)

func StringToNumberString(s string) string {
	if s == "0" {
		log.Println("Zerrrrrrrro")
	}
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