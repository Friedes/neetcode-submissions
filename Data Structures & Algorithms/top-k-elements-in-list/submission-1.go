import (
	"slices"
	"cmp"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums {
		m[num]++
	}

	items := make([]Item,0)
	for key ,v := range m {
		items = append(items, Item{N: key, F: v})
	}

	slices.SortFunc(items, func(a, b Item) int {
		if n := cmp.Compare(b.F, a.F); n != 0 {
			return n
		}
		return cmp.Compare(a.N, b.N)
	})

	for i:=0;i<k;i++{
		result = append(result, items[i].N)
	}

	return result
}

type Item struct {
	N int
	F int
}