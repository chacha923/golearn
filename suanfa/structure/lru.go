package structure

import "golearn/suanfa/list"

// hashmap + 双向链表 实现lru
// hashmap 用来保存数据
// 双向链表保存数据的访问顺序
// 最近访问的数据放在链表头，最少访问的数据放在链表尾。链表节点中存储键值对信息。(这里粗暴 直接存链表节点的指针)
// 使用 HashMap 来存储键值对信息以及对应的节点在链表中的位置。
type LRUCache struct {
	capacity int                         // 容量
	size     int                         // 当前大小
	hashMap  map[string]*list.DoubleNode // 存储数据，key：双向链表节点 key
	head     *list.DoubleNode            // dummy 节点永远指向链表头，移动时不用考虑边界
	tail     *list.DoubleNode            // dummy 节点永远指向链表尾
}

func NewLRUCache(cap int) *LRUCache {
	var head = list.NewEmptyDoubleNode()
	var tail = list.NewEmptyDoubleNode()
	lruCache := &LRUCache{
		capacity: cap,
		size:     0,
		head:     head,
		tail:     tail,
	}
	lruCache.hashMap = make(map[string]*list.DoubleNode, cap)

	return lruCache
}

func (l *LRUCache) Get(key string) int {
	if v, ok := l.hashMap[key]; ok {
		// 读到了，顺序移到追前面
		l.moveToHead(v)
		return v.Val
	}
	return -1
}

func (l *LRUCache) Put(key string, val int) {
	if v, ok := l.hashMap[key]; !ok {
		// key不存在，先插，再维护
		var node = list.NewDoubleNodeWithKey(key, val)
		l.hashMap[key] = node
		l.size += 1

		// 超了，删除最旧的
		if l.size > l.capacity {
			toRemoved := l.removeTail()      //删链表
			delete(l.hashMap, toRemoved.Key) //删hash表
			l.size -= 1
		}
	} else {
		// key 存在，刷新
		v.Val = val
		l.moveToHead(v)
	}
}

// 每次 put 或 get 已存在的key，刷新, 把一个节点插到链表头
func (l *LRUCache) moveToHead(node *list.DoubleNode) {
	if node == nil {
		return
	}
	l.removeNode(node)
	l.addToHead(node)
}

// 将新节点插到最前头
func (l *LRUCache) addToHead(node *list.DoubleNode) {
	node.Prev = l.head
	node.Next = l.head.Next
	l.head.Next.Prev = node
	l.head.Next = node
}

// 删除节点
func (l *LRUCache) removeNode(node *list.DoubleNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 删除最后一个节点，并返回
func (l *LRUCache) removeTail() *list.DoubleNode {
	node := l.tail.Prev
	l.removeNode(node)
	return node
}
