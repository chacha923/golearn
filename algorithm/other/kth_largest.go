package other

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	// 建一个最小堆长度k, 返回堆顶元素
	q := PriorityQueueInt{}
	for i := range nums {
		if len(q) <= k {
			heap.Push(&q, nums[i])
		} else {
			//如果nums[i] > 堆顶,  把堆顶pop, 插入num[i]
			if nums[i] > q[q.Len()-1] {
				heap.Pop(&q)
				heap.Push(&q, nums[i])
			}
		}
	}
	result := heap.Pop(&q).(int)
	return result
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueueInt []int

func (pq PriorityQueueInt) Len() int {
	return len(pq)
}

func (pq PriorityQueueInt) Less(i, j int) bool {
	// 注意：因为golang中的heap是按最小堆组织的，所以count越大，less()越小，越靠近堆顶.
	return pq[i] < pq[j]
}

func (pq PriorityQueueInt) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueueInt) Push(e any) {
	*pq = append(*pq, e.(int))
}

// Pop define
func (pq *PriorityQueueInt) Pop() any {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}
