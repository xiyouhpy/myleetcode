package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/container-with-most-water/
// 题目描述：给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水
// 解题思路：我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。
//         我们会使用变量 maxArea 来持续存储到目前为止所获得的最大面积。 在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxArea，并将指向较短线段的指针向较长线段那端移动一步
func main() {
	var height []int
	var area int

	height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area = maxArea(height)
	if area != 49 {
		fmt.Println("49 failed! case params:", area)
		return
	}

	height = []int{1, 1}
	area = maxArea(height)
	if area != 1 {
		fmt.Println("1 failed! case params:", area)
		return
	}

	height = []int{4, 3, 2, 1, 4}
	area = maxArea(height)
	if area != 16 {
		fmt.Println("16 failed! case params:", area)
		return
	}

	height = []int{1, 2, 1}
	area = maxArea(height)
	if area != 2 {
		fmt.Println("2 failed! case params:", area)
		return
	}

	fmt.Println("success!")
}

func maxArea(height []int) int {
	getMaxArea := 0
	start, end := 0, len(height)-1
	for start < end {
		tempArea := (end - start) * getMin(height[start], height[end])
		getMaxArea = getMax(tempArea, getMaxArea)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}

	return getMaxArea
}

func getMin(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
