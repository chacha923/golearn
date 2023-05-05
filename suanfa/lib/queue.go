package lib

// 用切片模拟一个普通的 FIFO 队列
// 下标小的为队头，下标大的为队尾
type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return -1
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
