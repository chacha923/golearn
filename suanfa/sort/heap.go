package sort

import "fmt"

//堆排序
//如果是实现升序排列, 使用大根堆
//将完全二叉树(堆)用数组表示, 初始下标为0, 那么i节点的左右孩子下标为2i+1, 2i+2
//大根堆堆特性:  一个完全二叉树, 任何一个父节点 >= 右孩子 与 左孩子 最大值

//https://studygolang.com/articles/3719
//https://www.cnblogs.com/outerspace/p/11098461.html
//根据上面的代码修改, 不改变less函数
func RunHeapSort() {
	MakeHeap()
	HeapSort()
	fmt.Println(heap)
}

//构建堆, 初始化操作, 初始化后堆的最后一个元素有序
func MakeHeap() {
	n := len(heap)
	for i := n/2 - 1; i >= 0; i-- { //从最后一个非叶子节点开始处理, 最后一个非叶子节点的下标为n/2 - 1
		downRecur(i, n)
	}
}

func HeapSort() {
	for i := len(heap) - 1; i > 0; i-- {
		//移除顶部元素到数组末尾,然后剩下的重建堆,依次循环, 每循环一次, 堆的末尾有一个元素变为有序, 数组从后往前逐渐有序
		swap(0, i)
		downRecur(0, i)
	}
}

//每当一个顶部元素与末尾交换, 进行一次down操作, 称为堆化
// i: 操作的索引, n: 堆大小
func down(i, n int) {
	for {
		j1 := left(i)          //找i节点左孩子
		if j1 >= n || j1 < 0 { // 越界退出
			break
		}
		//找出两个节点中最小的(less: a<b)
		j := j1                                     //中间变量, 记录左右孩子 当中最大的节点
		if j2 := right(i); j2 < n && less(j1, j2) { //i节点右孩子j2
			j = j2
		}
		//此时j记录左右孩子中较大的节点, 如果i节点比左右孩子大, 那么满足堆, 不需要交换
		if less(j, i) {
			break
		}
		swap(i, j) // 否则i和较大的孩子交换
		i = j      // 如果发生交换, 那么新的子节点可能不满足堆, 继续向下做down操作
	}
}

// down操作递归版本
func downRecur(i, n int) {
	l := left(i)
	r := right(i)
	max := i
	if l < n && less(max, l) {
		max = l
	}
	if r < n && less(max, r) {
		max = r
	}
	if max != i {
		Swap(heap, i, max)
		downRecur(max, n)
	}
}

/////////////////////////////////////////////////////////////////////////////////////////
//判断下标a元素是否比下标b小
func less(a, b int) bool {
	return heap[a] < heap[b]
}

//交换下标a与下标b元素
func swap(a, b int) {
	heap[a], heap[b] = heap[b], heap[a]
}

func Swap(slice []int, a, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}

// 严格小于
func Less(slice []int, a, b int) bool {
	return slice[a] < slice[b]
}

////////////////////////////////////下面通常不用///////////////////////////////
//由子节点到父节点重新开始建堆
func up(j int) {
	for {
		i := (j - 1) / 2 //得到父节点
		if i == j || less(j, i) {
			//less(子,父) !less(9,5) == true
			//父节点小于子节点,符合最小堆条件,break
			break
		}
		//子节点比父节点大,互换
		swap(i, j)
		j = i
	}
}

func Push(x interface{}) {
	heap = append(heap, x.(int))
	up(len(heap) - 1)
	return
}

func Pop() interface{} {
	n := len(heap) - 1
	swap(0, n)
	down(0, n)

	old := heap
	n = len(old)
	x := old[n-1]
	heap = old[0 : n-1]
	return x
}

func Remove(i int) interface{} {
	n := len(heap) - 1
	if n != i {
		swap(i, n)
		down(i, n)
		up(i)
	}
	return Pop()
}

/**
父节点i的左子节点在位置 (2i+1)
父节点i的右子节点在位置 (2i+2)
子节点i的父节点在位置  floor((i-1)/2)
*/
func RunHeapSort2() {
	m := len(heap)
	s := len(heap) / 2 //最后一个非叶子节点的索引
	for i := s; i > -1; i-- {
		heapSort2(heap, i, m-1)
	}
	for i := m - 1; i > 0; i-- {
		swap(i, 0)
		heapSort2(heap, 0, i-1)
	}
	fmt.Println(heap)
}

func heapSort2(arr []int, i, end int) {
	l := 2*i + 1 //左孩子
	if l > end {
		return
	}
	n := l
	r := 2*i + 2 //右孩子
	if r <= end && arr[r] > arr[l] {
		n = r
	}
	if arr[i] > arr[n] {
		return
	swap(i, n)
	heapSort2(arr, n, end)
}
