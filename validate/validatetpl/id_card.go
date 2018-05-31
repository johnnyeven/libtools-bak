package validatetpl

import (
	"regexp"
)

const (
	InvalidIDCardNoType  = "身份证号类型错误"
	InvalidIDCardNoValue = "无效的身份证号"
)

var (
	id_card_regexp = regexp.MustCompile(`^\d{17}(\d|x|X)$`)
)

func ValidateIDCardNo(v interface{}) (bool, string) {
	s, ok := v.(string)
	if !ok {
		return false, InvalidIDCardNoType
	}
	if !id_card_regexp.MatchString(s) {
		return false, InvalidIDCardNoValue
	}
	return true, ""
}

func ValidateIDCardNoOrEmpty(v interface{}) (bool, string) {
	s, ok := v.(string)
	if !ok {
		return false, InvalidIDCardNoType
	}
	if s != "" && !id_card_regexp.MatchString(s) {
		return false, InvalidIDCardNoValue
	}
	return true, ""
}
