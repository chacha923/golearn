

## 快排模版
这是一个前序遍历
```
func sort(nums []int, lo, hi int) {
	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, p+1, hi)
}
```

## 归并排序模版
这是一个后续遍历
```
func sort(nums []int, lo int, hi int) {
	    mid := lo + (hi-lo) / 2
		sort(nums, lo, mid)
		sort(nums, mid+1, hi)

		merge(nums, lo, mid, hi)
}
```