package hot100

import (
	"fmt"
	"testing"
)

func subsetsBinaryIter(nums []int) (ans [][]int) {

	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func TestSubsetsBinaryIter(t *testing.T) {
	fmt.Println(subsetsBinaryIter([]int{1, 2, 3}))
}
