package problem01

import "testing"

func TestTargetPairExists(t *testing.T) {
	type args struct {
		arr1      []int
		arr2      []int
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil_input_return_false", args{nil, nil, 3}, false},
		{"empty_input_return_false", args{[]int{}, []int{}, 2}, false},
		{"find", args{[]int{10, 40, 5, 280}, []int{234, 5, 2, 148, 23}, 42}, true},
		{"not_find", args{[]int{1, 2, 3}, []int{10, 19, 23}, -2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TargetPairExists(tt.args.arr1, tt.args.arr2, tt.args.targetSum); got != tt.want {
				t.Errorf("TargetPairExists() = %v, want %v", got, tt.want)
			}
		})
	}
}