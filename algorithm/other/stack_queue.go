package other

// 用两个栈实现队列
type CQueue struct {
	in  Stack
	out Stack
}

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) > 0 {
		return this.out.Pop()
	}
	if len(this.in) > 0 {
		for len(this.in) > 0 {
			this.out.Push(this.in.Pop())
		}
		return this.out.Pop()
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
