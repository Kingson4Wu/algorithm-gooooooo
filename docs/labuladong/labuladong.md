labuladong


----
# 精华总结

----
# 套路模版

----

# TODO标记


---

# 第一章 核心套路篇

## 1.1 学习算法和刷题的框架思维

+ 数据结构的存储方式只有两种:数组(顺序存储)和链表(链式存储)
+ 「图」的两种表示方法，邻接表就是链表，邻接矩阵就是二维数组
+ 「树」，用数组实现就是「堆」，因为「堆」是一个完全二叉树
+ 线性就是 for/while 迭代为代表，非线性就是递归为代表
+ N 叉树的遍历又可以扩展为图的遍历，因为图就是好几 N 叉棵树的结合 体。你说图是可能出现环的?这个很好办，用个布尔数组 visited 做标记就行


+ 先刷二叉树，先刷二叉树，先刷二叉树!

+ 试着从框架上看问题，而不要纠结于细节问题 ！！

## 1.2 动态规划解题套路框架

+ 动态规划问题的一般形式就是求最值
+ 求解动态规划的核心问题是穷举

+ 首先，动态规划的穷举有点特别，因为这类问题存在「重叠子问题」，如果 暴力穷举的话效率会极其低下，所以需要「备忘录」或者「DP table」来优 化穷举过程，避免不必要的计算。
+ 而且，动态规划问题一定会具备「最优子结构」，才能通过子问题的最值得 到原问题的最值。

+ 穷举所有可行解其实并不是一件容易的事，只有列出正确的「状态转移方 程」才能正确地穷举。

+ 重叠子问题、最优子结构、状态转移方程就是动态规划三要素。 ！！！

+ 辅助你思考状态转移方程:
明确「状态」 -> 定义 dp 数组/函数的含义 -> 明确「选择」-> 明确 base case。 ！！！

+ 三点：状态、选择、dp数组的定义

+ 套路框架：
```
# 初始化base case
dp[0][0][...] = base case
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态1 in 状态1的所有取值：
        for ...
            dp[状态1][状态2][...] = 求最值（选择1，选择2，...）

```

+ 递归的做法会存在大量重复计算 （重叠子问题-> 通过备忘录解决）

+ 自顶向下（递归）、自底向上（迭代）！

+ 带备忘录的递归解法的效率已经和迭代的动态规划解法一样了。实际
上，这种解法和迭代的动态规划已经差不多了，只不过这种方法叫做「自顶
向下」，动态规划叫做「自底向上」。
+ 啥叫「自顶向下」?注意我们刚才画的递归树(或者说图)，是从上向下延 伸，都是从一个规模较大的原问题比如说 f(20) ，向下逐渐分解规模，直 到 f(1) 和 f(2) 触底，然后逐层返回答案，这就叫「自顶向下」。
+ 啥叫「自底向上」?反过来，我们直接从最底下，最简单，问题规模最小的 f(1) 和 f(2) 开始往上推，直到推到我们想要的答案 f(20) ，这就是动 态规划的思路，这也是为什么动态规划一般都脱离了递归，而是由循环迭代
完成计算。

+ 千万不要看不起暴力解，动态规划问题最困难的就是写出状态转移方程，即 这个暴力解。优化方法无非是用备忘录或者 DP table，再无奥妙可言。 ！！！

+ 状态压缩！优化空间存储

### 例子
+ 斐波那契数列、凑零钱问题

## 1.3 回溯算法解题套路框架 （todo）

+ 解决一个回溯问题，实际上就是一个决 策树的遍历过程。你只需要思考 3 个问题:
    - 1、路径:也就是已经做出的选择。 
    - 2、选择列表:也就是你当前可以做的选择。 
    - 3、结束条件:也就是到达决策树底层，无法再做选择的条件。

+ 回溯算法的框架:
```
result = []
def backtrack(路径, 选择列表): 
    if 满足结束条件:
        result.add(路径) 
        return
    for 选择 in 选择列表: 
        做选择
        backtrack(路径, 选择列表) 
        撤销选择  
```  
+ 其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用 之后「撤销选择」！！！！

### 例子
+ 全排列问题、N皇后问题

## 1.4 BFS算法套路框架（todo）

+ BFS（Broad First Search，广度优先搜索）和 DFS（Depth First Search，深度优先搜索）。其中DFS可以算是回溯算法。

+ BFS 相对 DFS 的最主要的区别是:BFS 找到的路径一定是最短的，但代价 就是空间复杂度比 DFS 大很多

### 例子
+ 二叉树的最小高度、解开密码锁的最少次数

## 1.5 双指针技巧套路框架

+ 快慢指针（链表）、左右指针（数组）

### 1.5.1 快、慢指针的常用算法
1. 判断链表中是否有环
    - 跑得快的指针遇到null，说明不含环；两个指针最终相遇，说明含有环。
2. 已知链表含有环，返回这个环的起始位置
    - 数学问题，要画图
    - 第一次相遇时，把其中一个指针重新指向head起点，再同速前进，第二次相遇之处就是环的起点。
3. 寻找无环单链表的中点
    - 快指针到达尽头时，慢指针就处于链表的中间位置。
4. 寻找单链表的倒数第k个元素
    - 快指针先走k步，然后同速前进，快指针到达末尾时，慢指针就是倒数第k个节点

### 1.5.2 左右指针的常用算法
1. 二分搜索
2. 两数之和
    - 数组有序
3. 反转数组
    - 交换元素
4. 滑动窗口算法

## 1.6 二分查找解题套路框架

### 二分查找框架
```
int binarySearch(int[] nums, int target) {
    int left = 0, right = ...;
    while(...) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            ...
        } else if (nums[mid] < target) {
            left = ...
        } else if (nums[mid] > target) {
            right = ... }
        }
    return ...; 
}
```

## 1.7 滑动窗口解题套路框架 （todo）

```
int left = 0, right = 0;
while (right < s.size()) { // 增大窗口
    window.add(s[right]);
    right++;
    while (window needs shrink) { // 缩小窗口
        window.remove(s[left]);
        left++; 
    }
}
```

现在开始套模板，只需要思考以下四个问题:
1、当移动 right 扩大窗口，即加入字符时，应该更新哪些数据? 
2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口? 
3、当移动 left 缩小窗口，即移出字符时，应该更新哪些数据? 
4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新?
如果一个字符进入窗口，应该增加 window 计数器;如果一个字符将移出 窗口的时候，应该减少 window 计数器;当 valid 满足 need 时应该收缩 窗口;应该在收缩窗口的时候更新最终结果。

### 1.7.1 最小覆盖子串
+ 在 S (source)中找到包含 T (target)中全部字母的一个子串，且这 个子串一定是所有可能子串中最短的。
+ 左右指针
+ 实现细节很繁琐，TODO

### 1.7.2 字符串排列
+ 给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

+ 这种题目，是明显的滑动窗口算法，相当给你一个 S 和一个 T ，请问你 S 中是否存在一个子串，包含 T 中所有字符且不包含其他字符

### 1.7.3 找所有字母异位词
+  这是 LeetCode 第 438 题，Find All Anagrams in a String

### 1.7.4 最⻓无重复子串
+ 这是 LeetCode 第 3 题，Longest Substring Without Repeating Characters


# 第二章 动态规划系列

+ 解题基本步骤
    1. 找到“状态”和“选择”
    2. 明确dp数组/函数的定义
    3. 寻找状态之间的关系

## 2.1 最长递增子序列    
+ (Longest Increasing Subsequence，简写 LIS)
+ 比较容易想到的是动态规划解法，时间复杂度 O(N^2)，我们借 这个问题来由浅入深讲解如何写动态规划。比较难想到的是利用二分查找， 时间复杂度是 O(NlogN) 

+ 注意「子序列」和「子串」这两个名词的区别，子串一定是连续的，而子序列不一定是连续的。！！！

```
public int lengthOfLIS(int[] nums) { 
        int[] dp = new int[nums.length]; 
        // dp 数组全都初始化为 1 
        Arrays.fill(dp, 1);
        for (int i = 0; i < nums.length; i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j])
                    dp[i] = Math.max(dp[i], dp[j] + 1);
            } 
        }
        int res = 0;
        for (int i = 0; i < dp.length; i++) {
            res = Math.max(res, dp[i]);
        }
        return res; 
    }
```

