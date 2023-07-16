package structure

type IQueue[T any] interface {
	// 进队
	Push(v T)
	// 出队
	Pop() T
	// 队列长度
	Len() int
	// 观察队头节点
	Peek() T
}

// 用切片模拟一个普通的 FIFO 队列
// 下标 0 为队头，下标 len(data)-1 为队尾
type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() T {
	if len(q.data) == 0 {
		var zero T
		return zero
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue[T]) Peek() T {
	if len(q.data) == 0 {
		var zero T
		return zero
	}
	return q.data[0]
}

func (q *Queue[T]) Len() int {
	if q == nil {
		return 0
	}
	return len(q.data)
}

// 单调队列
// 既能够维护队列元素「先进先出」的时间顺序，又能够正确维护队列中所有元素的最值，这就是「单调队列」结构。
type MonotonicQueue struct {
	// 双链表，支持头部和尾部增删元素
	// 维护其中的元素自尾部到头部单调递增
	maxq []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		maxq: make([]int, 0),
	}
}

// 在队尾添加元素 n
// 如果每个元素被加入时都这样操作，最终单调队列中的元素大小就会保持一个单调递减的顺序
func (mq *MonotonicQueue) Push(n int) {
	// 将前面小于自己的元素都删除
	for len(mq.maxq) > 0 && mq.maxq[len(mq.maxq)-1] < n {
		mq.maxq = mq.maxq[:len(mq.maxq)-1]
	}
	mq.maxq = append(mq.maxq, n)
}

// 返回当前队列中的最大值
func (mq *MonotonicQueue) Max() int {
	// 队头的元素肯定是最大的
	return mq.maxq[0]
}

// 队头元素如果是 n，删除它
func (mq *MonotonicQueue) Pop(n int) {
	if n == mq.maxq[0] {
		mq.maxq = mq.maxq[1:]
	}
}
