
+ TODO
# 图

## 有向图
+ 入度表 indegrees（key为目标顶点, value为源顶点数值）
+ 邻接表 adjacency (key为源顶点, value为目标顶点列表）
+ DFS，递归！
+ 课程表（拓扑排序：入度表BFS法 / DFS法，清晰图解）: https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/

## DFS
+ DFS算法可以被认为是回溯算法！！！！！
+ 递归回溯
+ 三种类型
    - 子集型
    - 组合型
    - 排列型

+ 子集型
+ exercise/leetcode/hot100/78. 子集_test.go
```go
func subsetsDfs(nums []int) (ans [][]int) {

	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])//选
		dfs(cur + 1)
		set = set[:len(set)-1]//不选
		dfs(cur + 1)
	}
	dfs(0)
	return
}
```

+ 排列型
+ exercise/leetcode/hot100/46. 全排列_test.go
```go
func permute(nums []int) [][]int {

	var ans [][]int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), nums...))
			return
		}
		for i := cur; i < len(nums); i++ {
			nums[cur], nums[i] = nums[i], nums[cur]
			dfs(cur + 1)
			nums[cur], nums[i] = nums[i], nums[cur]
			//dfs(cur + 1)
		}
	}
	dfs(0)
	return ans
}
```

+ 组合型
+ exercise/leetcode/77. 组合_test.go
+ 注意 if len(set) == k 和 dfs(i + 1)
```go
func combine(n int, k int) [][]int {

	var ans [][]int
	var set []int

	var dfs func(int)
	dfs = func(cur int) {
		if len(set) == k {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		for i := cur; i < n; i++ {
			set = append(set, i+1)
			dfs(i + 1)
			set = set[:len(set)-1]
		}
	}

	dfs(0)
	return ans
}
```

## BFS
+ TODO



----


# TREE

+ 树遍历：先序、中序、后序（打印）；深度遍历，广度遍历（搜索）
+ 搜索树就是有序的意思（中序）- 二叉搜索树满足每个节点的左子树上的所有节点均严格小于当前节点且右子树上的所有节点均严格大于当前节点。（注意不能等于！！！！）
+ 树：left和right都为空才是叶子节点
+ 搜索树倒序，迭代法：右子树要全部入栈，直到右子数为空
  处理完当前结点，再处理左结点，重复左结点的右子树的入栈操作
  （er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof）
+ 完全二叉树--若二叉树的深度为 h，除第 h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）
  判断是否完全二叉树：广度优先搜索，队列，判断是不是连续两个不存在
  nowcoder/BM35 判断是不是完全二叉树.go
+ 平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
+ 堆的结构可以分为大顶堆和小顶堆，是一个完全二叉树


## 二叉树
+ 二叉树遍历(先序、中序、后序)
    - 先序（中左右），中序（左中右），后序（左右中）
    - 递归遍历和非递归遍历
    - 二叉树遍历(先序、中序、后序):<https://www.jianshu.com/p/456af5480cee>

+ 写代码前，先画图整理好思路！！！！

+ 匿名函数写法！！

## 题目

### 遍历

+ 前序遍历使用递归即可
+ 中序遍历使用栈（有递归解法....）
  `nowcoder/BM24 二叉树的中序遍历.go`
  递归真的好难理解！！！

```go
    //中序
   public void inorder(List<Integer> list, TreeNode root){
        //遇到空节点则返回
        if(root == null)
            return;
        //先去左子树
        inorder(list, root.left);
        //再访问根节点
        list.add(root.val);
        //最后去右子树
        inorder(list, root.right);
    }

```
+ 后序遍历同样可以使用递归（和中序递归一样处理）

### 对称的二叉树

+ nowcoder/BM31 对称的二叉树.go
  还是用递归！！！
```java
public class Solution {
    boolean recursion(TreeNode root1, TreeNode root2){
        //可以两个都为空
        if(root1 == null && root2 == null)
            return true;
        //只有一个为空或者节点值不同，必定不对称
        if(root1 == null || root2 == null || root1.val != root2.val)
            return false;
        //每层对应的节点进入递归比较
        return recursion(root1.left, root2.right) && recursion(root1.right, root2.left);
    }
    boolean isSymmetrical(TreeNode pRoot) {
        return recursion(pRoot, pRoot);
    }
}
```

### 两个指定节点的最近公共祖先

+ 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先！！！

方法一：搜索路径比较（推荐使用）
step 1：根据二叉搜索树的性质，从根节点开始查找目标节点，当前节点比目标小则进入右子树，当前节点比目标大则进入左子树，直到找到目标节点。这个过程成用数组记录遇到的元素。
step 2：分别在搜索二叉树中找到p和q两个点，并记录各自的路径为数组。
step 3：同时遍历两个数组，比较元素值，最后一个相等的元素就是最近的公共祖先。


方法二：一次遍历（扩展思路）
step 1：首先检查空节点，空树没有公共祖先。
step 2：对于某个节点，比较与p、q的大小，若p、q在该节点两边说明这就是最近公共祖先。
step 3：如果p、q都在该节点的左边，则递归进入左子树。
step 4：如果p、q都在该节点的右边，则递归进入右子树。

//pq在该节点两边说明这就是最近公共祖先(因为从根节点遍历的，那么最先遇到符合条件的就是最近公共祖先！！！！)
if((p >= root.val && q <= root.val) || (p <= root.val && q >= root.val))
return root.val;


+ nowcoder/BM38 在二叉树中找到两个节点的最近公共祖先.go
  (1)深度优先搜索，回溯法递归找路径，比较两个路径第一个不相同的前一个就是公共祖先
```java
 public boolean flag = false;
    //求得根节点到目标节点的路径
    public void dfs(TreeNode root, ArrayList<Integer> path, int o){
        if(flag || root == null)
            return;
        path.add(root.val);
        //节点值都不同，可以直接用值比较
        if(root.val == o){
            flag = true;
            return;
        }
        //dfs遍历查找
        dfs(root.left, path, o);
        dfs(root.right, path, o);
        //找到 - 在子结点找到了
        if(flag)
            return;
        //回溯
        path.remove(path.size() - 1);
    }
```
(2)二叉树递归

思路：

我们也可以讨论几种情况：

step 1：如果o1和o2中的任一个和root匹配，那么root就是最近公共祖先。
step 2：如果都不匹配，则分别递归左、右子树。
step 3：如果有一个节点出现在左子树，并且另一个节点出现在右子树，则root就是最近公共祖先.
step 4：如果两个节点都出现在左子树，则说明最低公共祖先在左子树中，否则在右子树。
step 5：继续递归左、右子树，直到遇到step1或者step3的情况。

### 根据前序和中序重建二叉树

+ 根据前序和中序重建二叉树
  （1）前序第一个为根节点
  （2）找到中序中根节点的位置，分成两组，分治递归构造二叉树

前序和中序的特点：
前序=【左子树】+【根节点】+【右子树】
中序=【根节点】+【左子树】+【右子树】

### 序列化二叉树

+ nowcoder/BM39 序列化二叉树.go
  (1)一个字符串的前序遍历（标识节点分割和空节点），可以重建二叉树（递归方式）
  (2)一个字符串的前序遍历和一个字符串的中序遍历（标识节点分割），可以重建二叉树（递归方式）

### 二叉树的右视图

+ nowcoder/BM41 输出二叉树的右视图.go -- 层次遍历！！！
  请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图
  构建完成二叉树之后，我们只需要对二叉树进行层次遍历。遍历到每一层时，将该层的最右边节点加入数组即可。

int size = queue.size();
for(int i = 0; i < size; i++){
if(i == size - 1) result.add(temp.val);
}

```java
 public ArrayList<Integer> right(TreeNode root){
        ArrayList<Integer> result = new ArrayList<Integer>();
        if(root == null)
            return result;
        // 层次遍历
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offer(root);
        while(!queue.isEmpty()){
            int size = queue.size();
            for(int i = 0; i < size; i++){
                TreeNode temp = queue.poll();
                // 记录每一层的最右边节点
                if(i == size - 1) result.add(temp.val);
                if(temp.left != null) queue.offer(temp.left);
                if(temp.right != null) queue.offer(temp.right);
            }
        }
        return result;
    }
/**
————————————————
版权声明：本文为CSDN博主「木水先生」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_39963192/article/details/115119704
*/

```

### 二叉搜索树中第K小的元素
+ 非递归的中序，使用栈，需要记录状态提前结束遍历的时候，使用栈的方式遍历，代码不会那么丑
  leetcode/tree/230. 二叉搜索树中第K小的元素.go


### 不同的二叉搜索树
+ leetcode/tree/95. 不同的二叉搜索树 II.go
  回溯递归！！！

```go
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }
    return helper(1, n)
}

func helper(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    allTrees := []*TreeNode{}
    // 枚举可行根节点
    for i := start; i <= end; i++ {
        // 获得所有可行的左子树集合
        leftTrees := helper(start, i - 1)
        // 获得所有可行的右子树集合
        rightTrees := helper(i + 1, end)
        // 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                currTree := &TreeNode{i, nil, nil}
                currTree.Left = left
                currTree.Right = right
                allTrees = append(allTrees, currTree)
            }
        }
    }
    return allTrees
}

```

### 二叉树的层级遍历

---

# 搜索

+ 树的遍历方法有广度优先（层序遍历），以及深度优先两种方法，分成先序遍历，中序遍历，后序遍历三种。
+ 深度优先使用栈，广度优先使用队列

## 深度优先搜索
+ 递归地实现 DFS 时，似乎不需要使用任何栈。但实际上，我们使用的是由系统提供的隐式栈，也称为调用栈（Call Stack）

## 栈实现

### 先序

(1)`root!=nil || len(stack)>0`
(2) root!=nil 处理逻辑，入栈，左子树赋值
(3) else 出栈，右子树赋值

```go
func (root *TreeNode) preorder() []int{       //非递归前序遍历
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}           //定义一个栈储存节点信息
	for root!=nil || len(stack)!=0{
		if root!=nil{
			res=append(res,root.data)
			stack=append(stack,root)        
			root=root.Lchild
		}else{
			root=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			root=root.Rchild
		}
	}
	return res
}

```

+ 中序

(1)`root!=nil || len(stack)>0`
(2) root!=nil 左子树循环入栈
(3)出栈，处理逻辑
(4)`root=root.Right` 右子树赋值

``` go
func (root *TreeNode) inorder()[]int{
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}
	for root!=nil || len(stack)!=0{
		if root!=nil{
			stack=append(stack,root)
			root=root.Lchild
		}else{
			root=stack[len(stack)-1]
			res=append(res,root.data)
			stack=stack[:len(stack)-1]
			root=root.Rchild
		}
	}
	return res
}
```

+ 判断一颗二叉树是否为二叉查询树
```go
func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}

```

+ 后序
  有点复杂，日

```go
func (root *TreeNode)postorder() []int {
	res:=[]int{}
	if root==nil{
		return res
	}
	stack:=[]*TreeNode{}
	pre:=&TreeNode{}
	stack=append(stack,root)
	for len(stack)!=0{
		cur:=stack[len(stack)-1]
		if (cur.Lchild==nil && cur.Rchild==nil) || (pre!=nil &&(pre==cur.Lchild || pre==cur.Rchild)){
			res=append(res,cur.data)
			pre=cur
			stack=stack[:len(stack)-1]
		}else{
			if cur.Rchild!=nil{
				stack=append(stack,cur.Rchild)
			}
			if cur.Lchild!=nil{
				stack=append(stack,cur.Lchild)
			}
		}
	}
	return res
}

```


递归

+ 深度优先搜索

+ 找树左下角的值 ： 给定一个二叉树，在树的最后一行找到最左边的值。
```go

func findBottomLeftValue(root *TreeNode) (curVal int) {
    curHeight := 0
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, height int) {
        if node == nil {
            return
        }
        height++
        dfs(node.Left, height)
        dfs(node.Right, height)
		//因为左边先递归，所以左边先赋值给curHeight， 所以val也是左边的值
        if height > curHeight {
            curHeight = height
            curVal = node.Val
        }
    }
    dfs(root, 0)
    return
}

```

+ 广度优先（层序遍历）
```java
void LevelOrderTraversal( BinTree BT)
{
    Queue Q;
    BinTree T;
    if(!BT) return;  //若是空树则直接返回
    Q = CreatQueue(maxsize); //创建并初始化队列
    AddQ(Q,BT);
    while(!IsEmpty(s))
    {
        T=deleteQ(Q);     //出队
        printf("%d\n",T->data);  //访问该节点
        if(T->Left) AddQ(Q,T->Left);   //分别将其左右子入队
        if(T->Right) AddQ(Q,T->Right);  
    }
}

```

### 修剪二叉搜索树
+ 给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树 不应该 改变保留在树中的元素的相对结构 (即，如果没有被移除，原有的父代子代关系都应当保留)。 可以证明，存在 唯一的答案 。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

+ leetcode/tree/669. 修剪二叉搜索树.go
```java
class Solution {
    public TreeNode trimBST(TreeNode root, int L, int R) {
        if (root == null) return root;
        if (root.val > R) return trimBST(root.left, L, R);
        if (root.val < L) return trimBST(root.right, L, R);

        root.left = trimBST(root.left, L, R);
        root.right = trimBST(root.right, L, R);
        return root;
    }
}

```

### 二叉树展开为链表

+ leetcode/tree/114. 二叉树展开为链表.go
  前两种方法都借助前序遍历，前序遍历过程中需要使用栈存储节点。有没有空间复杂度是 O(1) 的做法呢？

注意到前序遍历访问各节点的顺序是根节点、左子树、右子树。如果一个节点的左子节点为空，则该节点不需要进行展开操作。
如果一个节点的左子节点不为空，则该节点的左子树中的最后一个节点被访问之后，该节点的右子节点被访问。
该节点的左子树中最后一个被访问的节点是左子树中的最右边的节点，也是该节点的前驱节点。
因此，问题转化成寻找当前节点的前驱节点。

具体做法是，对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点，
将当前节点的右子节点赋给前驱节点的右子节点，然后将当前节点的左子节点赋给当前节点的右子节点，
并将当前节点的左子节点设为空。对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。
```go

func flatten(root *TreeNode)  {
    curr := root
    for curr != nil {
        if curr.Left != nil {
            next := curr.Left
            predecessor := next
            for predecessor.Right != nil {
                predecessor = predecessor.Right
            }
            predecessor.Right = curr.Right
            curr.Left, curr.Right = nil, next
        }
        curr = curr.Right
    }
}

```

### 填充每个节点的下一个右侧节点指针
+ 给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
  struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
  }
  填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
  初始状态下，所有 next 指针都被设置为 NULL。

