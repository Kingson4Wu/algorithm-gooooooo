package dp

import (
	"fmt"
	"testing"
)

/*
*
沿街有一排连续的房屋。每间房屋内都藏有一定的现金。现在有一位小偷计划从这些房屋中窃取现金。

由于相邻的房屋装有相互连通的防盗系统，所以小偷 不会窃取相邻的房屋 。

小偷的 窃取能力 定义为他在窃取过程中能从单间房屋中窃取的 最大金额 。

给你一个整数数组 nums 表示每间房屋存放的现金金额。形式上，从左起第 i 间房屋中放有 nums[i] 美元。

另给你一个整数 k ，表示窃贼将会窃取的 最少 房屋数。小偷总能窃取至少 k 间房屋。
返回小偷的 最小 窃取能力。

示例 1：

输入：nums = [2,3,5,9], k = 2
输出：5
解释：
小偷窃取至少 2 间房屋，共有 3 种方式：
- 窃取下标 0 和 2 处的房屋，窃取能力为 max(nums[0], nums[2]) = 5 。
- 窃取下标 0 和 3 处的房屋，窃取能力为 max(nums[0], nums[3]) = 9 。
- 窃取下标 1 和 3 处的房屋，窃取能力为 max(nums[1], nums[3]) = 9 。
因此，返回 min(5, 9, 9) = 5 。
示例 2：

输入：nums = [2,7,9,3,1], k = 2
输出：2
解释：共有 7 种窃取方式。窃取能力最小的情况所对应的方式是窃取下标 0 和 4 处的房屋。返回 max(nums[0], nums[4]) = 2 。

提示：

1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= k <= (nums.length + 1)/2

想不出来，只想到暴力穷举回溯
只能看题解

这题考的是：在不能偷相邻房子的条件下，选择至少 k 个房子，求所有方案中「最大金额最小」的那种偷法的最大单间金额（即窃取能力）的最小值。
二分答案 + 贪心判断

最小的窃取能力下界是：min(nums)，因为至少要偷一个房子。
上界是：max(nums)，因为最多可能偷到最大值。
我们对这个区间进行 二分查找，找到第一个满足条件的最小值

lower= min(nums), upper=max(nums)
令 middle=(lower+upper)/2，如果 f(middle)≥k，那么 upper=middle−1；否则 lower=middle+1。
当 lower≤upper 时，继续执行步骤 2；否则返回 lower 为结果。
*/
func minCapability(nums []int, k int) int {

	lower, upper := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > upper {
			upper = nums[i]
		} else if nums[i] < lower {
			lower = nums[i]
		}
	}

	for lower <= upper {
		middle := (lower + upper) / 2
		count, visited := 0, false
		for i := 0; i < len(nums); i++ {
			if !visited && nums[i] <= middle {
				count++
				if count >= k {
					break
				}
				visited = true
			} else {
				visited = false
			}
		}
		if count >= k {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}
	return lower
}

/**
执行用时分布
15
ms
击败
85.94%
复杂度分析
消耗内存分布
10.38
MB
击败
64.06%
*/

func TestMinCapability(t *testing.T) {
	//fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{2, 7, 9, 3, 1}, 2))
}

/**
好，我们就用题目里的例子：

---

## 🎯 示例：

```text
nums = [2, 3, 5, 9]
k = 2
```

目标：找一个最小的偷窃能力 `x`，在「只能偷金额 ≤ x 的房子」的前提下，**能偷至少 2 间且不能偷相邻的房子**。

---

## 🔍 一步步模拟 + 二分流程图讲解

我们用「二分 + 贪心」来找答案，设：

* `left = 0`（偷窃能力下限）
* `right = 9`（偷窃能力上限 = max(nums)）

我们开始二分：

---

### 第一步：mid = (0 + 9) / 2 = 4

💬 意思是：“我尝试只偷 ≤ 4 元的房子，可以偷够 2 间吗？”

执行贪心：

* i = 0：nums\[0] = 2，≤4 ✅，偷 → count = 1
* i = 1：相邻不能偷，跳过
* i = 2：nums\[2] = 5 ❌ 太贵，跳过
* i = 3：nums\[3] = 9 ❌ 太贵，跳过

✔ 最终只偷了 1 间 < k，不够 → ❌ 不行

所以我们把区间往右移：

```
left = 5, right = 9
```

---

### 第二步：mid = (5 + 9) / 2 = 7

💬 意思是：“我试试能力是 7 的时候行不行”

贪心过程：

* i = 0：nums\[0] = 2 ✅，偷 → count = 1
* i = 1：不能偷，跳
* i = 2：nums\[2] = 5 ✅，偷 → count = 2
* i = 3：不能偷相邻，跳

✔ 偷了 2 间 ✅ 成功！

说明能力 7 可行，那我们继续往小试：

```
left = 5, right = 7
```

---

### 第三步：mid = (5 + 7) / 2 = 6

💬 意思是：“试试能力为 6”

贪心过程同上：

* i = 0：2 ✅ 偷 → count = 1
* i = 1：跳
* i = 2：5 ✅ 偷 → count = 2
* i = 3：跳

✔ 成功 ✅

继续缩小右边：

```
left = 5, right = 6
```

---

### 第四步：mid = (5 + 6) / 2 = 5

💬 意思是：“试试最小可能的答案 5 行不行？”

执行：

* i = 0：2 ✅ 偷 → count = 1
* i = 1：跳
* i = 2：5 ✅ 偷 → count = 2
* i = 3：跳

✔ 成功 ✅

继续缩：

```
left = 5, right = 5
```

---

### 🎉 结束条件：left == right == 5

说明 **最小可行的偷窃能力就是 `5`**

---

## 🧾 最终答案：5

---

## ✅ 总结这个过程你只要记得两个点：

1. **“我试试只偷不超过 x 元的房子，能不能选出 k 间、且不相邻？”**（贪心判断）
2. **不断缩小这个 `x`，找最小可行值**（二分查找）


*/
