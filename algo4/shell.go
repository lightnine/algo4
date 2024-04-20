package algo4

// ShellSort 希尔排序
// 希尔排序是插入排序的一种改进版本，它与插入排序的不同之处在于，它会优先比较距离较远的元素。
func ShellSort(items SortInterface) {
	n := items.Len()
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for ; h >= 1; h = h / 3 {
		for i := h; i < n; i++ {
			for j := i; j >= h && items.Less(j, j-h); j -= h {
				items.Swap(j, j-h)
			}
		}
	}
}
