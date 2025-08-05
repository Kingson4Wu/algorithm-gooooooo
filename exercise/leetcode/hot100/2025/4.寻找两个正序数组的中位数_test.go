package _025

/*
*
核心思路：二分查找切分法
1. 切分两个数组
设数组 A（长度 m）和数组 B（长度 n），假设 m <= n（如果不是，交换 A 和 B）。
我们要在 A 数组上做二分查找，找到一个位置 i，把 A 分成两部分：
A_left: A[0 ~ i-1]
A_right: A[i ~ m-1]

同时在 B 数组上找到一个位置 j，使得：
j = (m + n + 1) / 2 - i （保证左边总元素个数等于右边或多1个）

2. 保证左边最大 <= 右边最小
我们要找到一个 i，使得：
A_left 的最大值 <= B_right 的最小值
B_left 的最大值 <= A_right 的最小值
即：
A[i-1] <= B[j]
B[j-1] <= A[i]
只要找到满足这个条件的 i 和 j，就找到了“中位数切分点”。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)

	if m == 0 && n == 0 {
		return 0
	}

	if m == 0 || n == 0 {
		nums := nums1
		if m == 0 {
			nums = nums2
		}
		if len(nums)%2 == 0 {
			return float64(nums[len(nums)/2-1]+nums[len(nums)]) / 2
		}
		return float64(nums[len(nums)/2-1])
	}

	low, high := 0, m-1

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for low <= high {
		i := (low + high) / 2
		j := (m+n)/2 - 1 - i //  %2 == 1
		if m+n%2 == 0 {
			j = (m+n)/2 - 2 - i
		}
		if j+1 < n && nums1[i] > nums2[j+1] {
			high = i - 1
			continue
		}
		if i+1 < m && nums2[j] > nums1[i+1] {
			low = i + 1
			continue
		}
		low = i
		break
	}
	i := low
	j := (m+n)/2 - 1 - i //  %2 == 1
	if m+n%2 == 0 {
		j = (m+n)/2 - 2 - i
	}
	if (m+n)%2 == 1 {
		return float64(max(nums1[i], nums2[j]))
	}
	//num1 := float64(max(nums1[i], nums2[j]))

	// 边界条件想到头疼
	// 越处理越复杂
	return 0
}

// 思想不难。写起来麻烦
/**
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证nums1是较短的那个
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := (m + n + 1)/2 - i

		var maxLeftA = math.Inf(-1)
		if i != 0 {
			maxLeftA = float64(nums1[i-1])
		}

		var minRightA = math.Inf(1)
		if i != m {
			minRightA = float64(nums1[i])
		}

		var maxLeftB = math.Inf(-1)
		if j != 0 {
			maxLeftB = float64(nums2[j-1])
		}

		var minRightB = math.Inf(1)
		if j != n {
			minRightB = float64(nums2[j])
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			// 找到合适的切分
			if (m+n)%2 == 1 {
				return math.Max(maxLeftA, maxLeftB)
			} else {
				return (math.Max(maxLeftA, maxLeftB) + math.Min(minRightA, minRightB)) / 2.0
			}
		} else if maxLeftA > minRightB {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return 0.0 // 理论上不会到这里
}
*/