+ 二分搜索解法：其实 最⻓递增子序列和一种叫做 patience game 的纸牌游戏有关，甚至有一种排 序方法就叫做 patience sorting(耐心排序)。

## 2.2 二维递增子序列：信封嵌套问题

+ 俄罗斯套娃信封
+ https://labuladong.github.io/zgnb/3/15/17/

+ 先对宽度w进行升序排列，如果遇到w相同的情况，则按照高度H进行降序排列（因为宽度一样不可能套娃，比如5 和5 不能同时进 去），然后把所有的H 作为一个数组，在这个数组上进行 计算（LIS）
+ 因为两个W相同的信封不能相互包含，W相同时将H逆序排序，则这些逆序H中至多只会有一个被选入递增子序列，保证了最终的信封序列不会出现W相同的情况

```
// envelopes = [[w, h], [w, h]...]
public int maxEnvelopes(int[][] envelopes) {
    int n = envelopes.length;
    // 按宽度升序排列，如果宽度一样，则按高度降序排列
    Arrays.sort(envelopes, new Comparator<int[]>() 
    {
        public int compare(int[] a, int[] b) {
            return a[0] == b[0] ? 
                b[1] - a[1] : a[0] - b[0];
        }
    });
    // 对高度数组寻找 LIS
    int[] height = new int[n];
    for (int i = 0; i < n; i++)
        height[i] = envelopes[i][1];

    return lengthOfLIS(height);
}
```

+ 正确排序后，就转为了一维数组的 LIS（最长递增子序列）问题

## 2.3 最大子数组问题
+ https://leetcode.cn/problems/maximum-subarray/
+ https://labuladong.github.io/zgnb/3/15/18/

+ 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

+ p[i] = Math.max(nums[i], nums[i] + dp[i - 1]);

```
int maxSubArray(int[] nums) {
    int n = nums.length;
    if (n == 0) return 0;
    int[] dp = new int[n];
    // base case
    // 第一个元素前面没有子数组
    dp[0] = nums[0];
    // 状态转移方程
    for (int i = 1; i < n; i++) {
        dp[i] = Math.max(nums[i], nums[i] + dp[i - 1]);
    }
    // 得到 nums 的最大子数组
    int res = Integer.MIN_VALUE;
    for (int i = 0; i < n; i++) {
        res = Math.max(res, dp[i]);
    }
    return res;
}
```

+ 状态压缩 todo


## 2.4 最优子结构及dp遍历方向
+ 正向、反向、斜向

## 2.5 最长公共子序列
+ 最⻓公共子序列(Longest Common Subsequence，简称 LCS)是一道非常经 典的面试题目，因为它的解法是典型的二维动态规划，大部分比较困难的字 符串问题都和这个问题一个套路，比如说编辑距离

+ 输入: str1 = "abcde", str2 = "ace" 输出: 3
解释: 最⻓公共子序列是 "ace"，它的⻓度是 3


+ 动态规划算法做的就是穷举 + 剪枝，它俩天生一对儿。所以可以说只要涉及子序列问题，十有八九都需要 动态规划来解决 ！！！

+ dp[i][j] 的含义是:对于 s1[1..i] 和 s2[1..j] ， 它们的LCS⻓度是 dp[i][j] 

```
def longestCommonSubsequence(str1, str2) -> int: 
m, n = len(str1), len(str2)
# 构建 DP table 和 base case
dp = [[0] * (n + 1) for _ in range(m + 1)]
# 进行状态转移
for i in range(1, m + 1):
        for j in range(1, n + 1):
            if str1[i - 1] == str2[j - 1]:
# 找到一个 lcs 中的字符
                dp[i][j] = 1 + dp[i-1][j-1]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    return dp[-1][-1]
```

## 2.6 编辑距离

+ https://labuladong.github.io/zgnb/3/15/19/

+ https://leetcode.cn/problems/edit-distance/

给你两个单词word1 和word2， 请返回将word1转换成word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

+ 为什么说它实用呢，因为前几天我就在日常生活中用到了这个算法。之前有一篇公众号文章由于疏忽，写错位了一段内容，我决定修改这部分内容让逻辑通顺。但是公众号文章最多只能修改 20 个字，且只支持增、删、替换操作（跟编辑距离问题一模一样），于是我就用算法求出了一个最优方案，只用了 16 步就完成了修改。

再比如高大上一点的应用，DNA 序列是由 A,G,C,T 组成的序列，可以类比成字符串。编辑距离可以衡量两个 DNA 序列的相似度，编辑距离越小，说明这两段 DNA 越相似，说不定这俩 DNA 的主人是远古近亲啥的。

+ 编辑距离问题就是给我们两个字符串 s1 和 s2，只能用三种操作，让我们把 s1 变成 s2，求最少的操作数。需要明确的是，不管是把 s1 变成 s2 还是反过来，结果都是一样的


```
def minDistance(s1, s2) -> int:
    # 定义：dp(i, j) 返回 s1[0..i] 和 s2[0..j] 的最小编辑距离
    def dp(i, j):
        # base case
        if i == -1: return j + 1
        if j == -1: return i + 1
        
        if s1[i] == s2[j]:
            return dp(i - 1, j - 1)  # 啥都不做
        else:
            return min(
                dp(i, j - 1) + 1,    # 插入
                dp(i - 1, j) + 1,    # 删除
                dp(i - 1, j - 1) + 1 # 替换
            )
    
    # i，j 初始化指向最后一个索引
    return dp(len(s1) - 1, len(s2) - 1)

```

+ 动态规划优化
对于重叠子问题呢，前文动态规划详解详细介绍过，优化方法无非是备忘录或者 DP table !!!
+ DP table 是自底向上求解，递归解法是自顶向下求解 !

```
int minDistance(String s1, String s2) {
    int m = s1.length(), n = s2.length();
    int[][] dp = new int[m + 1][n + 1];
    // base case 
    for (int i = 1; i <= m; i++)
        dp[i][0] = i;
    for (int j = 1; j <= n; j++)
        dp[0][j] = j;
    // 自底向上求解
    for (int i = 1; i <= m; i++)
        for (int j = 1; j <= n; j++)
            if (s1.charAt(i-1) == s2.charAt(j-1))
                dp[i][j] = dp[i - 1][j - 1];
            else               
                dp[i][j] = min(
                    dp[i - 1][j] + 1,
                    dp[i][j - 1] + 1,
                    dp[i-1][j-1] + 1
                );
    // 储存着整个 s1 和 s2 的最小编辑距离
    return dp[m][n];
}

int min(int a, int b, int c) {
    return Math.min(a, Math.min(b, c));
}

```

+ 状态压缩： 既然每个 dp[i][j] 只和它附近的三个状态有关，空间复杂度是可以压缩成 O(min(M, N)) 的（M，N 是两个字符串的长度）。不难，但是可解释性大大降低，读者可以自己尝试优化一下。

## 2.7 子序列问题解题模版：最长回文子序列
+ https://blog.csdn.net/fdl123456/article/details/103359919

### 两种思路
#### 1、第一种思路模板是一个一维的 dp 数组:
```
int n = array.length;
int[] dp = new int[n];
for (int i = 1; i < n; i++) {
    for (int j = 0; j < i; j++) {
        dp[i] = 最值(dp[i], dp[j] + ...)
    }
}
```

+ 「最⻓递增子序列」
+ 在子数组 array[0..i] 中，我们要求的子序列(最⻓递增子序列)的⻓度 是 dp[i] 

#### 2、第二种思路模板是一个二维的 dp 数组:

```
int n = arr.length;
int[][] dp = new dp[n][n];
for (int i = 0; i < n; i++) {
    for (int j = 0; j < n; j++) {
        if (arr[i] == arr[j])
            dp[i][j] = dp[i][j] + ...
        else
            dp[i][j] = 最值(...)
    } 
}
```
+ 涉及两个字符串/数组的子序列，比如 前文讲的「最⻓公共子序列」。本思路中 dp 数组含义又分为「只涉及一个 字符串」和「涉及两个字符串」两种情况
    - 涉及两个字符串/数组时(比如最⻓公共子序列)，dp 数组的含义如 下:在子数组 arr1[0..i] 和子数组 arr2[0..j] 中，我们要求的子序列(最⻓ 公共子序列)⻓度为 dp[i][j] 。
    - 只涉及一个字符串/数组时(比如本文要讲的最⻓回文子序列)，dp 数 组的含义如下:在子数组 array[i..j] 中，我们要求的子序列(最⻓回文子序列)的⻓度 为 dp[i][j] 。

