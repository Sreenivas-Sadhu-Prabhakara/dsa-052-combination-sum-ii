package main

import (
	"fmt"
	"sort"
)

// CombinationSum2 returns all unique combinations summing to target where
// each array element may be used at most once. The array may contain
// duplicates; equal values are skipped at the same depth so no combination
// repeats.
func CombinationSum2(candidates []int, target int) [][]int {
	a := append([]int{}, candidates...)
	sort.Ints(a)
	out := [][]int{}
	cur := []int{}
	var walk func(start, remaining int)
	walk = func(start, remaining int) {
		if remaining == 0 {
			out = append(out, append([]int{}, cur...))
			return
		}
		for i := start; i < len(a) && a[i] <= remaining; i++ {
			if i > start && a[i] == a[i-1] {
				continue
			}
			cur = append(cur, a[i])
			walk(i+1, remaining-a[i])
			cur = cur[:len(cur)-1]
		}
	}
	walk(0, target)
	return out
}

func main() {
	fmt.Println(CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
