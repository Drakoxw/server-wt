package utils

import (
	"fmt"
	"strconv"
)

func ToStringInterface(x interface{}) string {
	return fmt.Sprint(x)
}

func ToStringint(x int) string {
	return strconv.Itoa(x)
}

func ToStringlong(x int64) string {
	return strconv.FormatInt(x, 10)
}

func ToStringdouble(x float64) string {
	return fmt.Sprintf("%f", x)
}

func ToStringboolean(x bool) string {
	if x {
		return "true"
	}
	return "false"
}
