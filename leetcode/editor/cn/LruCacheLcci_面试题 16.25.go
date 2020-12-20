package main

import "fmt"

//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存
//被填满时，它应该删除最近最少使用的项目。
//
// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新
//的数据值留出空间。
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
// Related Topics 设计
// 👍 53 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size     int
	capacity int
	hashmap  map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key  int
	val  int
	pre  *DLinkedNode
	next *DLinkedNode
}

func initDLinkeNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		hashmap:  map[int]*DLinkedNode{},
		head:     initDLinkeNode(0, 0),
		tail:     initDLinkeNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.hashmap[key]; !ok {
		return -1
	}
	node := this.hashmap[key]
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCache) addHead(node *DLinkedNode) {
	//this.head.next = node
	//this.head.next.pre = node
	//node.pre = this.head
	//node.next = this.head.next
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.hashmap[key]; !ok {
		node := initDLinkeNode(key, value)
		this.size++
		this.hashmap[key] = node
		this.addHead(node)
		if this.size > this.capacity {
			remove := this.removeTail()
			delete(this.hashmap, remove.key)
			this.size--
		}
	} else {
		node := this.hashmap[key]
		node.val = value
		this.moveToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

func LRU(operators [][]int, k int) []int {
	lru := Constructor(k)
	var res []int
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lru.Put(operators[i][1], operators[i][2])
		} else if operators[i][0] == 2 {
			res = append(res, lru.Get(operators[i][1]))
		}
	}
	return res
}

func main() {
	var op = [][]int{{1, 1, 1}, {1, 2, 2}, {2, 1}, {1, 3, 3}, {2, 2}, {1, 4, 4}, {2, 1}, {2, 3}, {2, 4}}
	fmt.Print(LRU(op, 2))
}
