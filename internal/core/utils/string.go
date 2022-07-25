package utils

func Left(str string, length int) string {
	if len(str) < length {
		panic("length argument can not be bigger than given string's length!")
	}

	return str[len(str)-length:]
}

func Right(str string, length int) string {
	if len(str) < length {
		panic("length argument can not be bigger than given string's length!")
	}

	return str[:length]
}

// Gets a substring of a string from beginning of the string if it exceeds maximum length.
func Truncate(str string, maxLength int) string {
	if str == "" {
		return str
	}

	if len(str) <= maxLength {
		return str
	}
	return Left(str, maxLength)
}
