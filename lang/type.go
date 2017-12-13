package lang

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Type() {
	i := 1
	j := int64(1)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(unsafe.Sizeof(j))
}
