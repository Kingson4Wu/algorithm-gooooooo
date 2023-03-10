package nowcoder

/**

题目主要信息：
矩阵内是非负数，求最长的递增路径的长度
移动方向可以是上下左右，不能超出边界，这将是递归的判定条件
同一条路径不能有重复的单元格，需要有记忆
举一反三：
学习完本题的思路你可以解决如下题目：

BM57. 岛屿数量

方法一：深度优先搜索(推荐使用)
知识点：深度优先搜索（dfs）

深度优先搜索一般用于树或者图的遍历，其他有分支的（如二维矩阵）也适用。它的原理是从初始点开始，一直沿着同一个分支遍历，直到该分支结束，然后回溯到上一级继续沿着一个分支走到底，如此往复，直到所有的节点都有被访问到。

思路：

既然是查找最长的递增路径长度，那我们首先要找到这个路径的起点，起点不好直接找到，就从上到下从左到右遍历矩阵的每个元素。然后以每个元素都可以作为起点查找它能到达的最长递增路径。

如何查找以某个点为起点的最长递增路径呢？我们可以考虑深度优先搜索，因为我们查找递增路径的时候，每次选中路径一个点，然后找到与该点相邻的递增位置，相当于进入这个相邻的点，继续查找递增路径，这就是递归的子问题。因此递归过程如下：

终止条件： 进入路径最后一个点后，四个方向要么是矩阵边界，要么没有递增的位置，路径不能再增长，返回上一级。
返回值： 每次返回的就是本级之后的子问题中查找到的路径长度加上本级的长度。
本级任务： 每次进入一级子问题，先初始化后续路径长度为0，然后遍历四个方向（可以用数组表示，下标对数组元素的加减表示去往四个方向），进入符合不是边界且在递增的邻近位置作为子问题，查找子问题中的递增路径长度。因为有四个方向，所以最多有四种递增路径情况，因此要维护当级子问题的最大值。
具体做法：

step 1：使用一个dp数组记录i，j处的单元格拥有的最长递增路径，这样在递归过程中如果访问到就不需要重复访问。
step 2：遍历矩阵每个位置，都可以作为起点，并维护一个最大的路径长度的值。
step 3：对于每个起点，使用dfs查找最长的递增路径：只要下一个位置比当前的位置数字大，就可以深入，同时累加路径长度。


复杂度分析：

时间复杂度：O(mn)，m、n分别为矩阵的两边，遍历整个矩阵以每个点作为起点，然后递归相当于遍历填充dp矩阵
空间复杂度：O(mn)，辅助矩阵的空间是一个二维数组


import java.util.*;
public class Solution {
    //记录四个方向
    private int[][] dirs = new int[][] {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    private int n, m;
    //深度优先搜索，返回最大单元格数
    public int dfs(int[][] matrix, int[][] dp, int i, int j) {
        if(dp[i][j] != 0)
            return dp[i][j];
        dp[i][j]++;
        for (int k = 0; k < 4; k++) {
            int nexti = i + dirs[k][0];
            int nextj = j + dirs[k][1];
            //判断条件
            if(nexti >= 0 && nexti < n && nextj >= 0 && nextj < m && matrix[nexti][nextj] > matrix[i][j])
                dp[i][j] = Math.max(dp[i][j], dfs(matrix, dp, nexti, nextj) + 1);
        }
        return dp[i][j];
    }
    public int solve (int[][] matrix) {
        //矩阵不为空
        if (matrix.length == 0 || matrix[0].length == 0)
            return 0;
        int res = 0;
        n = matrix.length;
        m = matrix[0].length;
        //i，j处的单元格拥有的最长递增路径
        int[][] dp = new int[m + 1][n + 1];
        for(int i = 0; i < n; i++)
            for(int j = 0; j < m; j++)
                //更新最大值
                res = Math.max(res, dfs(matrix, dp, i, j));
        return res;
    }
}




方法二：广度优先搜索（扩展思路）
知识点：广度优先搜索（bfs）

广度优先搜索与深度优先搜索不同，它是将与某个节点直接相连的其它所有节点依次访问一次之后，再往更深处，进入与其他节点直接相连的节点。bfs的时候我们常常会借助队列的先进先出，因为从某个节点出发，我们将与它直接相连的节点都加入队列，它们优先进入，则会优先弹出，在它们弹出的时候再将与它们直接相连的节点加入，由此就可以依次按层访问。

思路：

我们可以将矩阵看成是一个有向图，一个元素到另一个元素递增，代表有向图的箭头。这样我们可以根据有向图的出度入度找到最长的路径，且这个路径在矩阵中就是递增的。

具体做法：

step 1：计算每个节点（单元格）所对应的出度（符合边界条件且递增），对于作为边界条件的单元格，它的值比所有的相邻单元格的值都要大，因此作为边界条件的单元格的出度都为0。利用一个二维矩阵记录每个单元格的出度
step 2：利用拓扑排序的思想，从所有出度为0的单元格开始进行广度优先搜索。
step 3：借助队列来广度优先搜索，队列中每次加入出度为0的点，即路径最远点，每次从A点到B点，便将A点出度减一。
step 4：每次搜索都会遍历当前层的所有单元格，更新其余单元格的出度，并将出度变为0的单元格加入下一层搜索。
step 5：当搜索结束时，搜索的总层数即为矩阵中的最长递增路径的长度，因为bfs的层数就是路径增长的层数。


复杂度分析：

时间复杂度：O(mn)，m、n分别为矩阵的两边，相当于遍历整个矩阵两次
空间复杂度：O(mn)，辅助矩阵的空间是一个二维数组
*/
