
# 排序
1. [稳定]冒泡排序：
   + 时间：最坏-O(n^2)、平均-O(n^2)、最好-O(n)； 空间：O(1)
   + 相邻元素两两比较交换，双重循环；
   + 单次遍历没有交换的，提前结束全部遍历
   + length<=1 的数组不需要处理
2. [不稳定]快速排序：
   + 时间：最坏-O(n^2)、平均-O(nlog2n)、最好-O(nlog2n)； 空间：O(nlog2n)
   + 快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。
   + （高低下标的值交替被替换，最后得到中间下标，把参照值设置回去）for low < high, 每次 v[low]= v[high], v[high]=v[low], 最后v[low]=temp
   + TopK，递归步骤进行剪枝：
     - mid>k, 处理左边部分
     - mid<k, 处理右边部分
     - 结果无序排序的时候可以用
3. [不稳定]堆排序：
   + 时间：最坏-O(nlog2n)、平均-O(nlog2n)、最好-O(nlog2n)； 空间：O(1)
   + 堆的结构可以分为大顶堆和小顶堆，是一个完全二叉树
   + 底层数组结构
   + 大顶堆：`arr(i)>arr(2*i+1) && arr(i)>arr(2*i+2)`；小顶堆：`arr(i)<arr(2*i+1) && arr(i)<arr(2*i+2)`
   + 根找左右：左：`(i+1)*2 - 1`，右：`(i+1)*2` ;
   + 左找根右：根：`(i+1)/2 - 1`，右：`i+1`
   + 右找根左：根：`i/2 - 1`, 左：`i-1`

   + 初始化堆
   + 从最后一个结点的父结点（arr.length / 2 - 1）开始 逐次调整位置，开始构建最大堆
       - 1 若父结点小于左结点，父结点与左结点互换，继续调整
       - 2 若父结点小于右结点，父结点与右结点互换（注意是经过1），继续调整

   + 1 初始化堆，比如小堆：加入堆最后一个结点后，与根节点对比，比根节点小则交换，并继续，否则结束完成（自底向上）
   + 2 调整堆（拿出堆中第一个元素后，即根最后一个结点交换后），将根节点和左右两个元素比较，与较小的那个交换，并继续，若都比自己大，或者已经是叶子结点，则结束完成（自上往下）
   + 如何实现
        1. 实现一个调整堆的方法，传参为nums []int, parent, len int， 即指定根结点parent，从上到下调整，根结点和子结点（先从左右两结点选择满足条件的即可）交换，直到没有子结点
        2. 初始化堆，从最下的根结点到最上的根结点，即从length/2 到0，每个根结点使用上述调整堆的方法
        3. 对堆进行排序，即把根结点（index=0），和最后一个元素交换，然后对[0,length-1]的堆进行调整, 遍历length-1到0， 交换和重新调整堆。
    

+ 其他： https://kingson4wu.github.io/2022/02/28/20220228-%E7%BB%8F%E5%85%B8%E6%8E%92%E5%BA%8F%E7%AE%97%E6%B3%95/

----

# 大文件

## 如何从 100 亿 URL 中找出相同的 URL
1. 分而治之，进行哈希取余，分开存文件；
2. 对每个子文件进行 HashSet 统计；
3. 如果分完还有很大的文件，那就文件再继续分，一直递归到文件符合设置的最大大小即可

+ https://github.com/Kingson4Wu/algorithm-gooooooo/blob/c5b9fc92e401878f7949771979816b3c78077885/Algorithm/interview/100billion/%E5%A6%82%E4%BD%95%E4%BB%8E%20100%20%E4%BA%BF%20URL%20%E4%B8%AD%E6%89%BE%E5%87%BA%E7%9B%B8%E5%90%8C%E7%9A%84%20URL.go

----

# 常见题目系列

## 丑数

+ Algorithm/interview/ugly

### 263. 丑数（简单）
+ 输入一个数字n，请你判断n是否为「丑数」。所谓「丑数」，就是只包含质因数2、3和5的正整数。
```go
var factors = []int{2, 3, 5}

func isUgly(n int) bool {
if n <= 0 {
return false
}
for _, f := range factors {
for n%f == 0 {
n /= f
}
}
return n == 1
} 
```
### 264. 丑数 II（中等）
+ 给你一个整数 n ，请你找出并返回第 n 个 丑数 。
+ 动态规划: 定义数组 dp，其中 dp[i]表示第 i个丑数，第 n个丑数即为 dp[n]。
    - 定义三个指针 p2,p3,p5，表示下一个丑数是当前指针指向的丑数乘以对应的质因数。初始时，三个指针的值都是 1
    - 当 2≤i≤n 时，令 dp[i]=min(dp[p2]×2,dp[p3]×3,dp[p5]×5)，然后分别比较 dp[i]和 dp[p2]×2,dp[p3]×3,dp[p5]×5 是否相等，如果相等则将对应的指针加 1
```go
func nthUglyNumber(n int) int {
    dp := make([]int, n+1)
    dp[1] = 1
    p2, p3, p5 := 1, 1, 1
    for i := 2; i <= n; i++ {
        x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
        dp[i] = min(min(x2, x3), x5)
        if dp[i] == x2 {
            p2++
        }
        if dp[i] == x3 {
            p3++
        }
        if dp[i] == x5 {
            p5++
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

### 1201. 丑数 III（中等）
+ 给你四个整数：n 、a 、b 、c ，请你设计一个算法来找出第 n 个丑数。
丑数是可以被 a 或 b 或 c 整除的 正整数 。

### 313. 超级丑数（中等）
+ 这道题和「264. 丑数 II」相似，区别在于，第 264 题规定丑数是只包含质因数 2、3 和 5 的正整数，这道题规定超级丑数是只包含数组 primes 中的质因数的正整数。

```go
func nthSuperUglyNumber(n int, primes []int) int {
    dp := make([]int, n+1)
    m := len(primes)
    pointers := make([]int, m)
    nums := make([]int, m)
    for i := range nums {
        nums[i] = 1
    }
    for i := 1; i <= n; i++ {
        minNum := math.MaxInt64
        for j := range pointers {
            minNum = min(minNum, nums[j])
        }
        dp[i] = minNum
        for j := range nums {
            if nums[j] == minNum {
                pointers[j]++
                nums[j] = dp[pointers[j]] * primes[j]
            }
        }
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

```

