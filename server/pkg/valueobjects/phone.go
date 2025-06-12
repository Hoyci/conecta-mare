package valueobjects

import "regexp"

func SanitizePhoneNumber(phone string) (string, bool) {
	re := regexp.MustCompile(`\D`)
	onlyDigits := re.ReplaceAllString(phone, "")

	if len(onlyDigits) == 10 || len(onlyDigits) == 11 {
		return onlyDigits, true
	}

	return "", false
}
