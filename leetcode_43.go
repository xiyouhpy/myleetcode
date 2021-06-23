package main

import "fmt"

// 题目链接：https://leetcode.com/problems/multiply-strings/
// 题目描述：给定两个非负整数字符串，进行大数据相乘
// 解题思路：大数据相乘，比大数据相加复杂；因为乘法需要进行交叉相乘，所以进位处理比较麻烦些；这里通过数组下标来进行进位处理比较巧妙
//         具体进位处理见下面注释，注意避免进位丢失的情况
func main() {
	var num1 string
	var num2 string

	num1 = "2"
	num2 = "3"
	if "6" != multiply(num1, num2) {
		fmt.Println("failed! case params:", num1, num2)
		return
	}

	num1 = "9"
	num2 = "9"
	if "81" != multiply(num1, num2) {
		fmt.Println("failed! case params:", num1, num2)
		return
	}

	num1 = "99"
	num2 = "99"
	if "9801" != multiply(num1, num2) {
		fmt.Println("failed! case params:", num1, num2)
		return
	}

	num1 = "123"
	num2 = "456"
	if "56088" != multiply(num1, num2) {
		fmt.Println("failed! case params:", num1, num2)
		return
	}

	fmt.Println("success!")
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}

	sumList := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		tmp1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			tmp2 := int(num2[j] - '0')

			// 进行各个位置相乘，并且需要加上之前结果的进位
			tmpSum := tmp1*tmp2 + sumList[i+j+1]

			// 整理当前位置，剔除进位的值后，保存起来
			sumList[i+j+1] = tmpSum % 10

			// 保存进位的值到下一位，这里一定要注意后面需要加 sumList[i+j]，防止覆盖上一层计算的值
			sumList[i+j] = tmpSum/10 + sumList[i+j]
		}
	}

	// 对计算的值进行转字符串处理，并剔除高位多余的0（该sumList为map，k为0的位置表示未使用）
	var result string
	for k, v := range sumList {
		// 这里不能去掉 v == 0 的判断，因为最高位的数组下标 k = 0，去掉的话会导致遗漏最高位的值
		if k == 0 && v == 0 {
			continue
		}
		result += string(v + '0')
	}

	return result
}
