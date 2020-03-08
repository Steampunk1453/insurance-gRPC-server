package utils

import (
	"regexp"
	"strings"
)

func GetNumberFormat(number string) string {
	re := regexp.MustCompile("[0-9]+")
	numberArray := re.FindAllString(number, -1)
	return strings.Join(numberArray, "")
}