### 最长回文子序列

+ 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。



示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-palindromic-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


+ 如果它俩相等，那么它俩加上s[i+1..j-1]中的最长回文子序列就是s[i..j]的最长回文子序列;
如果它俩不相等，说明它俩不可能同时出现在s[i..j]的最长回文子序列中，那么把它俩分别加入s[i+1..j-1]中，看看哪个子串产生的回文子序列更长即可

```
if (s[i] == s[j])    
    // 它俩一定在最长回文子序列中    
    dp[i][j] = dp[i + 1][j - 1] + 2;
else    
    // s[i+1..j] 和 s[i..j-1] 谁的回文子序列更长？    
    dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);
```

+ 为了保证每次计算dp[i][j]，左、下、左下三个方向的位置已经被计算出来，只能斜着遍历或者反着遍历

```
int longestPalindromeSubseq(string s) {    
    int n = s.size();    
    // dp 数组全部初始化为 0   
    vector<vector<int>> dp(n, vector<int>(n, 0));    
    // base case   
    for (int i = 0; i < n; i++)        
        dp[i][i] = 1;    
    // 反着遍历保证正确的状态转移    
    for (int i = n - 1; i >= 0; i--) {        
        for (int j = i + 1; j < n; j++) {            
            // 状态转移方程            
            if (s[i] == s[j])                
                dp[i][j] = dp[i + 1][j - 1] + 2;            
            else                
                dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);        
        }    
    }    
    // 整个 s 的最长回文子串长度    
    return dp[0][n - 1];
}
```

## 2.8 状态压缩

+ https://www.cnblogs.com/labuladong/p/13940269.html

+ 动态规划技巧对于算法效率的提升非常可观，一般来说都能把指数级和阶乘级时间复杂度的算法优化成 O(N^2)
+ 但是，动态规划本身也是可以进行阶段性优化的，比如说我们常听说的「状态压缩」技巧，就能够把很多动态规划解法的空间复杂度进一步降低，由 O(N^2) 降低到 O(N)，
+ 能够使用状态压缩技巧的动态规划都是二维 dp 问题，你看它的状态转移方程，如果计算状态 dp[i][j] 需要的都是 dp[i][j] 相邻的状态，那么就可以使用状态压缩技巧，将二维的 dp 数组转化成一维，将空间复杂度从 O(N^2) 降低到 O(N)。

+ 使用状态压缩技巧对二维 dp 数组进行降维打击之后，解法代码的可读性变得非常差了，如果直接看这种解法，任何人都是一脸懵逼的。算法的优化就是这么一个过程，先写出可读性很好的暴力递归算法，然后尝试运用动态规划技巧优化重叠子问题，最后尝试用状态压缩技巧优化空间复杂度。

## 2.9 以最小插入次数构造回文串
+ https://www.cnblogs.com/labuladong/p/13940284.html

+ 给你一个字符串s，每一次操作你都可以在字符串的任意位置插入任意字符。

请你返回让s成为回文串的最少操作次数。

「回文串」是正读和反读都相同的字符串。



示例 1：

输入：s = "zzazz"
输出：0
解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


+ 定义一个二维的 dp 数组，dp[i][j] 的定义如下：对字符串 s[i..j]，最少需要进行 dp[i][j] 次插入才能变成回文串。
+ 我们想求整个 s 的最少插入次数，根据这个定义，也就是想求 dp[0][n-1] 的大小（n 为 s 的长度）。
+ 同时，base case 也很容易想到，当 i == j 时 dp[i][j] = 0，因为当 i == j 时 s[i..j] 就是一个字符，本身就是回文串，所以不需要进行任何插入操作。


+ s[i] != s[j] 时的代码逻辑如下：

```
if (s[i] != s[j]) {
    // 步骤一选择代价较小的
    // 步骤二必然要进行一次插入
    dp[i][j] = min(dp[i + 1][j], dp[i][j - 1]) + 1;
}
```

## 2.10 动态规划之正则表达式

+ https://www.cnblogs.com/labuladong/p/12320270.html

## 2.11 四键盘问题

+ 四键键盘问题很有意思，而且可以明显感受到:对 dp 数组的不同定义需要 完全不同的逻辑，从而产生完全不同的解法。

+ 第一种思路 这种思路会很容易理解，但是效率并不高，我们直接走流程:对于动态规划
问题，首先要明白有哪些「状态」，有哪些「选择」。
具体到这个问题，对于每次敲击按键，有哪些「选择」是很明显的:4 种， 就是题目中提到的四个按键，分别是 A 、 C-A 、 C-C 、 C-V ( Ctrl 简 写为 C )。
接下来，思考一下对于这个问题有哪些「状态」?或者换句话说，我们需要 知道什么信息，才能将原问题分解为规模更小的子问题?
你看我这样定义三个状态行不行:第一个状态是剩余的按键次数，用 n 表 示;第二个状态是当前屏幕上字符 A 的数量，用 a_num 表示;第三个状态 是剪切板中字符 A 的数量，用 copy 表示。
如此定义「状态」，就可以知道 base case:当剩余次数 n 为 0 时， a_num 就是我们想要的答案。
结合刚才说的 4 种「选择」，我们可以把这几种选择通过状态转移表示出 来:

dp(n - 1, a_num + 1, copy), # A 解释:按下 A 键，屏幕上加一个字符 同时消耗 1 个操作数
dp(n - 1, a_num + copy, copy), # C-V 解释:按下 C-V 粘贴，剪切板中的字符加入屏幕 同时消耗 1 个操作数
dp(n - 2, a_num, a_num) # C-A C-C 解释:全选和复制必然是联合使用的，
剪切板中 A 的数量变为屏幕上 A 的数量
同时消耗 2 个操作数

+ 第二种思路
这种思路稍微有点复杂，但是效率高。继续走流程，「选择」还是那 4 个， 但是这次我们只定义一个「状态」，也就是剩余的敲击次数 n 。
这个算法基于这样一个事实，最优按键序列一定只有两种情况: 要么一直按 A :A,A,...A(当N比较小时)。 要么是这么一个形式:A,A,...C-A,C-C,C-V,C-V,...C-V(当 N 比较大时)。
因为字符数量少(N 比较小)时， C-A C-C C-V 这一套操作的代价相对比较 高，可能不如一个个按 A ;而当 N 比较大时，后期 C-V 的收获肯定很 大。这种情况下整个操作序列大致是:开头连按几个 A ，然后 C-A C-C 组合再接若干 C-V ，然后再 C-A C-C 接着若干 C-V ，循环下去。
换句话说，最后一次按键要么是 A 要么是 C-V 。


## 2.12 高楼扔鸡蛋问题

+ https://blog.csdn.net/Ostkakah/article/details/119004570
```java
int maxOperationCount(int eggSize, int Count)//count代表层数并非代表总高度
{
	int i, res = M;
	if(eggSize == 1) return Count;//只剩最后一个鸡蛋,我们就直接线性遍历,所以最坏的情况为全部遍历一遍。
	if(Count == 0) return 0;//高度为0时，就不需要再向下执行了。
	for(i = 1; i <= Count; i++){//为了保证找到对应的步数，我们必须遍历开始时不同的扔鸡蛋的位置
		res = Min(res,
			Max(
				maxOperationCount(eggSize - 1, i - 1),//碎了，鸡蛋数减一，所以我们只要求楼层在自己的基础上高度减一的层数对应的测试步数。 
				maxOperationCount(eggSize, Count - i)//没碎,楼层数为此层之上,所以我们只要求i到Count的层数对应的测试步数。
				)
			) + 1;//每步要加一;  
	}
	return res;
}
int Max(int a, int b)
{
	return (a >= b) ? a : b;
}
int Min(int a,int b)
{
	return (a >= b) ? b : a;
}
```

