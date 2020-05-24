/*
[34, -50, 42, 14, -5, 86] --> 42+14-5+86 = 137
Resulting indices are (2, 5)
*/

package max_sub_array_sum

import "testing"

func Test_maxSubArraySum(t *testing.T) {
	type args struct {
		L []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{"original", args{[]int{34, -50, 42, 14, -5, 86}}, 137, 2, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := maxSubArraySum(tt.args.L)
			if got != tt.want {
				t.Errorf("maxSubArraySum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("maxSubArraySum() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("maxSubArraySum() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
