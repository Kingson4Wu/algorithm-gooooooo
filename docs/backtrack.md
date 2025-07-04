整理一个尽量统一、相似的**排列、组合、子集**模板，涵盖**有重复和无重复**的情况，方便你刷题和记忆。

---

# 🗂 通用回溯模板总结

| 题型          | 递归参数 | 关键点               | 重复处理策略                 | 代码模板示例（Go 伪码简化）   |
| ----------- | ---- | ----------------- | ---------------------- | ----------------- |
| **排列**      | 无需起点 | 需要标记已用元素 `used[]` | 排序 + `used` + 跳过相邻重复元素 | 见排列 II 模板         |
| **组合 / 子集** | 需要起点 | 控制遍历起点，防止重复使用前面元素 | 排序 + 跳过同层相邻重复元素        | 见组合 II / 子集 II 模板 |

---

## 1. 排列（Permutation）

### 1.1 无重复元素 — 基础排列（46）

```go
func permute(nums []int) [][]int {
    var res [][]int
    var path []int
    used := make([]bool, len(nums))

    var dfs func()
    dfs = func() {
        if len(path) == len(nums) {
            res = append(res, append([]int(nil), path...))
            return
        }
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            used[i] = true
            path = append(path, nums[i])
            dfs()
            path = path[:len(path)-1]
            used[i] = false
        }
    }
    dfs()
    return res
}
```

### 1.2 有重复元素 — 排列 II（47）

```go
func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    var path []int
    used := make([]bool, len(nums))

    var dfs func()
    dfs = func() {
        if len(path) == len(nums) {
            res = append(res, append([]int(nil), path...))
            return
        }
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				//只能先用同一组重复数字的“第一个”，不能先用后面的。
				//如果现在选择了后一个重复元素，就会导致重复排列。
                continue
            }
            used[i] = true
            path = append(path, nums[i])
            dfs()
            path = path[:len(path)-1]
            used[i] = false
        }
    }
    dfs()
    return res
}
```

---

## 2. 组合 / 子集（Combination / Subset）

> 组合和子集本质相似，区别通常在题目描述和是否需要用所有元素。

### 2.1 无重复元素 — 组合 / 子集 I（如 78）

```go
func subsets(nums []int) [][]int {
    var res [][]int
    var path []int

    var dfs func(start int)
    dfs = func(start int) {
        res = append(res, append([]int(nil), path...))
        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            dfs(i + 1)
            path = path[:len(path)-1]
        }
    }
    dfs(0)
    return res
}
```

### 2.2 有重复元素 — 组合 / 子集 II（如 90）

```go
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    var path []int

    var dfs func(start int)
    dfs = func(start int) {
        res = append(res, append([]int(nil), path...))
        for i := start; i < len(nums); i++ {
            if i > start && nums[i] == nums[i-1] {
                continue
            }
            path = append(path, nums[i])
            dfs(i + 1)
            path = path[:len(path)-1]
        }
    }
    dfs(0)
    return res
}
```

---

# 🔑 统一模板结构

```go
func genericBacktrack(nums []int, allowRepeat bool, isPermutation bool) [][]int {
    if allowRepeat {
        sort.Ints(nums)
    }
    var res [][]int
    var path []int
    used := make([]bool, len(nums)) // 仅排列需要

    var dfs func(start int)
    if isPermutation {
        dfs = func(start int) {
            if len(path) == len(nums) {
                res = append(res, append([]int(nil), path...))
                return
            }
            for i := 0; i < len(nums); i++ {
                if used[i] {
                    continue
                }
                if allowRepeat && i > 0 && nums[i] == nums[i-1] && !used[i-1] {
                    continue
                }
                used[i] = true
                path = append(path, nums[i])
                dfs(0) // start 参数无意义，保留接口一致
                path = path[:len(path)-1]
                used[i] = false
            }
        }
    } else {
        dfs = func(start int) {
            res = append(res, append([]int(nil), path...))
            for i := start; i < len(nums); i++ {
                if allowRepeat && i > start && nums[i] == nums[i-1] {
                    continue
                }
                path = append(path, nums[i])
                dfs(i + 1)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0)
    return res
}
```

---

# 总结要点

| 特征         | 排列（Permutation）       | 组合 / 子集（Combination / Subset） |
| ---------- | --------------------- | ----------------------------- |
| 是否用 `used` | 是                     | 否（一般）                         |
| 是否排序       | 重复元素时必须排序             | 重复元素时必须排序                     |
| 是否有起点参数    | 可无（used法），有时用 `cur`   | 必须有起点参数 `cur`                 |
| 去重判断       | 有重复元素时判断 `!used[i-1]` | 有重复元素时跳过同层相同元素                |
| 递归参数       | 可无或固定（`dfs()`）        | 必须带起点（`dfs(cur int)`）         |

---

