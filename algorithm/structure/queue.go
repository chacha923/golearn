package structure

// 用切片模拟一个普通的 FIFO 队列
// 下标小的为队头，下标大的为队尾
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
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *Queue[T]) Len() int {
	if q == nil {
		return 0
	}
	return len(q.data)
}
