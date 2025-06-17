package maxConsecutiveOnesIII

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first test",
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
