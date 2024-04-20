package algo4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		items SortInterface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				items: StringSlice{"3", "2", "1", "45", "32", "21"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.items)
			assert.True(t, IsSorted(tt.args.items))
		})
	}
}
