package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	return str
}

func FileLines(fileName string) []string {
	data := strings.Split(strings.ReplaceAll(ReadFile(fileName), "\r\n", "\n"), "\n")
	if data[len(data)-1] == "" { //remove empty line if exists
		data = data[:len(data)-1]
	}

	return data
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func UniqueStrings(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
