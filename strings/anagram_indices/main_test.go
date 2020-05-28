package anagram_indices

import (
	"reflect"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple", args{"arroz", "zorra"}, true},
		{"almost", args{"arros", "zorra"}, false},
		{"empty", args{"", ""}, true},
		{"tricky", args{"acc", "accc"}, false},
		{"emptyTricky", args{"", "a"}, false},
		{"emptyTricky", args{"a", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	word string
	s    string
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"simple", args{"ab", "abxus"}, []int{0}},
	{"original", args{"ab", "abxaba"}, []int{0, 3, 4}},
	{"tricky", args{"ab", "ab"}, []int{0}},
	{"illust", args{"ab", "aba"}, []int{0}},
}

func Test_findAnagramIndicesNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagramIndicesNaive(tt.args.word, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagramIndicesNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAnagramIndices(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagramIndices(tt.args.word, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagramIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
