package hot100

import (
	"fmt"
	"testing"
)

/**
方法一：左右乘积列表
我们不必将所有数字的乘积除以给定索引处的数字得到相应的答案，而是利用索引左侧所有数字的乘积和右侧所有数字的乘积（即前缀与后缀）相乘得到答案。

对于给定索引 i，我们将使用它左边所有数字的乘积乘以右边所有数字的乘积。

方法二：空间复杂度 O(1)的方法
先把 answer 作为方法一的 L 数组。
这种方法的唯一变化就是我们没有构造 R 数组。而是用一个遍历来跟踪右边元素的乘积。并更新数组 answer[i]=answer[i]∗R。然后 R 更新为 R=R∗nums[i]，其中变量 R 表示的就是索引右侧数字的乘积。

作者：力扣官方题解
链接：https://leetcode.cn/problems/product-of-array-except-self/solutions/272369/chu-zi-shen-yi-wai-shu-zu-de-cheng-ji-by-leetcode-/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/**
自己做的，动态规划，但是

执行出错

runtime: out of memory: cannot allocate 4194304-byte block (469499904 in use)
fatal error: out of memory
runtime.throw({0x4c9a16?, 0x31?})
panic.go, line 992
runtime.(*mcache).allocLarge(0xc0000021a0?, 0x61a80, 0x1)
mcache.go, line 215
runtime.mallocgc(0x61a80, 0x4b4c40, 0x1)
malloc.go, line 1096
runtime.makeslice(0x4b4c40?, 0xc350?, 0xc350?)
slice.go, line 103
main.productExceptSelf({0xc000252000, 0xc350, 0x20000?})
solution.go, line 9
main.__helper__(...)
solution.go, line 41
main.main()
solution.go, line 69
runtime.main()
proc.go, line 250
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1571 +0x1 fp=0xc00005cfe8 sp=0xc00005cfe0 pc=0x45d841


重新思考后修改， 其实可以使用一维数组

*/

/*
*

给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
*/
/*func productExceptSelf(nums []int) []int {

	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		//初始化
		dp[i][i] = nums[i]
	}
	for length := 2; length <= len(nums)-2; length++ {
		for i := 0; i < len(nums)-1; i++ {
			j := i + length - 1
			if j > len(nums)-1 {
				break
			}
			dp[i][j] = dp[i][j-1] * dp[j][j]
		}
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = dp[1][1] * dp[2][len(nums)-1]
			continue
		}
		if i == len(nums)-1 {
			result[i] = dp[0][len(nums)-3] * dp[len(nums)-2][len(nums)-2]
			continue
		}
		result[i] = dp[0][i-1] * dp[i+1][len(nums)-1]
	}
	return result
}*/

/*
*
看过题解重做

时间
32 ms
击败
9.84%
内存
7.5 MB
击败
61.86%
*/
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	/** 算出i左边的乘积 */
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}
	/** 和右边相乘 */
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right *= nums[i]
	}

	return ans
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3}))
	fmt.Println(productExceptSelf([]int{1, 2}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
