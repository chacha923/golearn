package lang

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	age      int8
	language string
}

// 更新私有字段
func UpdatePrivateField() {
	p := Programmer{"stefno", 18, "go"}
	fmt.Println(p)
	// 看来，内存对齐是不需要在计算地址	偏移量的时候考虑的
	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int8(0)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"
	fmt.Println(p)
}

// 读取切片的长度和容量
func LenOfSlice() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
}
