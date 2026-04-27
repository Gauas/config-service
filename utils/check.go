package utils

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func ToLower(r rune) rune {
	if IsUpper(r) {
		return r + ('a' - 'A')
	}
	return r
}

func ToUpper(r rune) rune {
	if IsLower(r) {
		return r - ('a' - 'A')
	}
	return r
}
