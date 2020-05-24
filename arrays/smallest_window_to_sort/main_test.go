/*
Given an array of integers that are out of order, determine the bounds of
the smallest window that must be sorted in order for the whole array to be
sorted. For example [3, 7, 5, 6, 9] -> (1,3)
*/

package smallest_window_to_sort

import "testing"

type args struct {
	L []int
}

var tests = []struct {
	name  string
	args  args
	want  int
	want1 int
}{
	{"original", args{[]int{3, 7, 5, 6, 9}}, 1, 3},
	{"empty", args{[]int{}}, 0, 0},
	{"justOne", args{[]int{6}}, 0, 0},
	{"twoSorted", args{[]int{6, 7}}, 1, 0},
	{"twoNoSort", args{[]int{7, 6}}, 0, 1},
	{"extended", args{[]int{2, 3, 7, 5, 1, 6, 9}}, 0, 5},
	{"upDown", args{[]int{2, 3, 7, 5, 1, 6, 9, 4}}, 0, 7},
}

func Test_smallestWindowToSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := smallestWindowToSort(tt.args.L)
			if got != tt.want {
				t.Errorf("smallestWindowToSort() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("smallestWindowToSort() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_smallestWindowToSortSol(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := smallestWindowToSortSol(tt.args.L)
			if got != tt.want {
				t.Errorf("smallestWindowToSortSol() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("smallestWindowToSortSol() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
