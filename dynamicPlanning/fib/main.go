package main

import (
	"fmt"
	"time"
)

var memo []int64

func fib(n int) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] == -1 {
		memo[n] = fib(n-1) + fib(n-2)
	}
	return memo[n]
}

func fibDp(n int) int64 {
	memo[1] = 1
	memo[2] = 1
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func main() {
	time1 := time.Now()
	n := 10
	memo = make([]int64, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}

	res := fibDp(n)
	time2 := time.Now().Sub(time1)
	fmt.Println(res, "time", time2)
}
