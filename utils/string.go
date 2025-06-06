package utils

// CountChars counts the number of occurrences of a character in a string.
func CountChars(Str string, Ch rune) (n int) {
	for _, ch := range Str {
		if ch == Ch {
			n++
		}
	}
	return
}
