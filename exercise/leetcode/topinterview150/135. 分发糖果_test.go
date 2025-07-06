package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

示例 1：

输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例 2：

输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。

	第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
*/

/*
*
看题解后做
方法二：常数空间遍历
1. 糖果总是尽量少给，且从 1 开始累计，每次要么比相邻的同学多给一个，要么重新置为 1（记录升序的长度）
2. 降序期间，逐步数量累记加1，当 降序长度=升序长度 时，额外再加1，因为原来的最高位加入这个递减区间，并且比原来+1
1,2,3
1,2,3,1
1,2,3,2,1 		|| (1,2,3, 1,2)
1,2,4,3,2,1 	|| (1,2,3, 1,2,3) ->  (1,2,4, 1,2,3) ->  ->  (1,2,  1,2,3,4)
1,2,5,4,3,2,1 	|| (1,2,  1,2,3,4,5)

看了题解还是写不出来，以下代码很乱都是错的，变量控制有点乱
*/
func candy(ratings []int) int {

	total := 1
	pre := 1
	ins, dec := 1, 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			if dec > 0 {
				ins = 1
				dec = 0
				pre = 1
			}
			ins++
			pre++
			total += pre
		} else {

			if dec == 0 {
				if pre > 1 {
					pre = 0
				}
			}

			if ratings[i] == ratings[i-1] {
				ins = 0
			} else {
				dec++
				pre++
			}

			if dec == ins {
				ins = 0
				pre++
			}
			total += pre
		}
	}

	return total
}

/**
解答错误
27 / 49 个通过的测试用例

官方题解
输入
ratings =
[1,3,2,2,1]

添加到测试用例
输出
11
预期结果
7
*/

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
	fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
	fmt.Println(candy([]int{1, 2, 5, 4, 3, 2, 1, 2, 3}))
	fmt.Println(candy([]int{5, 3, 2, 1, 3, 4, 5, 2, 1, 4}))
}

/**
7
5
4
23
24
*/

/**
func candy(ratings []int) int {
    n := len(ratings)
    ans, inc, dec, pre := 1, 1, 0, 1
    for i := 1; i < n; i++ {
        if ratings[i] >= ratings[i-1] {
            dec = 0
            if ratings[i] == ratings[i-1] {
                pre = 1
            } else {
                pre++
            }
            ans += pre
            inc = pre
        } else {
            dec++
            if dec == inc {
                dec++
            }
            ans += dec
            pre = 1
        }
    }
    return ans
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/candy/solutions/533150/fen-fa-tang-guo-by-leetcode-solution-f01p/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
