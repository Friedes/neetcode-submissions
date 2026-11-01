func groupAnagrams(strs []string) [][]string {
	result := make([][]string,0)
	m := make(map[[26]int][]string)

	for _, str := range strs {
		var alps [26]int
		for _, c := range str {
			alps[c-'a']++
		}
		m[alps] = append(m[alps], str)
	}

	for _,v := range m {
		result = append(result, v)
	}

	return result
}
