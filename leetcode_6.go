package main

import "fmt"

// 题目链接：https://leetcode.com/problems/zigzag-conversion/
// 题目描述：给定一个字符串，和一个目标行数数字；然后按照倒着的"z"分布排列，最后按照水平方向读取字符串则为目标字符串
// 解题思路：观察图示可发现，数据的变化有对称性，两遍遇到的转折点都是各自的对称点；相互对称的位置互为结果字符串相邻的位置
func main() {
	var strZigzag string
	var numRows int

	strZigzag = "PAYPALISHIRING"
	numRows = 3
	if "PAHNAPLSIIGYIR" != convert(strZigzag, numRows) {
		fmt.Println("failed! case params:", strZigzag, numRows)
		return
	}

	strZigzag = "PAYPALISHIRING"
	numRows = 4
	if "PINALSIGYAHRPI" != convert(strZigzag, numRows) {
		fmt.Println("failed! case params:", strZigzag, numRows)
		return
	}

	strZigzag = "A"
	numRows = 1
	if "A" != convert(strZigzag, numRows) {
		fmt.Println("failed! case params:", strZigzag, numRows)
		return
	}

	fmt.Println("success!")
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	arrConvert := make([]string, numRows)

	// 在转折点进行反转，将同一行的字符聚在该行数组上，最终得到每个下标值都是当前行的字符集合
	var i, flag = 0, -1
	for _, val := range s {
		arrConvert[i] += string(val)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	// 最后将数组按照下标依次连接，就可以得到完整的字符串数据
	var strNew string
	for index := 0; index < numRows; index++ {
		strNew += arrConvert[index]
	}

	return strNew
}
