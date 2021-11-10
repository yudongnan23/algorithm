package _0211109

/**
 *  0～n-1中缺失的数字: https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
 */

func missingNumber(nums []int) int {
	return sum(len(nums)) - sliceSum(nums)
}

func sum(n int)int {
	return (n + 1) * n/2
}

func sliceSum(nums []int)int{
	s := 0
	for _, n := range nums {
		s = s + n
	}
	return s
}

func missingNumberMethod2(nums []int)int{
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	var r int
	for i := range nums {
		if nums[i] != i{
			r = i
			break
		}
	}

	return r
}