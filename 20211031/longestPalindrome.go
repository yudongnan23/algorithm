package main

/**
 * 最长回文子串：https://leetcode-cn.com/problems/longest-palindromic-substring/
 */
func longestPalindrome(s string) string {
	maxLengthStr := ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j >= i; j-- {
			if isPalindrome(s[i:j]) && (j-i) > len(maxLengthStr){
				maxLengthStr = s[i:j]
			}

			if len(maxLengthStr) >= (j - i - 1) {
				break
			}
		}

		if len(maxLengthStr) >= len(s)-i {
			break
		}
	}
	return maxLengthStr
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if string(s[i]) != string(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindromeMethod2(s string) string {
	i := 0
	j := len(s)
	maxLengthStr := ""
	for i < j{
		curMaxLengthStr := longestPalindromeWithMaxLength(s[i:j], 0)
		if len(curMaxLengthStr) > len(maxLengthStr){
			maxLengthStr = curMaxLengthStr
		}
		if len(maxLengthStr) > j - i {
			break
		}

		i++
		j--
	}
	return maxLengthStr

}

func longestPalindromeWithMaxLength(s string, curMaxLength int) string {
	isPalindrome, i := isPalindrome2(s)
	if isPalindrome && len(s) > curMaxLength {
		return s
	}

	leftMaxLengthStr := longestPalindromeWithMaxLength(s[:i+1], curMaxLength)
	rightMaxLengthStr := longestPalindromeWithMaxLength(s[i+1:], curMaxLength)
	if len(leftMaxLengthStr) > len(rightMaxLengthStr) {
		return leftMaxLengthStr
	}
	return rightMaxLengthStr
}

func isPalindrome2(s string) (bool, int) {
	i := 0
	j := len(s) - 1

	for i < j {
		if string(s[i]) != string(s[j]) {
			return false, i
		}
		i++
		j--
	}
	return true, 0
}