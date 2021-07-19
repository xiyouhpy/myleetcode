package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// 题目描述：给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合，给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母
// 解题思路：排列组合问题；第一个循环对目标输入进行拆分排列，第二个、第三个循环对拆分的元素进行组合，然后将结果进行追加
func main() {
	var digits string
	var resultList []string

	// 正确结果为： ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	digits = "23"
	resultList = letterCombinations(digits)
	fmt.Println(resultList)

	// 正确结果为： []
	digits = ""
	resultList = letterCombinations(digits)
	fmt.Println(resultList)

	// 正确结果为： ["a", "b", "c"]
	digits = "2"
	resultList = letterCombinations(digits)
	fmt.Println(resultList)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	digitsMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	result := digitsMap[digits[0]]
	for i := 1; i < len(digits); i++ {
		var tmp1 []string
		for j := 0; j < len(result); j++ {
			tmp2 := digitsMap[digits[i]]
			for k := 0; k < len(tmp2); k++ {
				tmp1 = append(tmp1, result[j]+tmp2[k])
			}
		}
		result = tmp1
	}

	return result
}
