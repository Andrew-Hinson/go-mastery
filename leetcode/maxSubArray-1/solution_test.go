package maxSubArray_1

import "testing"

func TestFindMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "first test",
			args: args{nums: []int{1, 12, -5, -6, 50, 3}, k: 4},
			want: 12.75000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
