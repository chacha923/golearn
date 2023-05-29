package structure

import (
	"fmt"
	"testing"
)

func TestMemo(t *testing.T) {
	memo := NewMemo()
	memo.Put(1, 2, 3)
	fmt.Println(memo.Get(1, 2))
}
