package _0211120

/**
 * 连续子数组的最大和: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
 */

func maxSubArray(nn []int) int {
	if len(nn) == 0 {
		return 0
	}

	d := make([]int, len(nn))
	d[0] = nn[0]
	max := nn[0]

	for i := 1; i < len(nn); i++ {
		sum := d[i-1] + nn[i]
		if sum > nn[i] {
			d[i] = sum
		}else{
			d[i] = nn[i]
		}

		if d[i] > max {
			max = d[i]
		}
	}
	return max
}
