package er

import (
	"regexp"
	"strconv"
)

/**
 * 纯数字
 */
func IsAllNumber(s string) (bool, error) {
	return regexp.MatchString("^[0-9]*$", s)
}

/**
 * N位数字
 */
func IsNNumber(s string, n int) (bool, error) {
	nStr := strconv.Itoa(n)
	return regexp.MatchString("^\\d{"+nStr+"}$", s)
}