+ leetcode/tree/117. 填充每个节点的下一个右侧节点指针 II.go
  方法二：使用已建立的 next 指针 ！！！！
  空间复杂度：O(1)，不需要存储额外的节点。

### 删除二叉搜索树中的节点

+ leetcode/tree/450. 删除二叉搜索树中的节点.go

root 为叶子节点，没有子树。此时可以直接将它删除，即返回空。
root 只有左子树，没有右子树。此时可以将它的左子树作为新的子树，返回它的左子节点。
root 只有右子树，没有左子树。此时可以将它的右子树作为新的子树，返回它的右子节点。

root 有左右子树，这时可以将 root 的后继节点（比 root 大的最小节点，即它的右子树中的最小节点，记为 successor作为新的根节点替代 root，并将 successor 从 root 的右子树中删除，使得在保持有序性的情况下合并左右子树。
简单证明，successor 位于 root 的右子树中，因此大于 root 的所有左子节点；successor 是 root 的右子树中的最小节点，因此小于 root 的右子树中的其他节点。以上两点保持了新子树的有序性。


### 出现次数最多的子树元素和
+ 给你一个二叉树的根结点root，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。

输入: root = [5,2,-3]
输出: [2,-3,4]  （2、-3、5+2+-3=4）
输入: root = [5,2,-5]
输出: [2] （5、2、5+2+-5=2）-> 2