+ 为了能够列出所有的情况，我们需要进行一个循环，将当前的楼层数全部循环一遍，找出每层鸡蛋会碎的情况和鸡蛋不会碎的情况，因此选择就分为 鸡蛋会碎和鸡蛋不会碎 ，因为要求最坏情况，所以我们要在每层的最多步数中找最小的值，最终就可以求出答案。

### 双蛋问题
+ https://zhuanlan.zhihu.com/p/260378938 !!!
最常见的问法是 N = 100，K = 2，也是就所谓的双蛋问题。跟只有 1 蛋相比，2 个鸡蛋可谓是富足，这个时候就没必要如履薄冰地线性搜索了，步子可以大一点。比如说：

等间隔丢
第一个鸡蛋可以每隔 10 层丢一次：10、20、30...100，如果碎了第二个鸡蛋再从前面的 9 层线性搜索。比如说第一个蛋在 10 层碎了，那么第二个蛋就在 [1,9] 之间试，也就是：

第一个鸡蛋尝试：10 20 30 40 50 60 70 80 90 100（最多尝试 10 次） 第二个鸡蛋最多尝试 9 次

因此，总的来说，最好的情况是第一个蛋在第 10 层就碎了，总次数是 10 次，最坏的情况是第一个蛋在第 100 层碎，总的次数就是 19 次。最好和最坏的情况相差比较大，这是因为第一个蛋每次都是等间隔丢，所以第二个蛋丢的时候，无论如何最坏都要尝试 9 次。

不等间隔丢
使用等间隔丢的方法，如果间距取得比较大，当第一个蛋碎的时候，第二个蛋要试的次数就比较大；当间距取比较小的时候，当 F 的位置越靠后，第一个蛋要试的次数就越大。

有没有办法让两个蛋丢的次数均衡一下呢，试试刚开始的时候间距取大一些，越往后间距逐渐缩小。即：第一次的间隔 为 n，如果没碎第二次的间隔为 n...，一直到最后一层间隔为 1。使用高斯公式可以知道  ，则  ，向上取整 n 等于 14，也就是：

第一个鸡蛋尝试：14 27 39 50 60 89 77 84 90 95 99 100（最多尝试12次） 第二个鸡蛋最多尝试的区间是 [1,13]，一共是13次

因此，总的来说，最好的情况是 12 次，最坏的情况是 14次。相对于等间距的 10 - 19 次要平均一些了，最坏情况的次数也更少。

###  N 层楼和  K个鸡蛋
+ 给定  N 层楼和  K个鸡蛋，要求找到扔下鸡蛋不碎的最高楼层（临界楼层 F ），那么最少尝试几次一定能找到这个临界楼层？我们可以定义问题如下：给定输入 N 、 K ，输出为最少尝试次数 Y
+ 假设在第 i 楼尝试，会存在两种情况（碎和不碎）： - 如果碎了，需要在低楼层 [1, i-1] 搜索，问题规模缩小为：y1=dp(i-1, k-1)  - 如果没碎，需要在高楼层 [i-1, N] 搜索，问题规模缩小为： y2=dp(N-i, k)

+ 提示：  1,10层 2 个鸡蛋 和  11,20层 2个鸡蛋，两个问题是等价的，都是 y=dp(10, 2) ，问题的关键是楼层数量和鸡蛋个数，而不是楼层编号，很好理解，对吧。
+ 因此，对于在第  i楼的尝试，最坏情况下的尝试次数max(y1,y2)  。而 i 可以在 [1, N] 中选择一个，根据题意，我们要找出这  种选择里最少的尝试次数

+ 检查重叠子问题：
这个问题是存在重叠子问题的，例如前面说的 [1,10] 层 2 个鸡蛋 和 [11,20] 层 2个鸡蛋，两个问题是等价的，问题的关键是楼层数量和鸡蛋个数，而不是楼层编号。假如我们曾经计算过函数值  ，那么下一次遇到 N = 10,K = 2 的问题，就可以直接返回前者的答案。

为此，我们需要使用“备忘录”把前者的答案记忆起来。用程序实现无非是基于数组或者基于散列表，这里使用二维数组

+ base case
当楼层数 N 等于 0 时，显然不需要扔鸡蛋；当鸡蛋数 K 为 1 时，显然只能线性扫描所有楼层：

def dp(K, N):
    if K == 1: return N
    if N == 0: return 0

```
def superEggDrop(K: int, N: int):

    memo = dict()
    def dp(K, N) -> int:
        # base case
        if K == 1: return N
        if N == 0: return 0
        # 避免重复计算
        if (K, N) in memo:
            return memo[(K, N)]

        res = float('INF')
        # 穷举所有可能的选择
        for i in range(1, N + 1):
            res = min(res, 
                      max(
                            dp(K, N - i), 
                            dp(K - 1, i - 1)
                         ) + 1
                  )
        # 记入备忘录
        memo[(K, N)] = res
        return res

    return dp(K, N)
```

+ 用二分搜索优化(替代线性搜索)也只是做了「剪枝」，减小了搜索空间，但本质思路没有变，还是穷举

```
// 把线性搜索改成二分搜索
// for (int m = 1; dp[K][m] < N; m++)
int lo = 1, hi = N;
while (lo < hi) {
    int mid = (lo + hi) / 2;
    if (... < N) {
        lo = ...
    } else {
        hi = ...
    }

    for (int k = 1; k <= K; k++)
        // 状态转移方程
}
```

+ 重新定义状态转移
    - dp 数组的定义，确定当前的鸡蛋个数和最多允许的扔鸡蛋次数，就知道能够确定 F 的最高楼层数


## 2.13 高楼扔鸡蛋（进阶）

+ 二分搜索优化！

## 2.14 戳气球问题
+ 好复杂的样子，先没仔细看
+ https://zhuanlan.zhihu.com/p/144384951

+ 「全排列」问题，我们前文 回溯算法框架套路详解 中有全排列算法的详解和代码，其实只要稍微改一下逻辑即可，伪码思路如下：

```
int res = Integer.MIN_VALUE;
/* 输入一组气球，返回戳破它们获得的最大分数 */
int maxCoins(int[] nums) {
    backtrack(nums, 0); 
    return res;
}
/* 回溯算法的伪码解法 */
void backtrack(int[] nums, int socre) {
    if (nums 为空) {
        res = max(res, score);
        return;
    }
    for (int i = 0; i < nums.length; i++) {
        int point = nums[i-1] * nums[i] * nums[i+1];
        int temp = nums[i];
        // 做选择
        在 nums 中删除元素 nums[i]
        // 递归回溯
        backtrack(nums, score + point);
        // 撤销选择
        将 temp 还原到 nums[i]
    }
}
```
+ 回溯算法就是这么简单粗暴，但是相应的，算法的效率非常低。这个解法等同于全排列，所以时间复杂度是阶乘级别

### 动态规划
+ 在一排气球 points 中，请你戳破气球 0 和气球 n+1 之间的所有气球（不包括 0 和 n+1），使得最终只剩下气球 0 和气球 n+1 两个气球，最多能够得到多少分？
+ dp[i][j] = x 表示，戳破气球 i 和气球 j 之间（开区间，不包括 i 和 j）的所有气球，可以获得的最高分数为 x
+ 题目要求的结果就是 dp[0][n+1] 的值，而 base case 就是 dp[i][j] = 0，其中 0 <= i <= n+1, j <= i+1，因为这种情况下，开区间 (i, j) 中间根本没有气球可以戳
+ dp[i][j] = dp[i][k] + dp[k][j] 
         + points[i]*points[k]*points[j]

```
int maxCoins(int[] nums) {
    int n = nums.length;
    // 添加两侧的虚拟气球
    int[] points = new int[n + 2];
    points[0] = points[n + 1] = 1;
    for (int i = 1; i <= n; i++) {
        points[i] = nums[i - 1];
    }
    // base case 已经都被初始化为 0
    int[][] dp = new int[n + 2][n + 2];
    // 开始状态转移
    // i 应该从下往上
    for (int i = n; i >= 0; i--) {
        // j 应该从左往右
        for (int j = i + 1; j < n + 2; j++) {
            // 最后戳破的气球是哪个？
            for (int k = i + 1; k < j; k++) {
                // 择优做选择
                dp[i][j] = Math.max(
                    dp[i][j], 
                    dp[i][k] + dp[k][j] + points[i]*points[j]*points[k]
                );
            }
        }
    }
    return dp[0][n + 1];
}
```

