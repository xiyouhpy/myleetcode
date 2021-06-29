package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/integer-to-roman/
// 题目描述：给定一个整数数字，将其转换成罗马数字
// 解题思路：枚举所有数字和罗马数字可能的映射关系，并建立罗马字母和数字的hashmap，遍历并找到对应关系得到结果
//         需要注意：1、golang中的map映射是无序的，无法顺序遍历；
//                 2、存在数字拆解为多个相同的罗马串，所以下面代码中第二层for是兼容这种情况
func main() {
	if "III" != intToRoman(3) {
		fmt.Println("failed! case params:3")
		return
	}

	if "IV" != intToRoman(4) {
		fmt.Println("failed! case params:4")
		return
	}

	if "IX" != intToRoman(9) {
		fmt.Println("failed! case params:9")
		return
	}

	if "LVIII" != intToRoman(58) {
		fmt.Println("failed! case params:58")
		return
	}

	if "CD" != intToRoman(400) {
		fmt.Println("failed! case params:400")
		return
	}

	if "MCMXCIV" != intToRoman(1994) {
		fmt.Println("failed! case params:1994")
		return
	}

	fmt.Println("success!")
}

func intToRoman(num int) string {
	// 建立 hashmap，列举可能的数字可能组成的基本单元
	romanMap := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intMap := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	for i := 0; i < len(romanMap); i++ {
		if num < intMap[i] {
			continue
		}

		// 若一个数字拆解为好几个同一个罗马串，比如 300 可以拆解为三个 C，即为：CCC；因此这个时候 i 的下标不能移动，这层 for 就是为了保证这种情况
		for ; num >= intMap[i]; num -= intMap[i] {
			result = result + romanMap[i]
		}
	}

	return result
}
