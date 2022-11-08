package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ToStringInterface(x interface{}) string {
	return fmt.Sprint(x)
}

func ToStringIdPrimite(x interface{}) string {
	strng := ToStringInterface(x)
	strng = strings.Replace(strng, `ObjectID("`, "", 1)
	strng = strings.Replace(strng, `")`, "", 1)
	return strng
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
