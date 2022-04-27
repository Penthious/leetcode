package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T){
	type test struct {
		input []int
		target int
		want []int
	}
	tests := []test{
		{
			input: []int{2,7,11,15},
			target:9,
			want: []int{0,1},
		},
		{
		input: []int{-1,-2,-3,-4,-5},
		target:-8,
		want: []int{2,4},
		},
	}

	for i, tt := range tests {
		got := TwoSum(tt.input, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d: TwoSum(%v, %d), got=%v: got!=want: want=%v", i, tt.input, tt.target, got, tt.want)
		}
	}

}
