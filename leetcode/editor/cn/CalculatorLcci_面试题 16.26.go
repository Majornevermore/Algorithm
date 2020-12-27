package main

import (
	"strconv"
	"strings"
)

//给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
//
// 表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。
//
// 示例 1:
//
// 输入: "3+2*2"
//输出: 7
//
//
// 示例 2:
//
// 输入: " 3/2 "
//输出: 1
//
// 示例 3:
//
// 输入: " 3+5 / 2 "
//输出: 5
//
//
// 说明：
//
//
// 你可以假设所给定的表达式都是有效的。
// 请不要使用内置的库函数 eval。
//
// Related Topics 字符串
// 👍 23 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	ss := strings.Replace(s, " ", "", -1) //字符串去空
	str := strings.SplitAfter(ss, "")     //转数组
	stack := make([]int, 0)
	c := "+"
	num := 0
	for i := 0; i <= len(str); i++ {
		if isDigit(str, i) {
			tmp, _ := strconv.Atoi(str[i])
			num = num*10 + tmp
		} else {
			switch c {
			case "+":
				stack = append(stack, num)
				break
			case "-":
				stack = append(stack, -num)
				break
			case "*":
				num *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
				break
			case "/":
				num = stack[len(stack)-1] / num
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
				break
			}
			num = 0
			if i < len(str) {
				c = str[i]
			}

		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func isDigit(s []string, i int) bool { //判断是否为数字
	if i >= len(s) {
		return false
	}
	_, err := strconv.Atoi(s[i])
	return err == nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
}
