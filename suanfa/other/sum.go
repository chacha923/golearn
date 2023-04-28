package other

// 查找两数之和，返回下标
func twoSum(nums []int, target int) []int {
	var hashTable = make(map[int]int)
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
