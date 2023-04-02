package hot100

/**
简单个毛，完全是数学技巧题。。。。

遍历异或运算最后的结果

*/
/*
*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

示例 1 ：

输入：nums = [2,2,1]
输出：1
示例 2 ：

输入：nums = [4,1,2,1,2]
输出：4
示例 3 ：

输入：nums = [1]
输出：1
*/
func singleNumber(nums []int) int {

	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

/**

如何才能做到线性时间复杂度和常数空间复杂度呢？

答案是使用位运算。对于这道题，可使用异或运算 ⊕。异或运算有以下三个性质。

任何数和 0 做异或运算，结果仍然是原来的数，即 a⊕0=a。
任何数和其自身做异或运算，结果是 0，即 a⊕a=0。
异或运算满足交换律和结合律，即 a⊕b⊕a = b⊕a⊕a = b⊕(a⊕a)= b⊕0 =b



func singleNumber(nums []int) int {
    single := 0
    for _, num := range nums {
        single ^= num
    }
    return single
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/single-number/solutions/242211/zhi-chu-xian-yi-ci-de-shu-zi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
