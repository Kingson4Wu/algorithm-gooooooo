package hot100

/*
*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:

输入: x = 1, y = 4

输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)

	↑   ↑

上面的箭头指出了对应二进制位不同的位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func hammingDistance(x int, y int) int {
	z := x ^ y
	//判断有多少个1
	count := 0
	for z > 0 {
		count++
		z = z & (z - 1)
	}
	return count
}

/**

n & (n-1) 每次消除一个1.直到n＝0，就能算出多少个1
或
sum += z & 1;
z = z>>1;
*/
