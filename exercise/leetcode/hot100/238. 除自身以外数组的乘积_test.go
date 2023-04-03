package hot100

import (
	"fmt"
	"testing"
)

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

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	/*for i := 0; i < len(nums); i++ {
		result[i]=
	}*/

	return result
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3}))
	fmt.Println(productExceptSelf([]int{1, 2}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
