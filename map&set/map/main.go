package main

import (
	"log"
	"runtime"
)

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v 分配内存= %vKB, GC次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}

func main() {
	printMemStats("初始化")
	hash := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		hash[i] = i
	}
	runtime.GC()
	printMemStats("增加map数据后")
	log.Println("删除前数组长度", len(hash))
	for i := 0; i < 10000; i++ {
		delete(hash, i)
	}
	log.Println("删除后数组长度：", len(hash))

	// 再次进行手动GC回收
	runtime.GC()
	printMemStats("删除map数据后")

	// 设置为nil进行回收
	hash = nil
	runtime.GC()
	printMemStats("设置为nil后")

}
