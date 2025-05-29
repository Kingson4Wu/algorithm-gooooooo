+ 字符串根据下标获取子串

```go
s[maxStart:maxLen]
```

+ ascii
```go
 c >= '0' && c <= '9' 
```

+ int 最大值最小值
```go
math.MaxInt32 
math.MinInt32
```
```go
	if res > math.MaxInt32 || res <= math.MinInt32 {
			return 0
		}
```

+ 排序 ： `sort.Ints`
+ 逆序：
```go
sort.Slice(nums, func(i, j int) bool {
    return i < j
})
```

+ Go 语言中，可以通过实现 sort.Interface 接口来自定义排序。该接口包括三个方法：

Len() 返回待排序元素的个数。
Less(i, j int) bool 比较索引为 i 和 j 的元素大小，如果第 i 个元素应该排在第 j 个元素之前，则返回 true，否则返回 false。
Swap(i, j int) 交换索引为 i 和 j 的元素。
下面是一个示例程序，展示如何使用 sort.Interface 接口进行自定义排序：

```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    people := []Person{
        {"Bob", 31},
        {"John", 20},
        {"Jane", 42},
        {"Mike", 19},
    }

    fmt.Println("Before sorting:", people)

    sort.Sort(ByAge(people))

    fmt.Println("After sorting:", people)
}

```


+ 对某一字符串排序

```go
str := "eat"
b := []byte(str)
sort.Slice(b, func(i, j int) bool {
	return b[i] < b[j]
})
str = string(b)

```

或

```go
func sortString(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

```