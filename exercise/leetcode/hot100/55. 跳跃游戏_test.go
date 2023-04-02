package hot100

import (
	"fmt"
	"testing"
)

/**
自己做的
官方答案使用贪心算法效率更高！！

时间
480 ms
击败
6.22%
内存
6.8 MB
击败
10.55%

*/

/**

public class Solution {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int rightmost = 0;
        for (int i = 0; i < n; ++i) {
            if (i <= rightmost) {
                rightmost = Math.max(rightmost, i + nums[i]);
                if (rightmost >= n - 1) {
                    return true;
                }
            }
        }
        return false;
    }
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/jump-game/solutions/203549/tiao-yue-you-xi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func canJump(nums []int) bool {

	if len(nums) <= 1 {
		return true
	}

	dp := make([]bool, len(nums))

	dp[0] = true

	for i := 0; i < len(nums); i++ {
		if !dp[i] {
			return false
		}
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if !dp[i+j] {
				dp[i+j] = true
			}
		}
	}
	return true
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
