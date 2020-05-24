/*
[34, -50, 42, 14, -5, 86] --> 42+14-5+86 = 137
Resulting indices are (2, 5)
*/

package max_sub_array_sum

func maxSubArraySum(L []int) (int, int, int) {
	running_max, max_so_far := 0, 0
	l, r := 0, 0
	for i, n := range L {
		if n > running_max+n {
			running_max = n
			l, r = i, i
		} else {
			running_max = running_max + n
		}

		if running_max > max_so_far {
			max_so_far = running_max
			r = i
		}
	}

	return max_so_far, l, r
}
