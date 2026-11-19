func longestConsecutive(nums []int) int {
	res := 0
	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	run := 0

	for len(m) > 0 {
		for k, _ := range m {
			run = k
			break
		}
		delete(m, run)

		first, last, count := run-1, run+1, 1

		for {
			if _, ok := m[first]; !ok {
				break
			}
			delete(m, first)
			first--
			count++
		}
		for {
			if _, ok := m[last]; !ok {
				break
			}
			delete(m, last)
			last++
			count++
		}

		res = max(res, count)
	}

	return res
}
