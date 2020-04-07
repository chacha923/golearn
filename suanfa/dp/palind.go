package dp

// 最长回文子串
// 动态规划方法, P(i,j) = true (子串Si-Sj是回文) | false 其他情况

// 中心展开法
// 事实上，只需使用恒定的空间，我们就可以在 O(n^2)O(n2) 的时间内解决这个问题。
// 我们观察到回文中心的两侧互为镜像。因此，回文可以从它的中心展开，并且只有 2n - 12n−1 个这样的中心。
// 你可能会问，为什么会是 2n - 12n−1 个，而不是 nn 个中心？原因在于所含字母数为偶数的回文的中心可以处于两字母之间（例如 \textrm{“abba”}“abba” 的中心在两个 \textrm{‘b’}‘b’ 之间）。

func longestPalindrome(s string) string {
	var start int
	var length int                //子串长度
	for i := 0; i < len(s); i++ { //奇数长度子串
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > length {
			start = left + 1
			length = right - left - 1
		}
	}
	for i := 0; i < len(s); i++ { //偶数长度
		left := i
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > length {
			start = left + 1
			length = right - left - 1
		}
	}
	return s[start : start+length]
}
