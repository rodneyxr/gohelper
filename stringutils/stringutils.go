package stringutils

// Reverse - Reverses a string.
func Reverse(s string) string {
	c := []rune(s)
	n := len(c)
	for i := 0; i < n/2; i++ {
		c[i], c[n-1-i] = c[n-1-i], c[i]
	}
	return string(c)
}
