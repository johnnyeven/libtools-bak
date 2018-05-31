package validatetpl

import (
	"regexp"
)

const (
	InvalidIPv4Type  = "IPv4类型错误"
	InvalidIPv4Value = "无效的IPv4"
)

var (
	ipv4_valid_regx = regexp.MustCompile(`^((\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])\.){3}(\d{1,2}|1\d{2}|2[0-4]\d|25[0-5])$`)
)

func ValidateIPv4(v interface{}) (bool, string) {
	s, ok := v.(string)
	if !ok {
		return false, InvalidIPv4Type
	}

	if !ipv4_valid_regx.MatchString(s) {
		return false, InvalidIPv4Value
	}
	return true, ""
}

func ValidateIPv4OrEmpty(v interface{}) (bool, string) {
	s, ok := v.(string)
	if !ok {
		return false, InvalidIPv4Type
	}

	if s != "" && !ipv4_valid_regx.MatchString(s) {
		return false, InvalidIPv4Value
	}
	return true, ""
}
