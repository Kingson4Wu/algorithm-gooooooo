package tencent

import (
	"fmt"
	"testing"
)

/**
比较简单

漏了本来就有序的测试用例！！

时间
16 ms
击败
8.83%
内存
4.2 MB
击败
39.21%
*/

/**
还有一种方法叫快慢指针

int removeDuplicates(int[] nums) {
    int n = nums.length;
    if (n == 0) return 0;
    int slow = 0, fast = 1;
    while (fast < n) {
        if (nums[fast] != nums[slow]) {
            slow++;
            // 维护 nums[0..slow] 无重复
            nums[slow] = nums[fast];
        }
    fast++;
    }
    // ⻓度为索引 + 1
    return slow + 1;
}
*/

/*
*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。

由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。

将最终结果插入 nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

判题标准:

系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

如果所有断言都通过，那么您的题解将被 通过。

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	length := 1
	index := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			length++
			if index >= 0 {
				nums[index] = nums[i+1]
				index++
			}

			continue
		}
		if index == -1 {
			index = i + 1
		}
	}

	return length
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 2, 3}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
