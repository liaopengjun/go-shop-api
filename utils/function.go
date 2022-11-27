package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenOrderNo() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 4; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	return timestamp + sb.String()
}

func In_array(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
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

func RemoveTopStructNew(fields map[string]string) string {
	var res string
	for _, err := range fields {
		if len(res) == 0 {
			res = err
			break
		}
	}
	return res
}

func MD5V(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
