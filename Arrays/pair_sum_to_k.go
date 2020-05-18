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

func isTherePairNaive(L []int, k int) bool {
	for i, p := range L {
		for _, q := range L[i+1:] {
			if p+q == k {
				return true
			}
		}
	}
	return false
}

func isTherePairAlmost(L []int, k int) bool {
	// Use a set to
	type void struct{}
	var member void

	someSet := make(map[int]void)
	for _, p := range L {
		someSet[p] = member
	}

	for _, p := range L {
		q := k - p
		if _, ok := someSet[q]; ok {
			//do something here
			if ok {
				return true
			}
		}
	}
	return false
}

func isTherePair(L []int, k int) bool {
	someMap := make(map[int]int)
	for _, p := range L {
		if _, ok := someMap[p]; ok {
			someMap[p]++
		} else {
			someMap[p] = 1
		}
	}

	for _, p := range L {
		q := k - p
		if count, ok := someMap[q]; ok {
			if ok && (p != q || count > 1) {
				return true
			}
		}
	}
	return false
}

func main() {
}
