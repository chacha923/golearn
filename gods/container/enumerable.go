package container

// value 值可以枚举，基于下标
type EnumerableWithIndex[T any] interface {
	// 容器的每个元素，执行一次 op
	Each(op func(index int, value T))
}
