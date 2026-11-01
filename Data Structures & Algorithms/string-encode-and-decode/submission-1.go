type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for i, str := range strs {
		if i == len(strs) {
			encoded += str
		}
		encoded += str + "ก"
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	result := make([]string, 0)
	str := ""
	for _, c := range encoded {
		if c == rune('ก') {
			result = append(result, str)
			str = ""
		} else {
			str += string(c)
		}
	}
	return result
}
