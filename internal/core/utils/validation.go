package utils

import "unicode"

// at least one uppercase letter, one lowercase letter, one number, one special character

func ValidatePassword(value string, minLength int) bool {
	return verify(value, minLength, true, true, true, true)
}

func verify(value string, minLength int, requiredDigit, requiredLowercase, requiredUppercase, requiredSpecialCharacter bool) bool {

	var (
		upp, low, num, sym bool
	)
	if len(value) < minLength {
		return false
	}

	for _, r := range value {
		switch {
		case unicode.IsUpper(r):
			if !requiredUppercase {
				return false
			}
			upp = true
		case unicode.IsLower(r) && requiredLowercase:
			if !requiredLowercase {
				return false
			}
			low = true
		case unicode.IsNumber(r) && requiredDigit:
			if !requiredDigit {
				return false
			}
			num = true
		case (unicode.IsPunct(r) || unicode.IsSymbol(r)):
			if !requiredSpecialCharacter {
				return false
			}
			sym = true
		default:
			return false
		}
	}

	if (requiredDigit && !num) ||
		(requiredUppercase && !upp) ||
		(requiredLowercase && !low) ||
		(requiredSpecialCharacter && !sym) {
		return false
	}

	return true
}
