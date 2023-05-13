package util

func RemoveDuplicates(slice []int) []int {
	// 定义一个 map，用于存储每个元素是否已经出现过
	seen := make(map[int]bool)
	// 定义一个结果 slice
	result := []int{}
	// 遍历原 slice，将不重复的元素添加到结果 slice 中
	for _, value := range slice {
		if _, exists := seen[value]; !exists {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}
