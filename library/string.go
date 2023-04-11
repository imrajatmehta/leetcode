package library

func IsOwelUpper(c byte) bool {
	if c == 'U' || c == 'O' || c == 'I' || c == 'E' || c == 'A' {
		return true
	}
	return false
}
func IsOwelLower(c byte) bool {
	if c == 'u' || c == 'o' || c == 'i' || c == 'e' || c == 'a' {
		return true
	}
	return false
}
func IsOwel(c byte) bool {
	return IsOwelLower(c) || IsOwelUpper(c)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
