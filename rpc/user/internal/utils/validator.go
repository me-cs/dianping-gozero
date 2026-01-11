package utils

import "regexp"

const (
	// 手机号正则
	PhoneRegex = `^1([38][0-9]|4[579]|5[0-3,5-9]|6[6]|7[0135678]|9[89])\d{8}$`
	// 邮箱正则
	EmailRegex = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	// 密码正则。4~32位的字母、数字、下划线
	PasswordRegex = `^\w{4,32}$`
	// 验证码正则, 6位数字或字母
	VerifyCodeRegex = `^[a-zA-Z\d]{6}$`
)

// IsPhoneInvalid 判断手机号是否无效
func IsPhoneInvalid(phone string) bool {
	if phone == "" {
		return true
	}
	matched, _ := regexp.MatchString(PhoneRegex, phone)
	return !matched
}

// IsEmailInvalid 判断邮箱是否无效
func IsEmailInvalid(email string) bool {
	if email == "" {
		return true
	}
	matched, _ := regexp.MatchString(EmailRegex, email)
	return !matched
}

// IsCodeInvalid 判断验证码是否无效
func IsCodeInvalid(code string) bool {
	if code == "" {
		return true
	}
	matched, _ := regexp.MatchString(VerifyCodeRegex, code)
	return !matched
}