+ leetcode/tree/508. 出现次数最多的子树元素和.go

```go
func findFrequentTreeSum(root *TreeNode) (ans []int) {
    cnt := map[int]int{}
    maxCnt := 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        sum := node.Val + dfs(node.Left) + dfs(node.Right)
        cnt[sum]++
        if cnt[sum] > maxCnt {
            maxCnt = cnt[sum]
        }
        return sum
    }
    dfs(root)

    for s, c := range cnt {
        if c == maxCnt {
            ans = append(ans, s)
        }
    }
    return
}

```

### 二叉树中所有距离为 K 的结点

+ leetcode/tree/863. 二叉树中所有距离为 K 的结点.go
  给定一个二叉树（具有根结点root），一个目标结点target，和一个整数值 k 。
  返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。

深度优先搜索 + 哈希表
若将 target 当作树的根结点，我们就能从 target 出发，使用深度优先搜索去寻找与 target 距离为 kk 的所有结点，即深度为 kk 的所有结点。

由于输入的二叉树没有记录父结点，为此，我们从根结点 root 出发，使用深度优先搜索遍历整棵树，同时用一个哈希表记录每个结点的父结点。

然后从 target 出发，使用深度优先搜索遍历整棵树，除了搜索左右儿子外，还可以顺着父结点向上搜索。

