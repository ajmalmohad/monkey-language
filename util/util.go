package util

func IsLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func IsDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
