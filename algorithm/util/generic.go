package util

import (
	"reflect"
)

// 函数无法直接返回泛型类型的默认值，因为 Golang 的类型系统是静态类型系统，需要在编译时确定类型，无法在运行时确定。
// 但是，可以通过在函数中使用类型断言，判断泛型类型的默认值类型并返回相应的默认值。
func DefaultValue[T any](v T) T {
	t := reflect.TypeOf(v)
	zero := reflect.Zero(t)
	return zero.Interface().(T)
}
