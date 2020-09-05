package er

import "regexp"

/**
 * 国内电话号码
 */
func IsTelephoneNo(s string) (bool, error) {
	return regexp.MatchString("\\d{3}-\\d{8}|\\d{4}-\\d{7}", s)
}

/**
 * 是否是手机号
 */
func IsPhoneNo(s string) (bool, error) {
	return regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$", s)
}

/**
 * 是否是身份证
 * TODO 详细实现
 */
func IsID(s string) (bool, error) {
	return regexp.MatchString("^\\d{15}|\\d{18}$", s)
}
