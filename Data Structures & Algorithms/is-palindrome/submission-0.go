func isPalindrome(s string) bool {
	filters := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, s)

	ls := strings.ToLower(filters)
	l, r := 0, len(ls)-1

	for l < r {
		if ls[l] != ls[r] {
			return false
		}
		l++
		r--
	}

	return true
}
