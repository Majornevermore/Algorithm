package main

import (
	"fmt"
	"strconv"
	"strings"
)

//实现一个基本的计算器来计算一个简单的字符串表达式的值。
//
// 字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格 。
//
// 示例 1:
//
// 输入: "1 + 1"
//输出: 2
//
//
// 示例 2:
//
// 输入: " 2-1 + 2 "
//输出: 3
//
// 示例 3:
//
// 输入: "(1+(4+5+2)-3)+(6+8)"
//输出: 23
//
// 说明：
//
//
// 你可以假设所给定的表达式都是有效的。
// 请不要使用内置的库函数 eval。
//
// Related Topics 栈 数学
// 👍 319 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func calculate1(s string) int {
	ss := strings.Replace(s, " ", "", -1) //字符串去空
	str := strings.SplitAfter(ss, "")     //转数组
	stack := make([]int, 0)
	c := "+"
	num := 0
	for len(str) > 0 {
		s := string(str[0])
		str = str[1:]
		if isDigit1(s) {
			tmp, _ := strconv.Atoi(s)
			num = num*10 + tmp
		} else {
			if s == "(" {
				num = calculate1(s)
			}
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

			c = s

		}
		if s == ")" {
			break
		}
	}
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func isDigit1(s string) bool { //判断是否为数字
	_, err := strconv.Atoi(s)
	return err == nil
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(calculate1("1+2"))

}
