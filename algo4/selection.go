package algo4

// SelectionSort 实现选择排序算法
func SelectionSort(items SortInterface) {
	n := items.Len()
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if items.Less(j, min) {
				min = j
			}
		}
		items.Swap(i, min)
	}
}
