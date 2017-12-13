package sort

//计数排序
//
//计数排序的过程类似小学选班干部的过程,如某某人10票,作者9票,那某某人是班长,作者是副班长
//大体分两部分,第一部分是拉选票和投票,第二部分是根据你的票数入桶

func CountSort(array []int) []int {
	if array == nil {
		return nil
	}

	length := len(array)
	count := make([]int, length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if array[j] < array[i] {
				count[i] = count[i] + 1
			}
		}
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		rank := count[i]
		result[rank] = array[i]
	}
	return result
}
