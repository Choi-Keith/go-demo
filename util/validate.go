package util

import (
	"errors"
	"regexp"
)

func IsValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, _ := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, _ = regexp.MatchString("^[a-zA-Z]", username)
	if !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

func IsValidatePassword(password, rePassword string) error {
	if len(password) == 0 {
		return errors.New("请输入密码")
	}
	if len(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}

func IsValidateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("请输入邮箱")
	}
	pattern := `^([A-Za-z0-9_\.\-])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		return errors.New("邮箱格式不符合规范")
	}
	return nil
}
