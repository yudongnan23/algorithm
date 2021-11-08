package _0211108

/**
 * 替换空格: https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
 */

func replaceSpace(s string)string {
	r := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			r = r + "%20"
			continue
		}

		r = r + string(s[i])
	}
	return r
}

