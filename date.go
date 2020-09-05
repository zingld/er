package er

import "regexp"

/**
 * 是否是日期格式
 */
func IsDateFormat(s string) (bool, error) {
	return regexp.MatchString("^\\d{4}-\\d{1,2}-\\d{1,2}", s)
}
