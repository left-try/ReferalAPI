package utils

import "strconv"

func GenerateCode(userId int64) string {
	code := strconv.Itoa(int(userId)) + "_referral_code"
	return code
}
