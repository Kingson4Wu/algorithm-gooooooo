package hot100

import (
	"fmt"
	"testing"
)

/**

推荐使用Floyd 判圈算法, 和双向链表判断入环点一样,根据这道题的特点,其实就是一个有环的链表

个人想法，跟
exercise/leetcode/hot100/448. 找到所有数组中消失的数字_test.go
一样
通过下标负数寻找
最后再恢复数组

可惜不让修改数组，题目虽然通过，实际是错的

方法一：二分查找
方法二：二进制
方法三：快慢指针

func findDuplicate(nums []int) int {
    slow, fast := 0, 0
    for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] { }
    slow = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/find-the-duplicate-number/solutions/261119/xun-zhao-zhong-fu-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

时间
76 ms
击败
91.70%
内存
8.5 MB
击败
31.62%

*/
/*
*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：

输入：nums = [1,3,4,2,2]
输出：2
示例 2：

输入：nums = [3,1,3,4,2]
输出：3

nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

进阶：

如何证明 nums 中至少存在一个重复的数字?
你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
*/
/*func findDuplicate(nums []int) int {

	result := 0
	for i := 0; i < len(nums); i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		if nums[index] < 0 {
			result = index
			break
		} else {
			nums[index] = -nums[index]
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
		}
	}

	return result
}*/

/*
*
时间
92 ms
击败
25.91%
内存
8.2 MB
击败
50.12%
*/
func findDuplicate(nums []int) int {

	/**
	必须初始值一样！！！！
	*/
	slow, fast := 0, 0
	slow, fast = nums[slow], nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p := 0
	for slow != p {
		slow = nums[slow]
		p = nums[p]
	}

	return p
}

func TestFindDuplicate(t *testing.T) {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