## 2.15 0-1 背包问题
+ https://www.cnblogs.com/labuladong/p/13927944.html
+ 给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少

+ 第一步要明确两点，「状态」和「选择」。

先说状态，如何才能描述一个问题局面？只要给几个物品和一个背包的容量限制，就形成了一个背包问题呀。所以状态有两个，就是「背包的容量」和「可选择的物品」。

再说选择，也很容易想到啊，对于每件物品，你能选择什么？选择就是「装进背包」或者「不装进背包」嘛


+ 第二步要明确 dp 数组的定义。
    - dp[i][w] 的定义如下：对于前 i 个物品，当前背包的容量为 w，这种情况下可以装的最大价值是 dp[i][w]
    - 如果 dp[3][5] = 6，其含义为：对于给定的一系列物品中，若只对前 3 个物品进行选择，当背包容量为 5 时，最多可以装下的价值为 6
    - 为什么要这么定义？便于状态转移，或者说这就是套路，记下来就行了。建议看一下我们的动态规划系列文章，几种套路都被扒得清清楚楚了。
    - 根据这个定义，我们想求的最终答案就是 dp[N][W]。base case 就是 dp[0][..] = dp[..][0] = 0，因为没有物品或者背包没有空间的时候，能装的最大价值就是 0。

 ```
for i in [1..N]:
    for w in [1..W]:
        dp[i][w] = max(
            dp[i-1][w], //不把物品 i 装进背包
            dp[i-1][w - wt[i-1]] + val[i-1] // 把物品 i 装进背包,
        )
return dp[N][W]
 ```   

 ```
int knapsack(int W, int N, vector<int>& wt, vector<int>& val) {
    // base case 已初始化
    vector<vector<int>> dp(N + 1, vector<int>(W + 1, 0));
    for (int i = 1; i <= N; i++) {
        for (int w = 1; w <= W; w++) {
            if (w - wt[i-1] < 0) {
                // 这种情况下只能选择不装入背包
                dp[i][w] = dp[i - 1][w];
            } else {
                // 装入或者不装入背包，择优
                dp[i][w] = max(dp[i - 1][w - wt[i-1]] + val[i-1], 
                               dp[i - 1][w]);
            }
        }
    }
    
    return dp[N][W];
}

 ```

## 2.16 子集背包问题 
+ https://www.cnblogs.com/labuladong/p/13927956.html
+ 416.分割等和子集
+ 给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少
+ 把问题转化为背包问题：

给一个可装载重量为 sum / 2 的背包和 N 个物品，每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满

+ 第一步要明确两点，「状态」和「选择」。

这个前文 经典动态规划：背包问题 已经详细解释过了，状态就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」。

第二步要明确 dp 数组的定义。

按照背包问题的套路，可以给出如下定义：

dp[i][j] = x 表示，对于前 i 个物品，当前背包的容量为 j 时，若 x 为 true，则说明可以恰好将背包装满，若 x 为 false，则说明不能恰好将背包装满。
比如说，如果 dp[4][9] = true，其含义为：对于容量为 9 的背包，若只是用前 4 个物品，可以有一种方法把背包恰好装满。

（因为sum是总户数的一半，所以恰好将背包装满时，不可能包含所有的物品！！！）

根据这个定义，我们想求的最终答案就是 dp[N][sum/2]，base case 就是 dp[..][0] = true 和 dp[0][..] = false，因为背包没有空间的时候，就相当于装满了，而当没有物品可选择的时候，肯定没办法装满背包。

第三步，根据「选择」，思考状态转移的逻辑。

```C++
bool canPartition(vector<int>& nums) {
    int sum = 0;
    for (int num : nums) sum += num;
    // 和为奇数时，不可能划分成两个和相等的集合
    if (sum % 2 != 0) return false;
    int n = nums.size();
    sum = sum / 2;
    vector<vector<bool>> 
        dp(n + 1, vector<bool>(sum + 1, false));
    // base case
    for (int i = 0; i <= n; i++)
        dp[i][0] = true;
    
    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= sum; j++) {
            if (j - nums[i - 1] < 0) {
               // 背包容量不足，不能装入第 i 个物品
                dp[i][j] = dp[i - 1][j]; 
            } else {
                // 装入或不装入背包
                dp[i][j] = dp[i - 1][j] || dp[i - 1][j-nums[i-1]];
            }
        }
    }
    return dp[n][sum];
}

```


+ 进行状态压缩
注意到 dp[i][j] 都是通过上一行 dp[i-1][..] 转移过来的，之前的数据都不会再使用了。

所以，我们可以进行状态压缩，将二维 dp 数组压缩为一维，节约空间复杂度


+ 唯一需要注意的是 j 应该从后往前反向遍历，因为每个物品（或者说数字）只能用一次，以免之前的结果影响其他的结果。
状态压缩部分没看懂，为什么要反向遍历？？？！！！


## 2.17 完全背包问题
 + https://www.cnblogs.com/labuladong/p/13927918.html
 + 518.零钱兑换II
 + 有一个背包，最大容量为 amount，有一系列物品 coins，每个物品的重量为 coins[i]，每个物品的数量无限。请问有多少种方法，能够把背包恰好装满？
 + 前面讲过的两个背包问题，有一个最大的区别就是，每个物品的数量是无限的，这也就是传说中的「完全背包问题」

 + 第一步要明确两点，「状态」和「选择」。
状态有两个，就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」嘛，背包问题的套路都是这样。

+ 第二步要明确 dp 数组的定义。

首先看看刚才找到的「状态」，有两个，也就是说我们需要一个二维 dp 数组。

dp[i][j] 的定义如下：

若只使用前 i 个物品，当背包容量为 j 时，有 dp[i][j] 种方法可以装满背包。

换句话说，翻译回我们题目的意思就是：

若只使用 coins 中的前 i 个硬币的面值，若想凑出金额 j，有 dp[i][j] 种凑法。

base case 为 dp[0][..] = 0， dp[..][0] = 1。因为如果不使用任何硬币面值，就无法凑出任何金额；如果凑出的目标金额为 0，那么“无为而治”就是唯一的一种凑法。

最终想得到的答案就是 dp[N][amount]，其中 N 为 coins 数组的大小。

```
for (int i = 1; i <= n; i++) {
    for (int j = 1; j <= amount; j++) {
        if (j - coins[i-1] >= 0)
            dp[i][j] = dp[i - 1][j] 
                     + dp[i][j-coins[i-1]];
return dp[N][W]

```

```java
int change(int amount, int[] coins) {
    int n = coins.length;
    int[][] dp = amount int[n + 1][amount + 1];
    // base case
    for (int i = 0; i <= n; i++) 
        dp[i][0] = 1;

    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= amount; j++)
            if (j - coins[i-1] >= 0)
                dp[i][j] = dp[i - 1][j] 
                         + dp[i][j - coins[i-1]];
            else 
                dp[i][j] = dp[i - 1][j];
    }
    return dp[n][amount];
}
```
+ 我们通过观察可以发现，dp 数组的转移只和 dp[i][..] 和 dp[i-1][..] 有关，所以可以压缩状态，进一步降低算法的空间复杂度

```java
int change(int amount, int[] coins) {
    int n = coins.length;
    int[] dp = new int[amount + 1];
    dp[0] = 1; // base case
    for (int i = 0; i < n; i++)
        for (int j = 1; j <= amount; j++)
            if (j - coins[i] >= 0)
                dp[j] = dp[j] + dp[j-coins[i]];
    
    return dp[amount];
}

```
+ 这个解法和之前的思路完全相同，将二维 dp 数组压缩为一维，时间复杂度 O(N*amount)，空间复杂度 O(amount)。

## 2.18 题目套路（打家劫舍）

### 2.18.1 线性排列情况
+ 打家劫舍I
### 2.18.2 环形排列情况
+ 打家劫舍II
### 2.18.3 树形排列情况
+ 打家劫舍III

## 动态规划和回溯算法的关系

