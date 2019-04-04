package sort

import "fmt"

var data = []int{278, 109, 63, 930, 589, 184, 505, 269, 8, 83}

func RunRadixSort() {
	radixSort()
	fmt.Println(data)
}

/**
 * 基数排序:先将低位数字用计数排序排好,再依次对高位进行排序
 */
func radixSort() {
	llen := len(data)
	bits := maxbit(arr) //数组中最大元素的位数
	tmp := make([]int, llen)
	count := make([]int, 10) //计数器
	bitValue := 0
	radix := 1

	for i := 1; i <= bits; i++ { //进行d次排序
		for j := 0; j < 10; j++ {
			count[j] = 0 //每次分配前清空计数器
		}
		for j := 0; j < llen; j++ {
			bitValue = (data[j] / radix) % 10 //统计每个桶中的记录数
			count[bitValue]++
		}
		for j := 1; j < 10; j++ {
			count[j] = count[j-1] + count[j] //将tmp中的位置依次分配给每个桶
		}
		for j := llen - 1; j >= 0; j-- {
			bitValue = (data[j] / radix) % 10
			tmp[count[bitValue]-1] = data[j]
			count[bitValue]--
		}
		for j := 0; j < llen; j++ { //将临时数组的内容复制到data中
			data[j] = tmp[j]
		}
		radix = radix * 10
	}

}

//辅助函数，求数组中最大元素的位数,即数组最大位数
func maxbit(data []int) int {
	llen := len(arr)
	maxData := data[0]
	for i := 1; i < llen; i++ {
		if maxData < data[i] {
			maxData = data[i]
		}
	}
	bits := 1
	for maxData > 0 {
		maxData /= 10
		bits++
	}
	return bits
}