代码实现时，由于每个结点值都是唯一的，哈希表的键可以用结点值代替。此外，为避免在深度优先搜索时重复访问结点，递归时额外传入来源结点 from，在递归前比较目标结点是否与来源结点相同，不同的情况下才进行递归。
(左右子树是往下的，父节点是往上的，所以可能相互重复)


```go

func distanceK(root, target *TreeNode, k int) (ans []int) {
    // 从 root 出发 DFS，记录每个结点的父结点
    parents := map[int]*TreeNode{}
    var findParents func(*TreeNode)
    findParents = func(node *TreeNode) {
        if node.Left != nil {
            parents[node.Left.Val] = node
            findParents(node.Left)
        }
        if node.Right != nil {
            parents[node.Right.Val] = node
            findParents(node.Right)
        }
    }
    findParents(root)

    // 从 target 出发 DFS，寻找所有深度为 k 的结点
    var findAns func(*TreeNode, *TreeNode, int)
    findAns = func(node, from *TreeNode, depth int) {
        if node == nil {
            return
        }
        if depth == k { // 将所有深度为 k 的结点的值计入结果
            ans = append(ans, node.Val)
            return
        }
		//上一个是往上遍历的，且当时的结点是左结点，别回来
        if node.Left != from {
            findAns(node.Left, node, depth+1)
        }
        //上一个是往上遍历的，且当时的结点是右结点，别回来
        if node.Right != from {
            findAns(node.Right, node, depth+1)
        }
		//上一个是往下遍历的，不要再往上了，别回来
        if parents[node.Val] != from {
            findAns(parents[node.Val], node, depth+1)
        }
    }
    findAns(target, nil, 0)
    return
}

```

