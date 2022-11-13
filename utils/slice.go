package utils

import (
	"strconv"
	"strings"
)

func IsValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func StringToSlice(id string) (ids []int) {
	strids := strings.Split(id, ",")
	for _, id := range strids {
		if id == "" {
			continue
		}
		idint, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		ids = append(ids, idint)
	}
	return
}

func SubStrLen(str string, length int) string {
	nameRune := []rune(str)
	if len(str) > length {
		return string(nameRune[:length-1]) + "..."
	}
	return string(nameRune)
}
