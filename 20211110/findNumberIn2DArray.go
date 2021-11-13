package _0211110

import (
	"fmt"
)

/**
 * 二维数组中的查找： https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
 */

// 二分查找
func findNumberIn2DArrayMethod2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	jstart := 0
	jend := len(matrix[0]) - 1
	istart := 0
	iend := len(matrix) - 1

	for istart < len(matrix) && iend >= 0 && jstart < len(matrix[0]) && jend >= 0 {
		if matrix[istart][jstart] == target || matrix[iend][jend] == target || matrix[istart][jend] == target || matrix[iend][jstart] == target {
			return true
		}

		if matrix[istart][jstart] > target || matrix[iend][jend] < target {
			return false
		}

		if matrix[istart][jend] < target {
			istart = istart + 1
			continue
		}

		if matrix[iend][jstart] < target {
			jstart = jstart + 1
			continue
		}

		if matrix[iend][jend] > target && matrix[istart][jend] > target {
			jend = jend - 1
			continue
		}

		if matrix[iend][jend] > target && matrix[iend][jstart] > target {
			iend = iend - 1
			continue
		}
	}
	return false
}

func searchRange(matrix [][]int, target int) (int, int, bool) {
	l := 0
	r := len(matrix)
	cope := make([][]int, 0)

	for l <= r {
		cope = append(cope, []int{l, r})
		m := (r + l) / 2
		if matrix[m][0] == target {
			return 0, 0, true
		}

		if matrix[m][0] > target {
			r = m - 1
		}

		if matrix[m][0] < target {
			l = m + 1
		}
	}

	return cope[len(cope)-1][0], cope[len(cope)-1][1], false
}

// 深度优先遍历
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	_format := "%d:%d"
	stack := [][]int{{0, 0}}
	dump := make(map[string]bool)
	dump[fmt.Sprintf(_format, 0, 0)] = true

	for len(stack) != 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if matrix[c[0]][c[1]] == target {
			return true
		}

		if _, ok := dump[fmt.Sprintf(_format, c[0]+1, c[1])]; !ok && c[0]+1 < len(matrix) && matrix[c[0]+1][c[1]] < target {
			stack = append(stack, []int{c[0] + 1, c[1]})
			dump[fmt.Sprintf(_format, c[0]+1, c[1])] = true
		}

		if _, ok := dump[fmt.Sprintf(_format, c[0], c[1]+1)]; !ok && c[1]+1 < len(matrix[0]) && matrix[c[0]][c[1]+1] < target {
			stack = append(stack, []int{c[0], c[1] + 1})
			dump[fmt.Sprintf(_format, c[0], c[1]+1)] = true
		}
	}
	return false
}
