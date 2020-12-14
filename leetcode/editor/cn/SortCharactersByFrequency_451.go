package main

import "sort"

//给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
// 示例 1:
//
//
//输入:
//"tree"
//
//输出:
//"eert"
//
//解释:
//'e'出现两次，'r'和't'都只出现一次。
//因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//
//
// 示例 2:
//
//
//输入:
//"cccaaa"
//
//输出:
//"cccaaa"
//
//解释:
//'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
//注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//
//
// 示例 3:
//
//
//输入:
//"Aabb"
//
//输出:
//"bbAa"
//
//解释:
//此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
//注意'A'和'a'被认为是两种不同的字符。
//
// Related Topics 堆 哈希表
// 👍 192 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	hashMap := make(map[rune]int)
	ret := make([]rune, 0)
	res := make([]rune, 0)
	r := []rune(s)
	for _, v := range r {
		if st, ok := hashMap[v]; ok {
			hashMap[v] = st + 1
		} else {
			hashMap[v] = 1
			ret = append(ret, v)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return hashMap[ret[i]] > hashMap[ret[j]]
	})
	for i := 0; i < len(ret); i++ {
		for j := 0; j < hashMap[ret[i]]; j++ {
			res = append(res, ret[i])
		}
	}
	return string(res)

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	frequencySort("tree")
}
