package structure

// 用切片模拟一个栈，下标大的元素靠近栈顶
type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) Len() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

// 注意，下标小的是栈底
func (s *Stack) ToSlice() []int {
	if s == nil {
		return []int{}
	}
	return s.data
}