### 判断是不是平衡二叉树
+ 定义一个函数，返回树的高度，如果不是平衡二叉树，则返回-1 ！！
```Java
class Solution {
    public boolean isBalanced(TreeNode root) {
        return recur(root) != -1;
    }

    private int recur(TreeNode root) {
        if (root == null) return 0;
        int left = recur(root.left);
        if(left == -1) return -1;
        int right = recur(root.right);
        if(right == -1) return -1;
        return Math.abs(left - right) < 2 ? Math.max(left, right) + 1 : -1;
    }
}
```

----

# 动态规划

## 方法技巧
+ 凑零钱问题
+ 带“备忘录”的递归
+ 自顶向下、自底向上

+ 状态转移方程

+ 动态规划的特性，穷举加备忘录/DP table 优化
+ 只要遇到求最值的算法问题，首先要思考的就是：如何穷举出所有可能的结果？
+ 穷举主要有两种算法，就是回溯算法和动态规划，前者就是暴力穷举，而后者是根据状态转移方程推导「状态」

+ 第一步要明确两点，「状态」和「选择」。
+ 第二步要明确 dp 数组的定义。

+ 二维动态规划压缩成一维动态规划吗？这就是状态压缩

## 题目

+ 总结公式：
+ 兑换零钱（一）
+ 打家劫舍（一）
  此转移方程为dp[i]=max(dp[i−1],nums[i−1]+dp[i−2])
+ 打家劫舍（二）
  这一问在第一问的基础上添加了房屋首尾相连的条件，所以首尾两个数字之间我们最多只能选取一个。

∙ \bullet∙ 既然这样我们就可以将整个数组进行分割，分为 [0,n-2] 与 [1,n-1] 两个部分，分别求解它们的最大值然后再选择两个之间的较大值作为最终结果，其余情况与第一问相同。
————————————————
版权声明：本文为CSDN博主「桃陉」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_46308081/article/details/119087704


