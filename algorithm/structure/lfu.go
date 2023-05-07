package structure

import "golearn/suanfa/list"

// 缓存算法设计并实现数据结构。它应该支持以下操作：get 和 put。

// get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
// put(key, value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。
// 当缓存达到其容量时，则应该在插入新项之前，使最不经常使用的项无效。
// 在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除最久未使用的键。

// 我们定义两个哈希表，第一个 freq_table 以频率 freq 为索引，每个索引存放一个双向链表，这个链表里存放所有使用频率为 freq 的缓存，
// 缓存里存放三个信息，分别为键 key，值 value，以及使用频率 freq。
// 第二个 key_table 以键值 key 为索引，每个索引存放对应缓存在 freq_table 中链表里的内存地址，
// 这样我们就能利用两个哈希表来使得两个操作的时间复杂度均为 O(1)。
// 同时需要记录一个当前缓存最少使用的频率 minFreq，这是为了删除操作服务的。
type LFUCache struct {
	cache               map[int]*list.FreqDequeNode
	freq                map[int]*DoubleList
	ncap, size, minFreq int
}

// 双向链表结构, 包含头尾指针
type DoubleList struct {
	head *list.FreqDequeNode
	tail *list.FreqDequeNode
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache: make(map[int]*Node),
		freq:  make(map[int]*DoubleList),
		ncap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// 查到了, 修改freq
		// todo
		return node.Val
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if ncap == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
		// key存在, 修改值
		node.Val = value
		incrFreq(node)
	} else {
		// key 不存在
		// 超出容量, 删除低频节点
		if this.size >= this.ncap {
			node := this.freq[this.minFreq].RemoveLast()
			delete(this.cache, node.Key)
			this.size--
		}
		x := &Node{key: key, val: value, freq: 1}
		this.cache[key] = x
		if this.freq[1] == nil {
			this.freq[1] = CreateDL()
		}
		this.freq[1].AddFirst(x)
		this.minFreq = 1
		this.size++
	}
}

func (this *LFUCache) incrFreq(node *list.FreqDequeNode) {
	f := node.Freq
	this.freq[f].Remove(node) // 对应freq表删除
	// 最小频率的双向表空了, 删掉
	if this.minFreq == f && this.freq[f].IsEmpty() {
		this.minFreq++
		delete(this.freq, f)
	}
	node.Freq++
	// 建表
	if this.freq[node.Freq] == nil {
		this.freq[node.freq] = CreateDL()
	}
	this.freq[node.freq].AddFirst(node)
}

func CreateDL() *DoubleList {
	head := &FreqDequeNode{}
	tail := &FreqDequeNode{}
	return &DoubleList{
		head: head,
		tail: tail,
	}
}

func (this *DoubleList) AddFirst(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.next.Prev = node
	this.head.next = node
}

func (this *DoubleList) Remove(node *list.DequeNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	node.Next = nil
	node.Prev = nil
}

func (this *DoubleList) RemoveLast() *list.DequeNode {
	if this.IsEmpty() {
		return nil
	}

	last := this.tail.Prev
	this.Remove(last)

	return last
}

func (this *DoubleList) IsEmpty() bool {
	return this.head.next == this.tail
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
