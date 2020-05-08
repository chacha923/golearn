package sort

// 计数排序对输入的数据有附加的限制条件：
// 1、输入的线性表的元素属于有限偏序集S；
// 2、设输入的线性表的长度为n，|S|=k（表示集合S中元素的总数目为k），则k=O(n)。
// 在这两个条件下，计数排序的复杂性为O(n)。

func CountSort(array []int) {
	maxVal := 0
	for i := range array {
		if array[i] > maxVal {
			maxVal = array[i]
		}
	}
	tmp := make([]int, maxVal+1) // idx: 元素 value: 出现的次数
	for i := range array {
		tmp[array[i]]++
	}
	var j int
	for i := 0; i < len(tmp); i++ {
		for tmp[i] > 0 {
			arr[j] = i
			j++
			tmp[i]--
		}
	}
}

func RunCountSort() {
	CountSort(data)
}
