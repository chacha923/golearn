package lang

import "fmt"

func ReturnBeforeDefer() string {
	return "return above defer"
	defer fmt.Println("defer keyword")
	return "return below defer"
}
