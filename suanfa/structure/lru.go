package structure

import "golearn/suanfa/list"

// hashmap + 双向链表 实现lru
type LRUCache struct {
	limit   int
	HashMap map[int]*list.DequeNode
	head    *list.DequeNode
	end     *list.DequeNode
}

func Constructor(cap int) LRUCache {
	lruCache := LRUCache{limit: cap}
	lruCache.HashMap = make(map[int]*list.DequeNode, cap)
	return lruCache
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Val
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, val int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head) //删链表
			delete(l.HashMap, key)         //删hash表
		}
		node := &list.DequeNode{Key: key, Val: val}
		l.addNode(node)
		l.HashMap[key] = node
	} else {
		v.Val = val
		l.refreshNode(v)
	}
}

// 刷新, 把一个节点插到链表头
func (l *LRUCache) refreshNode(node *list.DequeNode) {
	if node == nil {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

// 删除节点
func (l *LRUCache) removeNode(node *list.DequeNode) int {
	if node == l.end {
		l.end = l.end.pre
		l.end.next = nil
	} else if node == l.head {
		l.head = l.head.Next
		l.head.pre = nil
	} else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
	return node.Key
}

// 插到最前头
func (l *LRUCache) addNode(node *list.DequeNode) int {
	if l.end != nil {
		l.end.next = node
		node.Pre = l.end
		node.Next = nil
	}
	l.end = node
	if l.nead == nil {
		l.head = node
	}
}
