package utils

func StrReverse(str string) (result string) {
	for _, char := range str {
		result += string(char)
	}

	return result
}