+ https://blog.csdn.net/fdl123456/article/details/107054225

+ 给你一个整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

+ 任何算法的核心都是穷举，回溯算法是暴力穷举算法
+ 动态规划之所以比暴力算法快，是因为使用技巧消除了重叠子问题。

+ 变成背包问题的标准形式：

有一个背包，容量为 sum，现在给你 N 个物品，第 i 个物品的重量为 nums[i - 1]（注意 1 <= i <= N），每个物品只有一个，请问你有几种不同的方法能够恰好装满这个背包


+ 对照二维 dp，只要把 dp 数组的第一个维度全都去掉就行了，唯一的区别就是这里的 j 要从后往前遍历，原因如下：

因为二维压缩到一维的根本原理是，dp[j] 和 dp[j-nums[i-1]] 还没被新结果覆盖的时候，相当于二维 dp 中的 dp[i-1][j] 和 dp[i-1][j-nums[i-1]]。

那么，我们就要做到：在计算新的 dp[j] 的时候，dp[j] 和 dp[j-nums[i-1]] 还是上一轮外层 for 循环的结果。

如果你从前往后遍历一维 dp 数组，dp[j] 显然是没问题的，但是 dp[j-nums[i-1]] 已经不是上一轮外层 for 循环的结果了，这里就会使用错误的状态，当然得不到正确的答案。
————————————————
版权声明：本文为CSDN博主「labuladong」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/fdl123456/article/details/107054225

？？？？？？

+ 总结一下，回溯算法虽好，但是复杂度高，即便消除一些冗余计算，也只是「剪枝」，没有本质的改进。而动态规划就比较玄学了，经过各种改造，从一个加减法问题变成子集问题，又变成背包问题，经过各种套路写出解法，又搞出状态压缩，还得反向遍历。

现在搞得我都忘了自己是来干嘛的了。嗯，这也许就是动态规划的魅力吧。
————————————————
版权声明：本文为CSDN博主「labuladong」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/fdl123456/article/details/107054225


# 第三章 数据结构系列

## 3.1 LRU淘汰算法
+ https://www.cnblogs.com/labuladong/p/12320491.html
+ LRU 的全称是 Least Recently Used
+ 哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢。所以结合一下，形成一种新的数据结构：LinkedHashMap。

+ “为什么必须要用双向链表”的问题了，因为我们需要删除操作。删除一个节点不光要得到该节点本身的指针，也需要操作其前驱节点的指针，而双向链表才能支持直接查找前驱，保证操作的时间复杂度 𝑂(1)。

有了双向链表的实现，我们只需要在 LRU 算法中把它和哈希表结合起来即可。


```java
class LRUCache {
    // key -> Node(key, val)
    private HashMap<Integer, Node> map;
    // Node(k1, v1) <-> Node(k2, v2)...
    private DoubleList cache;
    // 最大容量
    private int cap;
    
    public LRUCache(int capacity) {
        this.cap = capacity;
        map = new HashMap<>();
        cache = new DoubleList();
    }
    
    public int get(int key) {
        if (!map.containsKey(key))
            return -1;
        int val = map.get(key).val;
        // 利用 put 方法把该数据提前
        put(key, val);
        return val;
    }
    
    public void put(int key, int val) {
        // 先把新节点 x 做出来
        Node x = new Node(key, val);
        
        if (map.containsKey(key)) {
            // 删除旧的节点，新的插到头部
            cache.remove(map.get(key));
            cache.addFirst(x);
            // 更新 map 中对应的数据
            map.put(key, x);
        } else {
            if (cap == cache.size()) {
                // 删除链表最后一个数据
                Node last = cache.removeLast();
                map.remove(last.key);
            }
            // 直接添加到头部
            cache.addFirst(x);
            map.put(key, x);
        }
    }
}

```

+ todo

## 3.2 LFU算法

+ https://www.cnblogs.com/labuladong/p/13975044.html

+  看caffine实现 TODO

+ 使用一个 HashMap 存储 key 到 freq 的映射，就可以快速操作 key 对应的 freq。

HashMap<Integer, Integer> keyToFreq;

+ LinkedHashSet 顾名思义，是链表和哈希集合的结合体。链表不能快速访问链表节点，但是插入元素具有时序；哈希集合中的元素无序，但是可以对元素进行快速的访问和删除。

## 3.3 二叉搜索树操作集锦

+ https://www.cnblogs.com/labuladong/p/13975087.html

1. 如何把二叉树所有节点中的值加一？递归
2. 如何判断两棵二叉树是否完全相同？递归

### 3.3.1 判断BST的合法性
### 3.3.2 在BST中查找一个树是否存在

+ 针对 BST 的遍历框架！！！
```
void BST(TreeNode root, int target) {
    if (root.val == target)
        // 找到目标，做点什么
    if (root.val < target) 
        BST(root.right, target);
    if (root.val > target)
        BST(root.left, target);
}
```

### 3.3.3 在BST中插入一个数
+ 先找再改

```
TreeNode insertIntoBST(TreeNode root, int val) {
    // 找到空位置插入新节点
    if (root == null) return new TreeNode(val);
    // if (root.val == val)
    //     BST 中一般不会插入已存在元素
    if (root.val < val) 
        root.right = insertIntoBST(root.right, val);
    if (root.val > val) 
        root.left = insertIntoBST(root.left, val);
    return root;
}
```

### 3.3.4 在BST中删除一个数
+ 先找再改

```
TreeNode deleteNode(TreeNode root, int key) {
    if (root.val == key) {
        // 找到啦，进行删除
    } else if (root.val > key) {
        root.left = deleteNode(root.left, key);
    } else if (root.val < key) {
        root.right = deleteNode(root.right, key);
    }
    return root;
}
```

+ 找到目标节点了，比方说是节点 A，如何删除这个节点，这是难点。因为删除节点的同时不能破坏 BST 的性质。
情况 1：A 恰好是末端节点，两个子节点都为空，那么它可以当场去世了。
情况 2：A 只有一个非空子节点，那么它要让这个孩子接替自己的位置。
情况 3：A 有两个子节点，麻烦了，为了不破坏 BST 的性质，A 必须找到左子树中最大的那个节点，或者右子树中最小的那个节点来接替自己。

```
TreeNode deleteNode(TreeNode root, int key) {
    if (root == null) return null;
    if (root.val == key) {
        // 这两个 if 把情况 1 和 2 都正确处理了
        if (root.left == null) return root.right;
        if (root.right == null) return root.left;
        // 处理情况 3
        TreeNode minNode = getMin(root.right);
        root.val = minNode.val;
        root.right = deleteNode(root.right, minNode.val);
    } else if (root.val > key) {
        root.left = deleteNode(root.left, key);
    } else if (root.val < key) {
        root.right = deleteNode(root.right, key);
    }
    return root;
}

TreeNode getMin(TreeNode node) {
    // BST 最左边的就是最小的
    while (node.left != null) node = node.left;
    return node;
} 

```

+ 删除操作就完成了。注意一下，这个删除操作并不完美，因为我们一般不会通过 root.val = minNode.val 修改节点内部的值来交换节点，而是通过一系列略微复杂的链表操作交换 root 和 minNode 两个节点。因为具体应用中，val 域可能会很大，修改起来很耗时，而链表操作无非改一改指针，而不会去碰内部数据。

## 3.4 完全二叉树的结点数

+ https://www.cnblogs.com/labuladong/p/13975114.html

+ 一棵完全二叉树的两棵子树，至少有一棵是满二叉树

+ 由于完全二叉树的性质，其子树一定有一棵是满的，所以一定会触发 hl == hr，只消耗 O(logN) 的复杂度而不会继续递归。

综上，算法的递归深度就是树的高度 O(logN)，每次递归所花费的时间就是 while 循环，需要 O(logN)，所以总体的时间复杂度是 O(logN*logN)。

所以说，「完全二叉树」这个概念还是有它存在的原因的，不仅适用于数组实现二叉堆，而且连计算节点总数这种看起来简单的操作都有高效的算法实现。

