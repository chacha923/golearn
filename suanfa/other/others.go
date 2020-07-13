package other

// 寻找完美数
// 对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”。
func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return false
	}
	var sum = 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum = sum + i + num/i
		}
	}
	if sum == num {
		return true
	}
	return false
}

// 和为k的子数组个数
// 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 我们定义 pre[i] 为 [0..i] 里所有数的和，则 pre[i] 可以由 pre[i−1] 递推而来，即：
// pre[i]=pre[i−1]+nums[i]
// 那么「[j..i] 这个子数组和为 kk 」这个条件我们可以转化为
// pre[i]−pre[j−1]==k
// 简单移项可得符合条件的下标 j 需要满足
// pre[j−1]==pre[i]−k
func subarraySum(nums []int, k int) int {
	count := 0         // 记录合适的连续字符串数量
	pre := 0           // 记录前面数字相加之和
	m := map[int]int{} // map记录前几个数字之和为K出现相同和的次数为V
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 如果前面数字之和加上这个数字正好等于K（存在一个数字加上nums[i]结果为K
		// 说明找到了
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

// 给定两个数组，编写一个函数来计算它们的交集。
// 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出: [4,9]
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
// 我们可以不考虑输出结果的顺序。

// 如果 nums1 元素个数大于 nums2，则交换数组元素。
// 对于 nums1 的每个元素，添加到 HashMap m 中，如果元素已经存在则增加对应的计数。
// 初始化 k = 0，记录当前交集元素个数。
// 遍历数组 nums2：
// 检查元素在 m 是否存在，若存在且计数为正：
// 将元素拷贝到 nums1[k]，且 k++。
// 减少 m 中对应元素的计数。
// 返回 nums1 前 k 个元素。
func intersect(nums1 []int, nums2 []int) []int {
	// 如果 nums1 元素个数大于 nums2，则交换数组元素。
	if len(nums1) > len(nums2) {
		intersect(nums2, nums1)
	}
	m := make(map[int]int)
	// 对于 nums1 的每个元素，添加到 HashMap m 中，如果元素已经存在则增加对应的计数。
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] += 1
	}
	var k int // 记录当前交集元素个数
	for j := 0; j < len(nums2); j++ {
		// 检查元素在 m 是否存在，若存在且计数为正
		if v, ok := m[nums2[j]]; ok {
			if v > 0 {
				nums1[k] = nums2[j]
				k++
				m[nums2[j]] = v - 1
			}
		}
	}
	return nums1[0:k]
}

//  x 的 n 次幂函数。
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 0 {
		tmp := myPow(x, n/2)
		return tmp * tmp
	} else {
		tmp := myPow(x, n/2)
		return x * tmp * tmp
	}
}

// 折半
func myPow1(x float64, n int) float64 {
	res := pow(x, n)
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := pow(x, n/2) // res 平方就是当前结果
	if n%2 == 0 {
		res = res * res
	} else {
		res = res * res * x
	}
	return res
}

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0 //慢, 快
	for ; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		// nums[j] == val, 跳过
	}
	return i
}

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，
// 并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 给定 nums = [0,1,2,2,3,0,4,2], val = 2,
// 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
// 注意这五个元素可为任意顺序。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0 //慢, 快,  慢指针只有在快指针访问到非val元素, 才往前走
	for ; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		// 快指针遇到val元素, 跳过一次循环
	}
	return i
}

// 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
// 初始化 A 和 B 的元素数量分别为 m 和 n。
// 输入:
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3
// 输出: [1,2,2,3,5,6]
func merge(A []int, m int, B []int, n int) {
	sorted := make([]int, 0, m+n)
	a := 0
	b := 0
	for a < len(A) {
		if a >= m {
			sorted = append(sorted, B[b:]...)
			break
		}
		if b >= n {
			sorted = append(sorted, A[a:m]...)
			break
		}
		if A[a] < B[b] {
			sorted = append(sorted, A[a])
			a++
		} else {
			sorted = append(sorted, B[b])
			b++
		}
	}
	for i := 0; i < len(sorted); i++ {
		A[i] = sorted[i]
	}
}

// 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
// 输入:
// "abccccdd"
// 输出:
// 7
func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]] += 1
	}
	var length int

	for k, v := range m {
		if v%2 == 0 {
			length += v
			m[k] = 0
		}
		if v%2 == 1 {
			length += (v / 2) * 2
			m[k] = m[k] - (v/2)*2
		}
	}
	for _, v := range m {
		if v >= 1 {
			length += 1
			break
		}
	}
	return length
}
