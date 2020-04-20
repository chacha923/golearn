package list

type LRUCache struct {
	limit   int
	HashMap map[int]*DequeNode
	head    *DequeNode
	tail    *DequeNode
}

func Constructor(cap int) LRUCache {
	lruCache := LRUCache{limit: cap}
	lruCache.HashMap = make(map[int]*DequeNode, cap)
	return lruCache
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.limit {
			oldKey := l.removeNode(l.head)
			delete(l.HashMap, key)
		}
		node := &DequeNode{Key: key, Value: value}
		l.addNode(node)
		h.HashMap[key] = node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

func (l *LRUCache) refreshNode(node *DequeNode) {
	if node == nil {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

// 删除节点
func (l *LRUCache) removeNode(node *DequeNode) int {
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
func (l *LRUCache) addNode(node *DequeNode) int {
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