```
public int countNodes(TreeNode root) {
    TreeNode l = root, r = root;
    // 记录左、右子树的高度
    int hl = 0, hr = 0;
    while (l != null) {
        l = l.left;
        hl++;
    }
    while (r != null) {
        r = r.right;
        hr++;
    }
    // 如果左右子树的高度相同，则是一棵满二叉树
    if (hl == hr) {
        return (int)Math.pow(2, hl) - 1;
    }
    // 如果左右高度不同，则按照普通二叉树的逻辑计算
    return 1 + countNodes(root.left) + countNodes(root.right);
}

```

## 3.5 各种遍历框架序列话和反序列化二叉树

+ leetcode 做过

## 3.6 二叉树最近公共祖先

## 3.7 单调栈

+ https://www.cnblogs.com/labuladong/p/13975051.html

+ 单调栈实际上就是栈，只是利用了一些巧妙的逻辑，使得每次新元素入栈后，栈内的元素都保持有序（单调递增或单调递减）。

+ for 循环要从后往前扫描元素，因为我们借助的是栈的结构，倒着入栈，其实是正着出栈。while 循环是把两个“高个”元素之间的元素排除，因为他们的存在没有意义，前面挡着个“更高”的元素，所以他们不可能被作为后续进来的元素的 Next Great Number 了。

### 3.7.3 Next Great Number 
+ 给你一个数组 T = [73, 74, 75, 71, 69, 72, 76, 73]，这个数组存放的是近几天的天气气温（这气温是铁板烧？不是的，这里用的华氏度）。你返回一个数组，计算：对于每一天，你还要至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填 0 。

举例：给你 T = [73, 74, 75, 71, 69, 72, 76, 73]，你返回 [1, 1, 4, 2, 1, 1, 0, 0]。

解释：第一天 73 华氏度，第二天 74 华氏度，比 73 大，所以对于第一天，只要等一天就能等到一个更暖和的气温。后面的同理。


```
vector<int> dailyTemperatures(vector<int>& T) {
    vector<int> ans(T.size());
    stack<int> s; // 这里放元素索引，而不是元素
    for (int i = T.size() - 1; i >= 0; i--) {
        while (!s.empty() && T[s.top()] <= T[i]) {
            s.pop();
        }
        ans[i] = s.empty() ? 0 : (s.top() - i); // 得到索引间距
        s.push(i); // 加入索引，而不是元素
    }
    return ans;
}

```

### 3.7.3 如何处理循环数组

+ 计算机的内存都是线性的，没有真正意义上的环形数组，但是我们可以模拟出环形数组的效果，一般是通过 % 运算符求模（余数），获得环形特效
+ 将原始数组“翻倍”，就是在后面再接一个原始数组，这样的话，按照之前“比身高”的流程，每个元素不仅可以比较自己右边的元素，而且也可以和左边的元素比较了

```
vector<int> nextGreaterElements(vector<int>& nums) {
    int n = nums.size();
    vector<int> res(n); // 存放结果
    stack<int> s;
    // 假装这个数组长度翻倍了
    for (int i = 2 * n - 1; i >= 0; i--) {
        while (!s.empty() && s.top() <= nums[i % n])
            s.pop();
        res[i % n] = s.empty() ? -1 : s.top();
        s.push(nums[i % n]);
    }
    return res;
}

```

## 3.8 单调队列

## 3.9 如何判断回文链表

+ https://www.cnblogs.com/labuladong/p/12320497.html

+ 寻找回文串的核心思想是从中心向两端扩展

+ 先通过「双指针技巧」中的快慢指针来找到链表的中点
+ 从slow开始反转后面的链表，现在就可以开始比较回文串了

```
ListNode left = head;
ListNode right = reverse(slow);

while (right != null) {
    if (left.val != right.val)
        return false;
    left = left.next;
    right = right.next;
}
return true;

```

+ 首先，寻找回文串是从中间向两端扩展，判断回文串是从两端向中间收缩。对于单链表，无法直接倒序遍历，可以造一条新的反转链表，可以利用链表的后序遍历，也可以用栈结构倒序处理单链表。

具体到回文链表的判断问题，由于回文的特殊性，可以不完全反转链表，而是仅仅反转部分链表，将空间复杂度降到 O(1)。

## 3.10 递归反转链表

## k个一组反转链表

# 第四章 算法思维系列

## 4.1 回溯算法解决子集、组合、排列问题

+ https://www.cnblogs.com/labuladong/p/13945718.html

+ 求子集（subset），求排列（permutation），求组合（combination）

```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择

```

### 子集

###  组合
+ k 限制了树的高度，n 限制了树的宽度

### 排列
+ 组合和子集问题都需要一个start变量防止重复，而排列不需要防重复。


就是排列组合和子集三个问题的解法，总结一下：

子集问题可以利用数学归纳思想，假设已知一个规模较小的问题的结果，思考如何推导出原问题的结果。也可以用回溯算法，要用 start 参数排除已选择的数字。

组合问题利用的是回溯思想，结果可以表示成树结构，我们只要套用回溯算法模板即可，关键点在于要用一个 start 排除已经选择过的数字。

排列问题是回溯思想，也可以表示成树结构套用算法模板，关键点在于使用 contains 方法排除已经选择的数字，前文有详细分析，这里主要是和组合问题作对比。

记住这几种树的形状，就足以应对大部分回溯算法问题了，无非就是 start 或者 contains 剪枝，也没啥别的技巧了。

## 4.2 回溯算法最佳实践：解数独

## 4.3 回溯算法最佳实践：括号生成
+ https://www.cnblogs.com/labuladong/p/13975768.html

+ 括号问题可以简单分成两类，一类是前文写过的 括号的合法性判断 ，一类是合法括号的生成。对于括号合法性的判断，主要是借助「栈」这种数据结构，而对于括号的生成，一般都要利用回溯递归的思想。

## 4.4 BFS算法暴力破解各种智力题

### 滑动拼图（华容道）

## 4.5 2Sum问题

### 4.5.1 2Sum I

+ 给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

+ hash表记录元素值到索引的映射

### 4.5.2 2Sum II

## 4.6 nSum问题

### 4.6.1 2Sum
### 4.6.2 3Sum
### 4.6.3 4Sum

### 4.6.4 100Sum


TODO


## 4.7 实现计算器

+ 多个步骤，多个模块，分而治之

## 4.8 摊烧饼
+ ？？？

## 4.9 前缀和技巧解决子数组问题 !!!

+ https://www.cnblogs.com/labuladong/p/13976538.html

给你一个整数数组 nums 和一个整数k ，请你统计并返回 该数组中和为k的连续子数组的个数。



示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 4.9.1 前缀和

+ 前缀和的思路是这样的，对于一个给定的数组 nums，我们额外开辟一个前缀和数组进行预处理：

```
int n = nums.length;
// 前缀和数组
int[] preSum = new int[n + 1];
preSum[0] = 0;
for (int i = 0; i < n; i++)
    preSum[i + 1] = preSum[i] + nums[i];

```

```
int subarraySum(int[] nums, int k) {
    int n = nums.length;
    // 构造前缀和
    int[] sum = new int[n + 1];
    sum[0] = 0; 
    for (int i = 0; i < n; i++)
        sum[i + 1] = sum[i] + nums[i];
    
    int ans = 0;
    // 穷举所有子数组
    for (int i = 1; i <= n; i++)
        for (int j = 0; j < i; j++)
            // sum of nums[j..i-1]
            if (sum[i] - sum[j] == k)
                ans++;

    return ans;
}

```


+ 优化的思路是：我直接记录下有几个 sum[j] 和 sum[i] - k 相等，直接更新结果，就避免了内层的 for 循环。我们可以用哈希表，在记录前缀和的同时记录该前缀和出现的次数。


+ 统计班上同学考试成绩在不同分数段的百分比，也可以利用前缀和技巧 !!!!

```
int[] scores; // 存储着所有同学的分数
// 试卷满分 150 分
int[] count = new int[150 + 1]
// 记录每个分数有几个同学
for (int score : scores)
    count[score]++
// 构造前缀和
for (int i = 1; i < count.length; i++)
    count[i] = count[i] + count[i-1];

```


+ 这样，给你任何一个分数段，你都能通过前缀和相减快速计算出这个分数段的人数，百分比也就很容易计算了。


## 4.10 扁平化嵌套列表

# 第五章 高频面试系列

