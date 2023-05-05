package lib

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Less(x, y int) bool {
	if x <= y {
		return true
	}
	return false
}
