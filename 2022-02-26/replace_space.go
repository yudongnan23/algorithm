package _022_02_26

// 替换字符串空格
// leetcode链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

/**
 * 思路：对字符串进行遍历，如果遇到的字符为空格，将其替换为指定字符
 */

const _replaceStr = "%20"

func replaceSpace(s string) string {
	var newS string

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			newS = newS + _replaceStr
			continue
		}

		newS = newS + string(s[i])
	}

	return newS
}
