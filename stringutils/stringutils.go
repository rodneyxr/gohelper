package stringutils

// Reverse - Reverses a string.
func Reverse(s string) string {
	c := []rune(s)
	n := len(c)
	stop := n / 2
	for i := 0; i < stop; i++ {
		tmp := n - 1 - i
		c[i], c[tmp] = c[tmp], c[i]
	}
	return string(c[n:])
}