## 5.1 如何高效寻找素数

+ 如果在 [2,sqrt(n)] 这个区间之内没有发现可整除因子，就可 以直接断定 n 是素数了，因为在区间 [sqrt(n),n] 也一定不会发现可整 除因子。

```
int countPrimes(int n) {
    boolean[] isPrim = new boolean[n];
    Arrays.fill(isPrim, true);
    for (int i = 2; i * i < n; i++)
        if (isPrim[i])
            for (int j = i * i; j < n; j += i)
                isPrim[j] = false;
    int count = 0;
    for (int i = 2; i < n; i++)
        if (isPrim[i]) count++;
    return count;
}
```

## 5.2 如何高效进行模幂运算

## 5.3 如何运用二分搜索算法

+ 当搜索空间有序的时候，就可以通过二分搜索“剪枝”，大幅提升效率
+ 替代线性搜索

## 5.4 如何高效解决接雨水问题

+ https://www.cnblogs.com/labuladong/p/12320514.html

+ 双指针解法 ！

```
int trap(vector<int>& height) {
    if (height.empty()) return 0;
    int n = height.size();
    int left = 0, right = n - 1;
    int ans = 0;
    
    int l_max = height[0];
    int r_max = height[n - 1];
    
    while (left <= right) {
        l_max = max(l_max, height[left]);
        r_max = max(r_max, height[right]);
        
        // ans += min(l_max, r_max) - height[i]
        if (l_max < r_max) {
            ans += l_max - height[left];
            left++; 
        } else {
            ans += r_max - height[right];
            right--;
        }
    }
    return ans;
}

```

## 5.5 如何去除有序数组的重复元素

+ 对于数组相关的算法问题，有一个通用的技巧:要尽量避免在中间删 除元素，那我就想先办法把这个元素换到最后去 。这样的话，最终待删除的 元素都拖在数组尾部，一个一个 pop 掉就行了，每次操作的时间复杂度也就 降低到 O(1) 了。
+ 按照这个思路呢，又可以衍生出解决类似需求的通用方式:双指针技巧。具体一点说，应该是快慢指针。 ！！！

+ 我们让慢指针 slow 走左后面，快指针 fast 走在前面探路，找到一个不 重复的元素就告诉 slow 并让 slow 前进一步。这样当 fast 指针遍历完 整个数组 nums 后， nums[0..slow] 就是不重复元素，之后的所有元素都 是重复元素。

```
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
```

## 5.6如何寻找最长回文子串

## 5.7 如何运用贪心算法玩跳跃游戏

### 5.7.1 跳跃游戏 I
### 5.7.2 跳跃游戏 II

## 5.8 如何用贪心算法做时间管理

## 5.9 如何判定括号的合法性
+ 对括号的合法性判断是一个很常⻅且实用的问题，比如说我们写的代码，编 辑器和编译器都会检查括号是否正确闭合。而且我们的代码可能会包含三种 括号 [](){} ，判断起来有一点难度。
本文就来聊一道关于括号合法性判断的算法题，相信能加深你对栈这种数据 结构的理解。

```
bool isValid(string str) {
    stack<char> left;
    for (char c : str) {
        if (c == '(' || c == '{' || c == '[')
            left.push(c);
        else // 字符 c 是右括号
            if (!left.empty() && leftOf(c) == left.top())
                left.pop();
            else
                // 和最近的左括号不匹配
                return false;
    }
    // 是否所有的左括号都被匹配了
    return left.empty();
}

char leftOf(char c) {
    if (c == '}') return '{';
    if (c == ')') return '(';
    return '[';
}

```

## 5.10 如何调度考试的座位
+ https://www.cnblogs.com/labuladong/p/13945728.html
+ 这是 LeetCode 第 885 题，有趣且具有一定技巧性。

+ 先来描述一下题目:假设有一个考场，考场有一排共 N 个座位，索引分别 是 [0..N-1] ，考生会陆续进入考场考试，并且可能在任何时候离开考场。
你作为考官，要安排考生们的座位，满足:每当一个学生进入时，你需要最 大化他和最近其他人的距离;如果有多个这样的座位，安排到他到索引最小 的那个座位。

+ 使用的 TreeSet 就是一个有序集合，目的就是为了保持线段⻓度的有 序性，快速查找最大线段，快速删除和插入。



## 5.11 Union-Find算法详解

+  Union-Find 算法，也就是常说的并查集算法，主要是解决图论中 「动态连通性」问题的。

## 5.12 Union-Find算法应用

## 5.13 一行代码就能解决的算法题

### 5.13.1 Nim游戏

+ 游戏规则是这样的:你和你的朋友面前有一堆石子，你们轮流拿，一次至少
拿一颗，最多拿三颗，谁拿走最后一颗石子谁获胜。
假设你们都很聪明，由你第一个开始拿，请你写一个算法，输入一个正整数
n，返回你是否能赢(true 或 false)。
比如现在有 4 颗石子，算法应该返回 false。因为无论你拿 1 颗 2 颗还是 3
颗，对方都能一次性拿完，拿走最后一颗石子，所以你一定会输。

+ 我们发现只要踩到 4 的倍数，就落入了圈套，永远逃不 出 4 的倍数，而且一定会输。所以这道题的解法非常简单

```
bool canWinNim(int n) {
// 如果上来就踩到 4 的倍数，那就认输吧 // 否则，可以把对方控制在 4 的倍数，必胜 return n % 4 != 0;
}
```

### 5.13.2 石子游戏

+ 游戏规则是这样的:你和你的朋友面前有一排石头堆，用一个数组 piles 表 示，piles[i] 表示第 i 堆石子有多少个。你们轮流拿石头，一次拿一堆，但是 只能拿走最左边或者最右边的石头堆。所有石头被拿完后，谁拥有的石头 多，谁获胜。
假设你们都很聪明，由你第一个开始拿，请你写一个算法，输入一个数组 piles，返回你是否能赢(true 或 false)。
注意，石头的堆的数量为偶数，所以你们两人拿走的堆数一定是相同的。石
头的总数为奇数，也就是你们最后不可能拥有相同多的石头，一定有胜负之
分。

+ 只要对规则深入思考，就会大惊失色:只要你足够聪明，你是必胜无疑
的，因为你是先手。

+ 这是为什么呢，因为题目有两个条件很重要:一是石头总共有偶数堆，石头 的总数是奇数。这两个看似增加游戏公平性的条件，反而使该游戏成为了一 个割韭菜游戏。我们以 piles=[2, 1, 9, 5] 讲解，假设这四堆石头从左到 右的索引分别是 1，2，3，4。
如果我们把这四堆石头按索引的奇偶分为两组，即第 1、3 堆和第 2、4 堆， 那么这两组石头的数量一定不同，也就是说一堆多一堆少。因为石头的总数 是奇数，不能被平分。
而作为第一个拿石头的人，你可以控制自己拿到所有偶数堆，或者所有的奇
数堆。

+ 也就是说，你可以在第一步就观察好，奇数堆的石头总数多，还是偶数堆的
石头总数多，然后步步为营，就一切尽在掌控之中了。知道了这个漏洞，可
以整一整不知情的同学了。



### 5.13.3 电灯开关问题

+ 这个问题是这样描述的:有 n 盏电灯，最开始时都是关着的。现在要进行 n 轮操作:
第 1 轮操作是把每一盏电灯的开关按一下(全部打开)。
第 2 轮操作是把每两盏灯的开关按一下(就是按第 2，4，6... 盏灯的开关，
它们被关闭)。
第 3 轮操作是把每三盏灯的开关按一下(就是按第 3，6，9... 盏灯的开关，
有的被关闭，比如 3，有的被打开，比如 6)... 如此往复，直到第 n 轮，即只按一下第 n 盏灯的开关。
现在给你输入一个正整数 n 代表电灯的个数，问你经过 n 轮操作后，这些电 灯有多少盏是亮的?

```
int bulbSwitch(int n) {
    return (int)Math.sqrt(n);
}
```

---


+ split -b 40m labuladong的算法小抄完整版.pdf labuladong的算法小抄完整版.part-
+ cat labuladong的算法小抄完整版.part-* > ~/downloads/labuladong的算法小抄完整版.pdf