package kmp

//构造next数组
func GetNextValueArray(sub []byte) (next []int) {
	var (
		length        int = len(sub)
		middle        int
		compare_left  int
		compare_right int
		match_count   int
	)

	next = make([]int, length)
	next[0] = 0
	next[1] = 0

	for i := 2; i < length; i++ {
		middle = i / 2
		match_count = 0
		if i%2 == 0 {
			for j := 0; j < middle; j++ {
				compare_left = 0
				compare_right = i - 1 - j
				for compare_left <= j {
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left ++
					compare_right++
				}
				if compare_left == j+1 {
					match_count++
				}
			}
			match_count++
		} else {
			for j := 0; j <= middle; j++ {
				compare_left = 0
				compare_right = i - 1 - j
				for compare_left <= j {
					if sub[compare_left] != sub[compare_right] {
						break
					}
					compare_left++
					compare_right++
				}
				if compare_left == j+1 {
					match_count++
				}
			}
			next[i] = match_count
		}
	}
	return next
}

//next数组优化
func ReviseNextValueArray(next []int) []int {
	var length int = len(next)
	for i := 2; i < length; i++ {
		if next[i] == next[next[i]] {
			next[i] = next[next[i]]
		}
	}

	return next
}

func RunKMP(content []byte, start_index int, end_index int, sub []byte) (index int) {

	var (
		next       []int = ReviseNextValueArray(GetNextValueArray(sub))
		sub_index  int   = 0
		sub_length int   = len(sub)
	)

	for i := start_index; i <= end_index; i++ {
		if content[i] == sub[sub_index] {
			match_index := i
			for j := sub_index; j <= sub_length; j++ {
				if j == sub_length {
					return match_index - sub_index
				}
				if i >= end_index || content[i] != sub[j] {
					sub_index = next[j]
					break
				}
				i++
			}
		}
	}
	return -1
}
