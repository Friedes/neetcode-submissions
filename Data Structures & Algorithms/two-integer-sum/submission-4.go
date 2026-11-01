func twoSum(nums []int, target int) []int {
    m := make(map[int][]int)
    for i:=0;i<len(nums);i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = []int{i}
        } else {
            m[nums[i]] = append(m[nums[i]], i)
        }
    }
    for i:=0;i<len(nums)-1;i++ {
        if v, ok := m[target-nums[i]]; ok {
            if nums[i] == target-nums[i] { 
                if len(v) > 1 {
                    return []int{i, v[1]}
                }
            } else {
                return []int{i, v[0]}
            }   
        }
    }

    return []int{}
}
