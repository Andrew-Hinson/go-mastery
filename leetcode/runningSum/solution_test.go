package runningSum

import (
	"reflect"
	"testing"
)

func Test_runningSumOf1dArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first test",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "second test",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRunningSumOf1dArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSumOf1dArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
