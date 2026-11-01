func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	bef := make([]int, n)
	aft := make([]int, n)
	
	bef[0], aft[n-1] = 1,1
	for i:=1;i<n;i++ {
		bef[i] = nums[i-1] * bef[i-1]
	}
	for i:=n-2;i>=0;i--{
		aft[i] = nums[i+1] * aft[i+1]
	}
	for i:=0;i<n;i++{
		res[i] = bef[i]*aft[i]
	}

	return res
}
