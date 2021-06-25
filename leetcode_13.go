package main

import "fmt"

// 题目链接：https://leetcode.com/problems/roman-to-integer/
// 题目描述：给定一个罗马数字，将其转换成整数
// 解题思路：建立罗马字母和数字的hashmap，遍历字符串并比较相邻的两位值，若后一位比前一位char表示的数字小则表示这两位一起为，否则一个char表示一个数字
func main() {
	if 3 != romanToInt("III") {
		fmt.Println("failed! case params:III")
		return
	}

	if 4 != romanToInt("IV") {
		fmt.Println("failed! case params:IV")
		return
	}

	if 9 != romanToInt("IX") {
		fmt.Println("failed! case params:IX")
		return
	}

	if 58 != romanToInt("LVIII") {
		fmt.Println("failed! case params:LVIII")
		return
	}

	if 1994 != romanToInt("MCMXCIV") {
		fmt.Println("failed! case params:LVIII")
		return
	}

	fmt.Println("success!")
}

func romanToInt(s string) int {
	// 建立 hashmap
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	i, retInt := 0, 0

	// 这里遍历时不处理最后一个字符
	for i = 0; i < len(s)-1; i++ {
		tmp1, tmp2 := s[i], s[i+1]
		if romanMap[tmp1] < romanMap[tmp2] {
			i++
			retInt += romanMap[tmp2] - romanMap[tmp1]
		} else {
			retInt += romanMap[tmp1]
		}
	}

	// 当下标 i == len(s)-1，表示上面最后一次判断是经过 else 逻辑后退出的，因此需要把最后一个元素的值加起来
	// 当下标 i > len(s)-1，上面上面最后一次判断是经过 if 逻辑后退出的，那么最后一个值已经被处理过，无需进行操作
	if i == len(s)-1 {
		retInt += romanMap[s[len(s)-1]]
	}

	return retInt
}
