package lang

import (
	"fmt"

	"github.com/golang/glog"
)

func NewMap() map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, 0)
	m["foo"] = "bar"
	m["f11"] = "b11"
	return m
}

func NewArray() [3]int {
	var array [3]int
	array[0] = 0
	array[1] = 1
	array[2] = 2
	return array
}

func NewSlice() []int {
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	return slice
}

func RangeMap() {
	m := NewMap()
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func RangeSlice(slice []int) []int {
	for k := range slice {
		slice[k] = slice[k] + 1
	}
	return slice
}

func RangeArrayPoint(array [3]int) [3]int {
	glog.Errorln("RangeArrayPoint before range", array)
	for k, v := range &array {
		array[k] = v + 1
	}
	glog.Errorln(array)
	return array
}

func RangeArray(array [3]int) [3]int {
	glog.Errorln("RangeArray before range", array)
	for k, v := range array {
		array[k] = v + 1
	}
	glog.Errorln(array)
	return array
}
