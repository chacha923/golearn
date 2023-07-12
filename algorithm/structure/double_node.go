package structure

// 双向链表节点
type DoubleNode[K comparable, V any] struct {
	// 数据域
	Key K // 键值对
	Val V
	// 指针域
	Next *DoubleNode[K, V] // 后一个
	Prev *DoubleNode[K, V] // 前一个
}

func NewDoubleNode[K comparable, V any](key K, val V) *DoubleNode[K, V] {
	return &DoubleNode[K, V]{Key: key, Val: val}
}

func NewEmptyDoubleNode[K comparable, V any]() *DoubleNode[K, V] {
	var zeroKey K
	var zeroVal V
	return NewDoubleNode[K, V](zeroKey, zeroVal)
}

// 带使用频率的双端队列节点, lfu
type FreqDequeNode[K comparable, V any] struct {
	Freq int
	DoubleNode[K, V]
}

// 双向链表接口
type IDoubleList[K comparable, V any] interface {
	PushBack(x *DoubleNode[K, V])
	Remove(x *DoubleNode[K, V])
	RemoveFirst() *DoubleNode[K, V]
	Size() int
}

type DoubleList[K comparable, V any] struct {
	head, tail *DoubleNode[K, V]
	size       int
}

func NewDoubleList[K comparable, V any]() *DoubleList[K, V] {
	// head tail 是 dummy 节点
	head := NewEmptyDoubleNode[K, V]()
	tail := NewEmptyDoubleNode[K, V]()
	head.Next, tail.Prev = tail, head
	return &DoubleList[K, V]{head: head, tail: tail, size: 0}
}

// 在链表尾部添加节点 x，时间 O(1)
func (d *DoubleList[K, V]) PushBack(x *DoubleNode[K, V]) {
	x.Prev = d.tail.Prev
	x.Next = d.tail
	d.tail.Prev.Next = x
	d.tail.Prev = x
	d.size += 1
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (d *DoubleList[K, V]) Remove(x *DoubleNode[K, V]) {
	// 不需要遍历链表，因为给的是目标节点
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
	d.size -= 1
}

// 删除链表中第一个节点，并返回该节点，时间 O(1)
func (d *DoubleList[K, V]) RemoveFirst() *DoubleNode[K, V] {
	if d.head.Next == d.tail {
		return nil
	}
	first := d.head.Next
	d.Remove(first)
	return first
}

// 返回链表长度，时间 O(1)
func (d *DoubleList[K, V]) Size() int {
	return d.size
}
