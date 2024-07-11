package util

import (
	"fmt"
	"regexp"
)

func Match(pattern string, str string) bool {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return regex.MatchString(str)
}
