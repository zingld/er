package er

import (
	"regexp"
	"strconv"
)

/**
 * 汉字
 */
func IsChinese(s string) (bool, error) {
	return regexp.MatchString("^[\u4e00-\u9fa5]{0,}$", s)
}

/**
 * 英文和数字
 */
func IsEnglishAndNumber(s string) (bool, error) {
	return regexp.MatchString("^[A-Za-z0-9]+$", s)
}

/**
 * 长度为 m - n 的所有字符,暂不兼容汉字
 */
func IsLenBetweenMN(s string, m, n int) (bool, error) {
	mStr := strconv.Itoa(m)
	nStr := strconv.Itoa(n)
	return regexp.MatchString("^.{"+mStr+","+nStr+"}$", s)
}
