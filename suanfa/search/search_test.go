package search

import (
	"testing"
	"fmt"
)

var (
	array = []int{2, 4, 5, 8, 12, 16, 23, 37, 70, 100}
)

func TestBinSearch(t *testing.T) {
	fmt.Println(binarySearch(array, 100))
}
