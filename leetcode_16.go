package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/3sum-closest/
// 题目描述：给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
//		   返回这三个数的和。假定每组输入只存在唯一答案
// 解题思路：先对 nums 数组进行排序，遍历 nums 的每个元素，我们记录为 num1，该数的索引为 index1，再用剩余的数组合与 target 的差值进行比较，
//		   那么剩余两个数从 nums 中的索引范围为 [index1, len(nums) - 1]
//         则 num2 从前往后遍历的位置为 start = index1 + 1； nums3 从后往前遍历的开始位置 end = len(nums) - 1；
//         与 target 比较的结果可能如下：
//         		1、差值为0，说明已找到目标组合，返回 target
//				2、差值小于0，则需要继续增加三个数的和，保存这三个值与 target 的差值到 numSub 数组，并进行start++
// 				3、差值大于0，则需要继续缩小三个数的和，保存这三个值与 target 的差值到 numSub 数组，并进行end--
//		   最终求出 numSub 元素最小的值，该值与 target 求和即为题目的目标值
func main() {
	var nums []int

	nums = []int{-1, 2, 1, -4}
	if 2 != threeSumClosest(nums, 1) {
		fmt.Println("failed! case params:", nums, ", target: 2")
		return
	}

	fmt.Println("success!")
}

func threeSumClosest(nums []int, target int) int {
	// 1、排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	// 2、求所有差值
	var numSub []int
	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			// 得出差值
			tempSub := nums[i] + nums[start] + nums[end] - target
			if tempSub == 0 {
				return target
			}
			if tempSub < 0 {
				start++
			}
			if tempSub > 0 {
				end--
			}
			// 保存差值
			numSub = append(numSub, tempSub)
		}
	}

	// 3、取最小差值，这里注意正数和负数，切记差值最小不是值最小！！！
	minSub := numSub[0]
	for i := 1; i < len(numSub); i++ {
		tempMinSub, tempNumSub := minSub, numSub[i]
		if minSub < 0 {
			tempMinSub = 0 - minSub
		}
		if numSub[i] < 0 {
			tempNumSub = 0 - numSub[i]
		}

		if tempMinSub > tempNumSub {
			minSub = numSub[i]
		}
	}

	return minSub + target
}
