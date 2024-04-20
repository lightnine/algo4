package algo4

// InsertionSort implements insertion sort algorithm.
// 主要思想就是每次在内循环中将小的元素往前移动
func InsertionSort(items SortInterface) {
	n := items.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && items.Less(j, j-1); j-- {
			items.Swap(j, j-1)
		}
	}
}
