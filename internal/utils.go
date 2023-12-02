package internal

func IsRuneANumber(inRune rune) bool {
	if inRune >= '0' && inRune <= '9' {
		return true
	}
	return false
}
