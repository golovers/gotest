package rotate_test

import (
	"testing"

	"github.com/golovers/gotest/rotate"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int32
		rotate int32
		ouput  []int32
	}{
		{
			name:   "empty slice",
			input:  []int32{},
			rotate: 10000,
			ouput:  []int32{},
		},
		{
			name:   "rotate an even number of times compare to len of the input",
			input:  []int32{1, 2, 3, 4, 5},
			rotate: 10000,
			ouput:  []int32{1, 2, 3, 4, 5},
		},
		{
			name:   "rotate an odd number of times compare to len of the input",
			input:  []int32{1, 2, 3, 4, 5},
			rotate: 10003,
			ouput:  []int32{4, 5, 1, 2, 3},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := rotate.RotateLeft(test.input, test.rotate)
			if !sliceEqual(got, test.ouput) {
				t.Errorf("got %v, want %v", got, test.ouput)
			}
		})
	}
}

func sliceEqual(first []int32, second []int32) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func BenchmarkRotate(b *testing.B) {
	arr := [1000000]int32{1, 2, 3, 4, 5}
	slice := arr[:]
	for i := 0; i < b.N; i++ {
		rotate.RotateLeft(slice, int32(289734292))
	}
}

func BenchmarkRotateOrg(b *testing.B) {
	arr := [100000000]int32{1, 2, 3, 4, 5}
	slice := arr[:]
	for i := 0; i < b.N; i++ {
		rotate.RotateLeftOrg(slice, int32(289734292))
	}
}
