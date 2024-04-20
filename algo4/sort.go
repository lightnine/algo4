package algo4

type SortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func IsSorted(items SortInterface) bool {
	for i := 1; i < items.Len(); i++ {
		if items.Less(i, i-1) {
			return false
		}
	}
	return true
}
