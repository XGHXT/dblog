// Package tools provides ...
package tools

import (
	"github.com/sirupsen/logrus"
	"regexp"
)

var regexpEmail = regexp.MustCompile(`^(\w)+([\.\-]\w+)*@(\w)+((\.\w+)+)$`)

// ValidateEmail 校验邮箱
func ValidateEmail(e string) bool {
	return regexpEmail.MatchString(e)
}

var regexpPhoneNo = regexp.MustCompile(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)

// ValidatePhoneNo 校验手机号
func ValidatePhoneNo(no string) bool {
	logrus.Info(no)
	return regexpPhoneNo.MatchString(no)
}

// ValidatePassword 校验米阿莫
func ValidatePassword(pwd string) bool {
	return len(pwd) > 5 && len(pwd) < 32
}
