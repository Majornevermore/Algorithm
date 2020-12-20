package main

import "math/rand"

//给定一个包含 [0，n ) 中独特的整数的黑名单 B，写一个函数从 [ 0，n ) 中返回一个不在 B 中的随机整数。
//
// 对它进行优化使其尽量少调用系统方法 Math.random() 。
//
// 提示:
//
//
// 1 <= N <= 1000000000
// 0 <= B.length < min(100000, N)
// [0, N) 不包含 N，详细参见 interval notation 。
//
//
// 示例 1:
//
//
//输入:
//["Solution","pick","pick","pick"]
//[[1,[]],[],[],[]]
//输出: [null,0,0,0]
//
//
// 示例 2:
//
//
//输入:
//["Solution","pick","pick","pick"]
//[[2,[]],[],[],[]]
//输出: [null,1,1,1]
//
//
// 示例 3:
//
//
//输入:
//["Solution","pick","pick","pick"]
//[[3,[1]],[],[],[]]
//Output: [null,0,0,2]
//
//
// 示例 4:
//
//
//输入:
//["Solution","pick","pick","pick"]
//[[4,[2]],[],[],[]]
//输出: [null,1,3,1]
//
//
// 输入语法说明：
//
// 输入是两个列表：调用成员函数名和调用的参数。Solution的构造函数有两个参数，N 和黑名单 B。pick 没有参数，输入参数是一个列表，即使参数为空，
//也会输入一个 [] 空列表。
// Related Topics 排序 哈希表 二分查找 Random
// 👍 51 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	sz      int
	mapping map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	l := Solution{
		mapping: map[int]int{},
	}

	l.sz = N - len(blacklist)

	for i := 0; i < len(blacklist); i++ {
		l.mapping[blacklist[i]] = 999
	}
	last := N - 1
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= l.sz {
			continue
		}
		for {
			if _, ok := l.mapping[last]; ok {
				last--
			} else {
				break
			}

		}
		l.mapping[blacklist[i]] = last
		last--
	}
	return l
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.sz)
	if _, ok := this.mapping[index]; ok {
		return this.mapping[index]
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	l := Constructor(4, []int{2, 1})
	for i := 0; i < 3; i++ {
		l.Pick()
	}
}
