package main

import (
	"fmt"
	"sort"
	"strconv"
)

func FourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && (nums[i]+nums[i+1]+nums[i+2]+nums[i+3]) < target; i++ {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n-2 && (nums[j]+nums[i+2]+nums[i+3]) < target-nums[i]; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ans
}

func treeSum(nums []int, targetNum int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for frist := 0; frist < n; frist++ {
		if frist >= 1 && nums[frist] == nums[frist-1] {
			continue
		}
		thrid := n - 1
		target := targetNum - nums[frist]
		for second := frist + 1; second < n; second++ {
			if second < frist+1 && nums[second-1] == nums[second] {
				continue
			}
			for nums[second]+nums[thrid] > target && thrid > second {
				thrid--
			}
			if second == thrid {
				break
			}
			if nums[second]+nums[thrid] == target {
				ans = append(ans, []int{nums[frist], nums[second], nums[thrid]})
			}
		}
	}
	return ans
}

func main() {
	ip := ""
	var ips []string
	for i := 0; i < 10; i++ {
		changedIp(&ip, i)
		//tmp := ip
		ips = append(ips, ip)
	}
	fmt.Println(ips)
}

func changedIp(ip *string, i int) {
	*ip = "1992" + strconv.Itoa(i)
}
