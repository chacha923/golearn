package container

// 迭代器，基于下标
type IteratorWithIndex[T any] interface {
	Next() bool
	Value() T
	Index() int
	Begin()
	First() bool
	NextTo(func(index int, value T) bool) bool
}
