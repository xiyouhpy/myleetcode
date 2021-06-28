package main

import (
	"fmt"
	"strconv"
)

// 题目链接：https://leetcode.com/problems/add-binary/
// 题目描述：给定两个二进制字符串，返回它们的和（用二进制表示），输入为非空字符串且只包含数字 1 和 0。
// 解题思路：需要注意string和int的转换；是否有进位，长度是否有变化
func main() {
	var a, b string

	//a, b = "11", "1"
	//fmt.Printf("a=%s, b=%s, result:%s\n", a, b, addBinary(a, b))

	a, b = "1010", "1011"
	fmt.Printf("a=%s, b=%s, result:%s\n", a, b, addBinary(a, b))
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	if aLen == 0 {
		return b
	}
	if bLen == 0 {
		return a
	}

	// 进行对位相加，并保留进位
	strSum := ""
	flag := 0
	i, j := aLen-1, bLen-1
	for i >= 0 && j >= 0 {
		intAi, _ := strconv.Atoi(string(a[i]))
		intBj, _ := strconv.Atoi(string(b[j]))
		tempSum := intAi + intBj + flag
		flag = tempSum / 2
		strSum = strconv.Itoa(tempSum%2) + strSum
		i--
		j--
	}

	// 处理加数剩余长度
	for i >= 0 {
		intAi, _ := strconv.Atoi(string(a[i]))
		tempSum := intAi + flag
		flag = tempSum / 2
		strSum = strconv.Itoa(tempSum%2) + strSum
		i--
	}

	// 处理被加数剩余长度
	for j >= 0 {
		intBj, _ := strconv.Atoi(string(b[j]))
		tempSum := intBj + flag
		flag = tempSum / 2
		strSum = strconv.Itoa(tempSum%2) + strSum
		j--
	}

	// 补充最高进位值
	if flag == 1 {
		strSum = "1" + strSum
	}

	return strSum
}
