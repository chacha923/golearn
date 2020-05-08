package sort

// 1. 取最大堆的根结点元素。
// 2. 取集合最末尾的元素，放到根结点，调用maxHeapify进行调整。重复步骤1.

// 首先可以看到堆建好之后堆中第0个数据是堆中最小的数据。取出这个数据再执行下堆的删除操作。
// 这样堆中第0个数据又是堆中最小的数据，重复上述步骤直至堆中只有一个数据时就直接取出这个数据。
// 由于堆也是用数组模拟的，故堆化数组后，第一次将A[0]与A[n - 1]交换，再对A[0…n-2]重新恢复堆。
// 第二次将A[0]与A[n – 2]交换，再对A[0…n - 3]重新恢复堆，重复这样的操作直到A[0]与A[1]交换。
// 由于每次都是将最小的数据并入到后面的有序区间，故操作完成后整个数组就有序了。有点类似于直接选择排序。

// 注意使用最小堆排序后是递减数组，要得到递增数组，可以使用最大堆。
func heapSortV2() {
	length := len(heap)
	// 第一次构建堆, 只调整非叶子节点
	for i := length/2 - 1; i >= 0; i-- {
		// sink(i)
		down(i, length)
	}

	for i := length - 1; i > 0; i-- {
		Swap(heap, 0, i)
		// sink(i)
		down(0, i)
	}
}

// 返回父亲索引
func parent(root int) int {
	return (root - 1) / 2
}

// 左孩子索引
func left(root int) int {
	return root*2 + 1
}

// 右孩子索引
func right(root int) int {
	return root*2 + 2
}

/* 上浮第 k 个元素，以维护最大堆性质 */
func swim(k int) {
	for {
		if !(k > 0 && less(parent(k), k)) {
			break
		}

		Swap(heap, parent(k), k)
		k = parent(k)
	}
}

/* 下沉第 k 个元素，以维护最大堆性质 */
func sink(k int) {
	for {
		if k >= len(heap) || k < 0 {
			break
		}
		var older = k // 假设父亲最大
		// 有左孩子, 且左孩子比older大
		if left(k) < len(heap) && less(older, left(k)) {
			older = left(k)
		}
		// 有右孩子, 且右孩子比older大
		if right(k) < len(heap) && less(older, right(k)) {
			older = right(k)
		}
		// 结点 k 比俩孩子都大，就不必下沉了
		if older == k {
			break
		}
		// 交换父亲和最大孩子
		Swap(heap, k, older)
		k = older
	}
}
