package structure

// 定义栈接口，可以用来写伪代码
type IStack[T any] interface {
	// 进栈
	Push(v T)
	// 出栈
	Pop() T
	// 栈长度
	Len() int
	// 栈是否为空
	Empty() bool
	// 观察栈顶
	Top() T
	// 观察栈底
	Bottom() T
	// 清空栈
	Clear()
	// 转化为切片
	ToSlice() []T
}

// 用切片模拟一个栈，下标大的元素靠近栈顶
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Bottom() T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	return s.data[0]
}

func (s *Stack[T]) Clear() {
	for s.Len() > 0 {
		s.Pop()
	}
}

// 注意，下标小的是栈底
func (s *Stack[T]) ToSlice() []T {
	if s == nil {
		return []T{}
	}
	return s.data
}
