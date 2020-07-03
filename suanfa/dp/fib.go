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
