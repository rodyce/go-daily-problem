/*
Given an array of integers that are out of order, determine the bounds of
the smallest window that must be sorted in order for the whole array to be
sorted. For example [3, 7, 5, 6, 9] -> (1,3)
*/

package smallest_window_to_sort

// This solution attempt is wrong.
func smallestWindowToSort(L []int) (int, int) {
	var l, r int = 0, 0
	var max_in_range int

	for i, n := range L {
		if i == 0 {
			continue
		}
		if n >= L[i-1] {
			if l == r {
				l = i
				r = i
				max_in_range = n
			} else if n > max_in_range {
				return l, r
			} else {
				r++
			}
		} else {
			r = i
			for l > 0 && L[l-1] > n {
				l--
			}
		}
	}

	return l, r
}

func smallestWindowToSortSol(L []int) (int, int) {
	size := len(L)
	if size <= 0 {
		return 0, 0
	}

	l, r := 0, 0
	max_l, min_r := L[l], L[r]

	for i := 0; i < len(L); i++ {
		if L[i] > max_l {
			max_l = L[i]
		}
		if L[i] < max_l {
			r = i
		}
	}

	for i := size - 1; i >= 0; i-- {
		if L[i] < min_r {
			min_r = L[i]
		}
		if L[i] > min_r {
			l = i
		}
	}

	return l, r
}
