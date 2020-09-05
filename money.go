package er

import "regexp"

/**
 * 是否是货币格式
 */
func IsMoney(s string) (bool, error) {
	return regexp.MatchString("^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$", s)
}
