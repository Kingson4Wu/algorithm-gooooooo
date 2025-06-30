# 无向图
+ 由两个部分组成：
顶点（Vertices）：图中的节点。
边（Edges）：连接两个顶点的线段。

+ 边用集合表示：一条边连接两个顶点，用 {A, B} 表示；而不是有向图中的 (A, B)。
  度（Degree）：一个顶点的度是连接它的边的数量（不区分方向）。

+ 无向图可以表示为：
  顶点：{A, B, C}
  边：{{A, B}, {B, C}}
+ 图形大致如下：
  A —— B —— C

+ 无向图的深度优先搜索
  起始于某个顶点；
  标记为“已访问”；
  遍历它的邻居；
  对于每一个未访问过的邻居，递归执行 DFS；
  遇到死胡同（没有未访问的邻居）就回退。

+ 递归实现 DFS：
```python
def dfs(graph, start, visited=None):
    if visited is None:
        visited = set()

    print(start)  # 访问当前节点
    visited.add(start)

    for neighbor in graph[start]:
        if neighbor not in visited:
            dfs(graph, neighbor, visited)

# 调用
dfs(graph, 'A')

```
+ 非递归实现（使用栈）
```python
def dfs_iterative(graph, start):
    visited = set()
    stack = [start]

    while stack:
        node = stack.pop()
        if node not in visited:
            print(node)
            visited.add(node)
            # 为了保持顺序，反转邻居顺序
            for neighbor in reversed(graph[node]):
                if neighbor not in visited:
                    stack.append(neighbor)

dfs_iterative(graph, 'A')
```

+ 无向图 DFS 的注意事项

防止死循环：一定要用 visited 集合记录已访问节点，因为无向图的边是双向的，否则可能在 A-B-A-B 之间反复访问。
图不连通时，只 DFS 一个起点可能无法遍历整个图；可以对所有节点进行一次 DFS

```python
def dfs_all(graph):
    visited = set()
    for node in graph:
        if node not in visited:
            dfs(graph, node, visited)
```

# 有向图

## 有向图的拓扑排序

+ 有向图的拓扑排序（Topological Sorting of a Directed Graph）是一种将**有向无环图（DAG，Directed Acyclic Graph）**中的所有顶点排成线性序列的方法，使得对于每一条有向边 u → v，顶点 u 都排在 v 的前面。

+ 拓扑排序就是给有依赖关系的事情排一个顺序。例如：

学习顺序：先学 A，再学 B，最后学 C。
项目任务：某个任务必须在另一个任务完成后才能进行。
如果我们把每个任务看作图的一个节点，每个“先做 A 才能做 B”的关系看作一条从 A 指向 B 的有向边，整个问题就变成了对一个 DAG 做拓扑排序。

+ 适用范围
  必须是有向无环图（DAG）。
  如果图中存在环，则无法拓扑排序。

+ 拓扑排序的常用算法：
  ✅ 方法一：入度表 + 队列（Kahn 算法）

统计每个顶点的入度（有多少边指向它）。
将所有入度为 0 的点加入队列。
从队列中取出一个顶点 u，加入结果序列。
删除所有从 u 出发的边（即相邻顶点 v 的入度减 1）。
如果 v 的入度变为 0，把它加入队列。
重复直到队列为空。
如果结果序列中包含了所有节点，说明拓扑排序成功；否则图中有环。

✅ 方法二：DFS（深度优先搜索）法

从一个未访问的节点开始 DFS。
每访问一个节点，递归访问它所有的后继节点。
后继节点都访问完后，把当前节点加入栈。
所有节点访问完成后，从栈顶到栈底就是拓扑序列。

应用场景
编译器：代码编译顺序（模块依赖）
项目管理：任务调度（任务依赖）
数据管道：处理步骤排序
课程表安排（如 Leetcode 207, 210）

```go

package main

import (
  "fmt"
)

// 拓扑排序（Kahn 算法）
func topologicalSort(graph map[string][]string) ([]string, bool) {
  inDegree := make(map[string]int)
  var result []string

  // 初始化入度表
  for u := range graph {
    if _, ok := inDegree[u]; !ok {
      inDegree[u] = 0
    }
    for _, v := range graph[u] {
      inDegree[v]++
    }
  }

  // 入度为0的节点入队
  var queue []string
  for node, deg := range inDegree {
    if deg == 0 {
      queue = append(queue, node)
    }
  }

  // 拓扑排序
  for len(queue) > 0 {
    // 出队一个元素
    node := queue[0]
    queue = queue[1:]
    result = append(result, node)

    // 遍历它的邻接点
    for _, neighbor := range graph[node] {
      inDegree[neighbor]--
      if inDegree[neighbor] == 0 {
        queue = append(queue, neighbor)
      }
    }
  }

  // 判断是否是 DAG（拓扑排序成功）
  if len(result) != len(inDegree) {
    return nil, false // 有环
  }
  return result, true
}

func main() {
  // 示例图（没有B→C）
  graph := map[string][]string{
    "A": {"B", "C"},
    "B": {"D"},
    "C": {"D"},
    "D": {},
  }

  order, ok := topologicalSort(graph)
  if !ok {
    fmt.Println("图中存在环，无法拓扑排序")
  } else {
    fmt.Println("拓扑排序结果：", order)
  }
}

```

+ Kahn 算法之所以可以完成拓扑排序，关键在于它 逐步移除入度为 0 的节点，并确保所有依赖顺序都被满足。
+ Kahn 算法的关键逻辑：
  每次只处理 入度为 0 的节点，说明这个节点没有任何前驱（没有依赖）。
  它可以排在最前面，因为没人依赖它，它也不依赖别人。
  处理完一个节点（加入结果），就“删除”它对别人的依赖（让它指向的节点入度减 1）。
  只有在所有前驱都处理完之后，一个节点才能变为入度 0 → 再被处理。
  ➡️ 这个过程正好保证了拓扑排序中“先处理前驱，再处理后继”的要求。
+ 为什么只适用于有向无环图（DAG）？

如果图中有 环，那么某个节点永远不会变成入度为 0（因为它始终有一个前驱在环里）。
这样队列就会耗尽，剩下的节点都还有入度 ≠ 0，导致结果节点数量 < 图中节点数。
✅ 因此：可以用这个算法顺便检测是否有环（如果排序结果节点数 < 图节点数，说明有环）

+ Kahn 算法之所以成立，是因为它始终只处理“没有依赖的节点”，每次处理都维护了拓扑顺序的合法性。只要图是 DAG，就一定能处理完所有节点，得到一个正确的拓扑排序。
