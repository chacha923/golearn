package dp

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	ans := math.MaxInt64
	for _, c := range coins {
		if amount-c < 0 {
			continue
		}
		sub := coinChange(coins, amount-c)
		if sub == -1 {
			continue
		}
		ans = min(ans, sub+1)
	}
	if ans == math.MaxInt64 {
		return -1
	} else {
		return ans
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
