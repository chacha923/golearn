package search

import (
	"fmt"
	"testing"
)

var (
	array = []int{2, 4, 5, 8, 12, 16, 23, 37, 70, 100}
)

func TestBinSearch(t *testing.T) {
	fmt.Println(binarySearch(array, 100))
}

func TestBinarySearch2(t *testing.T) {
	fmt.Println(binarySearch2(array, 0, len(array)-1, 100))
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(8))
}
