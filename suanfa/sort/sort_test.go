package sort

import "testing"

func TestHeapSort(t *testing.T) {
	RunHeapSort()
}

func TestRunFastSort(t *testing.T) {
	RunFastSort()
}

func TestRunMergeSort(t *testing.T) {
	RunMergeSort()
}

func TestRunRadixSort(t *testing.T) {
	RunRadixSort()
}
