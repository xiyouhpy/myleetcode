package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/3sum/
// 题目描述：给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组
// 解题思路：先对数组进行排序，然后从左边小到大遍历，如果当前遍历到的值大于 0，则后续的值均会大于 0 无法组成符合预期的三元组；
//         若当前值不大于 0，则使用两个下标对数组剩下的元素分别从小到大和从大到小遍历，遍历中满足和为 0 的三元组则记录下来；
//         需要注意的是数组中可能存在相同的元素，导致遍历过程中产生相同的三元组
func main() {
	var arrList []int
	var retList [][]int

	// 结果如下：[[-1 -1 2] [-1 0 1]]
	arrList = []int{-1, 0, 1, 2, -1, -4}
	retList = threeSum(arrList)
	fmt.Println(retList)

	// 结果如下：[]
	arrList = []int{}
	retList = threeSum(arrList)
	fmt.Println(retList)

	// 结果如下：[]
	arrList = []int{0}
	retList = threeSum(arrList)
	fmt.Println(retList)
}

func threeSum(nums []int) [][]int {
	numLen := len(nums)
	if numLen < 3 {
		return [][]int{}
	}

	// 使用选择排序对 nums 进行排序
	for i := 0; i < numLen; i++ {
		for j := i + 1; j < numLen; j++ {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	retList := make([][]int, 0)
	for i := 0; i < numLen-2; i++ {
		// 数组已排序，排除三元组全为整数或者负数的情况
		if nums[i] > 0 {
			break
		}
		// 数组中可能存在重复的数字，会导致产出的结果存在重复的三元组，这里需要对后第一个元素 i 做去重逻辑
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 判断并找出合适的三元组
		var start, end = i + 1, numLen - 1
		for start < end {
			if nums[i]+nums[start]+nums[end] < 0 {
				start++
				continue
			}
			if nums[i]+nums[start]+nums[end] > 0 {
				end--
				continue
			}

			// 记录已找到的元素，执行该部分逻辑则肯定是 nums[i] + nums[start] + nums[end] = 0
			retList = append(retList, []int{nums[i], nums[start], nums[end]})

			// 分别移动 start 和 end 下标，进行下一轮操作处理
			start, end = start+1, end-1

			// 数组中可能存在重复的数字，会导致产出的结果存在重复的三元组，这里需要对后两个元素 start、end 做去重逻辑
			for start < end && nums[start] == nums[start-1] {
				start++
			}
			for start < end && nums[end] == nums[end+1] {
				end--
			}
		}
	}

	return retList
}
