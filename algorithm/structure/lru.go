package structure

// hashmap + 双向链表 实现lru
// hashmap 用来保存数据
// 双向链表保存数据的访问顺序
// 1、如果我们每次默认从链表尾部添加元素，那么显然越靠尾部的元素就是最近使用的，越靠头部的元素就是最久未使用的。
// 2、对于某一个 key，我们可以通过哈希表快速定位到链表中的节点，从而取得对应 val。
// 3、链表显然是支持在任意位置快速插入和删除的，改改指针就行。只不过传统的链表无法按照索引快速访问某一个位置的元素，而这里借助哈希表，可以通过 key 快速映射到任意一个链表节点，然后进行插入和删除。
type LRUCache[K comparable, V any] struct {
	// 最大容量
	capacity int
	// key -> Node(key, val)
	hashMap map[K]*DoubleNode[K, V] // 存储数据，key：双向链表节点 key
	// Node(k1, v1) <-> Node(k2, v2)...
	cache *DoubleList[K, V]
}

func NewLRUCache[K comparable, V any](cap int) *LRUCache[K, V] {
	lruCache := &LRUCache[K, V]{
		capacity: cap,
	}
	lruCache.hashMap = make(map[K]*DoubleNode[K, V], cap)
	// TODO init double list
	lruCache.cache = NewDoubleList[K, V]()
	return lruCache
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	if v, ok := l.hashMap[key]; ok {
		// 读到了，将该数据提升为最近使用的
		l.makeRecently(key)
		return v.Val, true
	}
	var zeroVal V
	return zeroVal, false
}

func (l *LRUCache[K, V]) Put(key K, val V) {
	if _, exist := l.hashMap[key]; !exist {
		// key不存在，先插，再维护
		var node = NewDoubleNode[K, V](key, val)
		l.hashMap[key] = node

		// 容量满了，先淘汰
		if l.cache.Size() == l.capacity {
			l.removeLeastRecently()
		}
		l.addRecently(key, val)
	} else {
		// key 存在，删除旧的数据
		l.deleteKey(key)
		// 新插入的数据为最近使用的数据
		l.addRecently(key, val)
	}

}

// 将某个 key 提升为最近使用的
func (this *LRUCache[K, V]) makeRecently(key K) {
	x := this.hashMap[key]
	// 先从链表中删除这个节点
	this.cache.Remove(x)
	// 重新插到队尾
	this.cache.PushBack(x)
}

// 添加最近使用的元素
func (this *LRUCache[K, V]) addRecently(key K, val V) {
	x := NewDoubleNode[K, V](key, val)
	// 链表尾部就是最近使用的元素
	this.cache.PushBack(x)
	// 别忘了在 map 中添加 key 的映射
	this.hashMap[key] = x
}

// 删除某一个 key
func (this *LRUCache[K, V]) deleteKey(key K) {
	x, ok := this.hashMap[key]
	if !ok {
		return
	}
	// 从链表中删除
	this.cache.Remove(x)
	// 从 map 中删除
	delete(this.hashMap, key)
}

// 删除最久未使用的元素
func (this *LRUCache[K, V]) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deletedNode := this.cache.RemoveFirst()
	// 同时别忘了从 map 中删除它的 key
	deletedKey := deletedNode.Key
	delete(this.hashMap, deletedKey)
}
