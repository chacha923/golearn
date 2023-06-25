package container

// 一个容器类应当实现这些方法
type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}
