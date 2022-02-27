package _022_02_27

// 统计重复数字在排序数组中出现的次数
// leetcode链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

/**
 * 思路：二分查找
 */

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	var start int

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			start = mid
			break
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}

	return newStatCounter(nums, target, start).statLeftCount() +
		newStatCounter(nums, target, start+1).statRightCount()
}

type statCounter struct {
	nums   []int
	target int
	start  int
}

func newStatCounter(nums []int, target int, start int) *statCounter {
	return &statCounter{
		nums:   nums,
		target: target,
		start:  start,
	}
}

func (s *statCounter) statLeftCount() int {
	var count int

	for i := s.start; i >= 0; i-- {
		if s.nums[i] == s.target {
			count++
		}
	}

	return count
}

func (s *statCounter) statRightCount() int {
	var count int

	for i := s.start; i < len(s.nums); i++ {
		if s.nums[i] == s.target {
			count++
		}
	}

	return count
}
