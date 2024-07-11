package util

import (
	"fmt"
	"strconv"
)

func StringToUint(str string) uint {
	uintStr, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err.Error())
	}
	return uint(uintStr)
}
