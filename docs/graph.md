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
