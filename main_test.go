package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{[]int{1}, 2, [][]int{}},
		{[]int{}, 3, [][]int{}},
	}
	for _, c := range cases {
		if got := CombinationSum2(c.in, c.target); !reflect.DeepEqual(got, c.want) {
			t.Errorf("CombinationSum2(%v,%d) = %v, want %v", c.in, c.target, got, c.want)
		}
	}
}