+ leetcode/dp/70. 爬楼梯.go
  climbStairs(n) = climbStairs(n-2) + climbStairs(n-1)

+ leetcode/dp/119. 杨辉三角 II.go
  算好规律

+ leetcode/dp/5. 最长回文子串.go  !!!!


+ leetcode/dp/64. 最小路径和.go
  dp[i,j]=min(dp[i-1,j],dp[i,j-1]) + val(i,j)


+ leetcode/dp/300. 最长递增子序列.go !!!!

+ leetcode/dp/55. 跳跃游戏.go !!!
  贪心算法！！ 遍历更新最大可达位置！！！


+ leetcode/dp/45. 跳跃游戏 II.go
  遍历计算到达每一步的最小步数！
  需考虑数组长度为1 的情况 ！！！！

+ leetcode/dp/139. 单词拆分.go !!!!

```go
func wordBreak(s string, wordDict []string) bool {
    wordDictSet := make(map[string]bool)
    for _, w := range wordDict {
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordDictSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}
```

+ leetcode/dp/221. 最大正方形.go ！！！


----

# 链表
+ 链表的问题基本都挺容易想，但是写起来比价麻烦，需要许多中间变量，所以做之前要先画图整理清楚，否则写的时候很容易乱

+ 相交链表
    - 双指针，为空后分别重新赋值指向对方，这样长度相同，遇到相等则相交，都为空则不相交

+ 闭合为环的技巧，知道长度可以取模，(n - 1) - (k mod n) ，转化成逆序遍历长度
+ exercise/leetcode/tencent/61. 旋转链表_test.go

## 反转单链表

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


----

# 字符串
## 返回一个字符串的最长不重复子串的长度
+ "滑动窗口"（左右两个指针），hashset判断重复元素，遍历一遍即可。左指针向右移动时删除hashset中的元素，左指针向右移动时增加hashset中的元素。
+ algorithm-gooooooo/exercise/leetcode/hot100/3. 最长不含重复字符的子字符串_test.go
+ algorithm-gooooooo/leetcode/剑指 Offer 48. 最长不含重复字符的子字符串.go （以这个为准！）

# 其他

## 顺时针打印矩阵
定义上下左右下标标量，分四个方向遍历，直到左大于右 或 上大于下

## 两数之和
+ hash表的使用

## 统计一个文档中出现频率最多的k个单词
+ <http://www.lxway.com/4084220604.htm>
1. 1.统计文档中所有不同word出现的频率,hash
2. 求topk(堆排序)（nlogn）

## 判断一个无序数列是否是等差数列
+ 如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。
1. 对数组进行第一次遍历，找出数组中的min.max
2. 如果是等差数列，那么公差必然是（max-min）/(n-1)   n为元素个数
3. 有了公差，有了首项，有了尾项。这个等差数列实际上就模拟出来了。

```java
// Utils4Java/kingson/src/main/java/com/kxw/interview/ArithmeticProgressionJudge.java
public class ArithmeticProgressionJudge {

    //如果公差为0，那么说明数列中所有的数都是相同的 ，判断所有数是否相同
    //公差不为0

    public static void main(String[] args) {

        int[] seq = {2, 4, 8, 6};
        System.out.println(judge(seq));

    }

    public static boolean judge(int[] arr) {

        int min = arr[0], max = arr[0];
        for (int i = 1; i < arr.length; i++) {//对数组进行第一次遍历,找出数组中min ,max

            if (arr[i] < min) {
                min = arr[i];
            } else if (arr[i] > max) {
                max = arr[i];
            }
        }

        int dif = (max - min) / (arr.length - 1);

        if (dif == 0) {//如果公差为0，那么说明数列中所有的数都是相同的 ，判断所有数是否相同
            for (int i = 0; i < arr.length; i++) {
                if (arr[i] != min) {
                    return false;
                }
            }
        } else {
            int[] positionArr = new int[arr.length];
            for (int i = 0; i < arr.length; i++) {
                int position = (arr[i] - min) / dif;
                if (position >= arr.length || positionArr[position] == 1) {
                    return false;
                } else {
                    positionArr[position] = 1;
                }
            }
            return true;

        }
        return true;
    }

}

```