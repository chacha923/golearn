package kmp

import (
	"testing"
	"fmt"
)

func TestRunKMP(t *testing.T) {
	content := []byte("why every programming language use the hello world as the first test???")
	sub := []byte("hello world")
	fmt.Println(RunKMP(content, 0, len(content)-1, sub))
}
