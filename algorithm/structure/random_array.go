package structure

import "math/rand"

// 常数时间插入、删除和获取随机元素
type RandomizedSet struct {
	nums       []int       // 存储元素的值
	valToIndex map[int]int // 记录每个元素对应在 nums 中的索引
}

/** 如果 val 不存在集合中，则插入并返回 true，否则直接返回 false */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	// 若 val 不存在，插入到 nums 尾部，
	// 并记录 val 对应的索引值
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** 如果 val 在集合中，则删除并返回 true，否则直接返回 false */
func (this *RandomizedSet) Remove(val int) bool {
	// 若 val 不存在，不用再删除
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}
	// RandomizedSet 只能随机读，所以修改元素位置无所谓
	// 先拿到 val 的索引
	index := this.valToIndex[val]
	// 将最后一个元素对应的索引修改为 index
	this.valToIndex[this.nums[len(this.nums)-1]] = index
	// 交换 val 和最后一个元素
	this.nums[index], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[index]
	// 在数组中删除元素 val
	this.nums = this.nums[:len(this.nums)-1]
	// 删除元素 val 对应的索引
	delete(this.valToIndex, val)
	return true
}

/** 从集合中等概率地随机获得一个元素 */
func (this *RandomizedSet) GetRandom() int {
	// 随机获取 nums 中的一个元素
	return this.nums[rand.Intn(len(this.nums))]
}
