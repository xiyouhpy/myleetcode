package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/longest-palindromic-substring/
// 题目描述：给定一个字符串s，求出返回最长的回文子串
// 解题思路：中心扩展法；遍历整个字符串的所有元素，分别以每个元素为中心向两边等步长扩展判断是否相等，若相等则该中心继续扩展，若不想等则该中心不再继续扩展（需要注意回文子串分别存在奇数和偶数情况）
func main() {
	// 测试case，第一个string表示参数，第二个string表示预期结果
	testList := map[string]string{
		"cbbd":  "bb",
		"babad": "bab",
		"a":     "a",
		"ac":    "a",
	}

	for k, v := range testList {
		result := longestPalindrome(k)
		if v != result {
			fmt.Println("failed! case params:", k, ", Expected:", v, ", Output:", result)
			return
		}
	}

	fmt.Println("success!")
}

// longestPalindrome 最大回文子串——中心扩展法
func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	var max, start, end = 0, 0, 0
	for i := 0; i < sLen; i++ {
		// 对奇数情况做扩展查找
		max1, start1, end1 := centerFind(s, i, i)

		// 对偶数情况做扩展查找
		max2, start2, end2 := centerFind(s, i, i+1)

		// 比较得出较大的长度，以及对应的左右下标
		if max2 > max1 && max2 > max {
			max = max2
			start = start2
			end = end2
		}
		if max1 > max2 && max1 > max {
			max = max1
			start = start1
			end = end1
		}
	}

	return s[start : end+1]
}

// centerFind 中心扩展查找逻辑
func centerFind(s string, left int, right int) (int, int, int) {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}

	// 减2是因为上面条件不满足退出时left--过且right++过
	return right - left + 1 - 2, left + 1, right - 1
}
