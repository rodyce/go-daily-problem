/*
PROBLEM:
Given a list of numbers and a number k, return whether any two numbers from
the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

APPROACH:
Given a list L and a constant k, find if there is a pair (p, q) such that
p + q = k.
*/

package main

import (
	"testing"
)

var k1, k2 int = 17, 18
var L []int = []int{10, 15, 3, 7}

type args struct {
	L []int
	k int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	// TODO: Add test cases.
	{"mainK1", args{L, k1}, true},
	{"mainK2", args{L, k2}, true},
	{"double", args{L, 2 * L[1]}, false},
	{"empty", args{[]int{}, k1}, false},
	{"justOne", args{[]int{k1}, k1}, false},
}

func Test_isTherePairNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTherePairNaive(tt.args.L, tt.args.k); got != tt.want {
				t.Errorf("isTherePairNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTherePairAlmost(t *testing.T) {
	// Note: This function fails the "double" test.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTherePairAlmost(tt.args.L, tt.args.k); got != tt.want {
				t.Errorf("isTherePairAlmost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTherePair(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTherePair(tt.args.L, tt.args.k); got != tt.want {
				t.Errorf("isTherePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
