package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/plus-one/
// 题目描述：给定由整数组成的非空数组所表示的非负整数，在该数的基础上加一；最高位数字存放数组的首位， 数组每个元素只存一个数字；除 0 外，该数不会以零开头。
// 解题思路：需要注意溢出进位情况
func main() {
	fmt.Printf("digits:1, 2, 3, result:%+v\n", plusOne([]int{1, 2, 3}))
	fmt.Printf("digits:4, 3, 2, 1, result:%+v\n", plusOne([]int{4, 3, 2, 1}))
	fmt.Printf("digits:9, 9, 9, 9, result:%+v\n", plusOne([]int{9, 9, 9, 9}))
	fmt.Printf("digits:0, result:%+v\n", plusOne([]int{0}))
}

func plusOne(digits []int) []int {
	// 判断最低位，若 !=9，则 +1 后不进位，则当前值 +1 后返回（其他位置不变化）
	digitsLen := len(digits)
	if digits[digitsLen-1] != 9 {
		digits[digitsLen-1]++
		return digits
	}

	// 遍历数组，若当前值 =9，则 +1 会进位，将当前值赋为 0，继续判断；若当前值 !=9，则 +1 不进位，返回结果
	digits[digitsLen-1] = 0
	for i := digitsLen - 2; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}

	// 处理溢出情况
	digits = append([]int{1}, digits...)

	return digits
}
