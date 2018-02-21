package goalgorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	implementations := []struct {
		name string
		sort func([]int)
	}{
		{"SelectionSort", SelectionSort},
		{"SelectionSortTemp", SelectionSortTemp},
	}
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"Mixed", []int{1, 3, 4, 5, 2, 9, 8, 0}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Already sorted", []int{0, 1, 2, 3, 4, 5, 8, 9}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Almost sorted", []int{0, 1, 2, 3, 4, 5, 9, 8}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Reversed", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Large set of 50 values",
			[]int{
				703, 741, 755, 275, 283, 800, 120, 902, 744, 848, 473, 529, 277, 449, 172, 141, 773, 746, 308, 103,
				263, 787, 11, 259, 253, 211, 569, 613, 110, 990, 664, 588, 434, 600, 930, 145, 188, 293, 896, 719,
				534, 721, 23, 476, 671, 763, 254, 123, 838, 208,
			},
			[]int{
				11, 23, 103, 110, 120, 123, 141, 145, 172, 188, 208, 211, 253, 254, 259, 263, 275, 277, 283, 293,
				308, 434, 449, 473, 476, 529, 534, 569, 588, 600, 613, 664, 671, 703, 719, 721, 741, 744, 746, 755,
				763, 773, 787, 800, 838, 848, 896, 902, 930, 990,
			},
		},
	}
	for _, impl := range implementations {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				impl.sort(tt.list)
				got := tt.list
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", impl.name, tt.list, got, tt.want)
				}
			})
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	tests := [][]int{
		[]int{4, 1, 9, 6, 2, 5, 3, 7, 8, 0},
		[]int{
			703, 741, 755, 275, 283, 800, 120, 902, 744, 848, 473, 529, 277, 449, 172, 141, 773, 746, 308, 103,
			263, 787, 11, 259, 253, 211, 569, 613, 110, 990, 664, 588, 434, 600, 930, 145, 188, 293, 896, 719,
			534, 721, 23, 476, 671, 763, 254, 123, 838, 208,
		},
	}

	implementations := []struct {
		name string
		sort func([]int)
	}{
		{"SelectionSort", SelectionSort},
		{"SelectionSortTemp", SelectionSortTemp},
	}
	for _, tt := range tests {
		for _, impl := range implementations {
			b.Run(impl.name+fmt.Sprintf("%d", len(tt)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					c := make([]int, len(tt))
					copy(c, tt)
					impl.sort(c)
				}
			})
		}
	}
}
