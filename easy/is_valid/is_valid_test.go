package is_valid

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	type test struct {
		name  string
		input string
		want  bool
	}
	tests := []test{
		{
			name:  "is true",
			input: "(({[]}))",
			want:  true,
		},
		{
			name:  "is false",
			input: "(({)})",
			want:  false,
		},
		{
			name:  "not same size",
			input: "(({)}",
			want:  false,
		},
		{
			name:  "all matching",
			input: "()[]{}",
			want:  true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.input)
			if got != tt.want {
				t.Errorf("%d: TwoSum(%s), got=%v: got!=want: want=%v", i, tt.input, got, tt.want)
			}
		})
	}

}
