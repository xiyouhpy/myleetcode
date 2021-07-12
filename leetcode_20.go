package main

import "fmt"

// 题目链接：https://leetcode.com/problems/valid-parentheses/
// 题目描述：给定一个字符串，包含'(', ')', '[', ']', '{', '}'，判断字符串是否有效，左括号必须被右括号关闭
// 解题思路：声明新数组来记录配对情况，最大下标来判断是否配对，未配对则追加到数组后，配对则减去配对的元素，最终判断是否栈数组是否为空即可
func main() {
	if true != isValid("()") {
		fmt.Println("failed! case params:()")
		return
	}

	if true != isValid("()[]{}") {
		fmt.Println("failed! case params:()[]{}")
		return
	}

	if false != isValid("(]") {
		fmt.Println("failed! case params:(]")
		return
	}

	if false != isValid("([)]") {
		fmt.Println("failed! case params:([)]")
		return
	}

	if true != isValid("{[]}") {
		fmt.Println("failed! case params:{[]}")
		return
	}

	fmt.Println("success!")
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '}' {
			if len(res) != 0 && res[len(res)-1] == '{' {
				res = res[:len(res)-1]
			} else {
				return false
			}
		} else if s[i] == ')' {
			if len(res) != 0 && res[len(res)-1] == '(' {
				res = res[:len(res)-1]
			} else {
				return false
			}
		} else if s[i] == ']' {
			if len(res) != 0 && res[len(res)-1] == '[' {
				res = res[:len(res)-1]
			} else {
				return false
			}
		} else {
			res = append(res, s[i])
		}
	}

	if len(res) != 0 {
		return false
	} else {
		return true
	}
}
