package lang

import (
	"fmt"
	"runtime"
)

func deleteMapKey() {
	map1 := make(map[string]int)
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("插入 k-v 之前", m.Alloc)
	for i := 0; i < 10000; i++ {
		map1[fmt.Sprintf("%d", i)] = i
	}
	m = &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("插入 k-v 之后", m.Alloc)
	for i := 0; i < 10000; i++ {
		delete(map1, fmt.Sprintf("%d", i))
	}
	runtime.GC()
	m = &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("GC 之后", m.Alloc)
	fmt.Println("插入 k-v 之前", m.Alloc)
	for i := 0; i < 10000; i++ {
		map1[fmt.Sprintf("%d", i)] = i
	}
	m = &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("插入 k-v 之后", m.Alloc)
	runtime.GC()
	m = &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("GC 之后", m.Alloc)
	runtime.GC()
	m = &runtime.MemStats{}
	runtime.ReadMemStats(m)
	fmt.Println("GC 之后", m.Alloc)
}
