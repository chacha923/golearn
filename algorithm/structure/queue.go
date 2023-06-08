package structure

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
