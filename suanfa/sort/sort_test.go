package sort

import "testing"

func TestHeapSort(t *testing.T) {
	RunHeapSort()
}

func TestRunHeapSort2(t *testing.T) {
	RunHeapSort2()
}

func TestRunFastSort(t *testing.T) {
	RunFastSort()
}

func TestRunFastSort2(t *testing.T) {
	FastSort2(arr)
}

func TestRunMergeSort(t *testing.T) {
	RunMergeSort()
}

func TestRunMergeSort2(t *testing.T) {
	RunMergeSort2()
}

func TestRunRadixSort(t *testing.T) {
	RunRadixSort()
}
