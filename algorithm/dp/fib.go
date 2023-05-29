package dp

// 菲波那切数列
func Fib(n int) int {
	fib := make([]int, n+1)
	fib[1] = 1
	fib[2] = 1
	for i := 3; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func FibByRecurse(n int) int {
	// base case
	if n == 1 || n == 2 {
		return 1
	}
	// 备忘录
	var memo = make([]int, n+1)

	var dp func(memo []int, i int) int
	dp = func(memo []int, i int) int {
		if val := memo[i]; val > 0 {
			return val
		}
		memo[n] = dp(memo, n-1) + dp(memo, n-2)
		return memo[n]
	}

	return dp(memo, n)
}
