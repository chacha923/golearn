package sort

import (
	"container/list"
	"fmt"
)

//基数排序
//基于两种不同的排序顺序我们将基数排序分为
//LSD（Least significant digital）或  MSD（Most significant digital）
//LSD的排序方式由数值的最右边（低位）开始，而MSD则相反，由数值的最左边（高位）开始
//这里实现LSD
func CreateBucket(arrLen int) []list.List {
	bucket := make([]list.List, 0, 10)
	return bucket
}

//清空桶
func ClearBucket(bucket []list.List) {

}

//按桶内的顺序读取数组, 赋值给arr
func ReadBucket(arr []int, bucket []list.List) []int{
	i := 0
	for _,sublist := range bucket {
		pVal := sublist.Front()
		for index := 0; index < sublist.Len(); index++{
			arr[i] = pVal.Value.(int)
			i++
			pVal = pVal.Next()
		}
		sublist.Init()
	}
	return arr
}

func PutIn(bucket []list.List, arr []int, time int){
	for _, ele := range arr {
		NumOfBit := ele /(10^(time-1)) % (10^time)
		bucket[NumOfBit].PushBack(ele)
	}
}

//d 表示 arr 内元素最大的位数
func CardinalSort(arr []int, d int){
	lenArr := len(arr)
	bucket := CreateBucket(lenArr)
	for i := 1; i <= d ; i++{
		PutIn(bucket, arr, i)
		ReadBucket(arr, bucket)
	}
	fmt.Println(arr)
}