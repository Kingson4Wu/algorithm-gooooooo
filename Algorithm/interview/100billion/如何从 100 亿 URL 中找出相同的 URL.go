package algorithm

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func seek() {

	hashToFile("a.txt", 10)
	hashToFile("b.txt", 10)

	result := compare("a.txt", "b.txt", 10)

	for _, url := range result {
		fmt.Println(url)
	}

}

func compare(filenameA, filenameB string, fileNum int) []string {

	var result []string
	for i := 0; i < fileNum; i++ {
		hashA := loadFileToHash(hashFileName(filenameA, i))
		hashB := loadFileToHash(hashFileName(filenameB, i))

		for url := range hashA {
			if ok, _ := hashB[url]; ok {
				result = append(result, url)
			}
		}
	}
	return result
}

func loadFileToHash(filename string) map[string]bool {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hash := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		hash[line] = true
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return hash
}

func hashToFile(filename string, fileNum int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	files := createFile(filename, fileNum)

	defer func() {
		for _, f := range files {
			if err := f.Flush(); err != nil {
				panic(err)
			}
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			hash := sha256.Sum256([]byte(line))
			index := int(hash[0]) % 10

			writer := files[index]
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func hashFileName(filename string, i int) string {
	return "output/" + filename + "_" + strconv.Itoa(i)
}

func createFile(filename string, fileNum int) map[int]*bufio.Writer {

	files := make(map[int]*bufio.Writer, 10)
	for i := 0; i < fileNum; i++ {
		name := hashFileName(filename, i)
		if _, err := os.Stat(name); err == nil {
			if err := os.Remove(name); err != nil {
				panic(err)
			}
			//fmt.Printf("File %s has been removed.\n", name)
		} else if os.IsNotExist(err) {
			//fmt.Printf("File %s does not exist.\n", name)
		} else {
			panic(err)
		}

		file, err := os.Create(name)
		if err != nil {
			panic(err)
		}
		//defer file.Close()

		writer := bufio.NewWriter(file)

		files[i] = writer
	}
	return files
}

/**
https://mp.weixin.qq.com/s/hDmObRRRYWa3U0R9ucdd9w

题目描述

给定 a、b 两个文件，各存放 50 亿个 URL，每个 URL 各占 64B，内存限制是 4G。请找出 a、b 两个文件共同的 URL。

解答思路

每个 URL 占 64B，那么 50 亿个 URL占用的空间大小约为 320GB。
“5, 000, 000, 000 * 64B ≈ 5GB * 64 = 320GB
由于内存大小只有 4G，因此，我们不可能一次性把所有 URL 加载到内存中处理。对于这种类型的题目，一般采用分治策略 ，即：把一个文件中的 URL 按照某个特征划分为多个小文件，使得每个小文件大小不超过 4G，这样就可以把这个小文件读到内存中进行处理了。
思路如下 ：
首先遍历文件 a，对遍历到的 URL 求 hash(URL) % 1000 ，根据计算结果把遍历到的 URL 存储到 a0, a1, a2, ..., a999，这样每个大小约为 300MB。使用同样的方法遍历文件 b，把文件 b 中的 URL 分别存储到文件 b0, b1, b2, ..., b999 中。这样处理过后，所有可能相同的 URL 都在对应的小文件中，即 a0 对应 b0, ..., a999 对应 b999，不对应的小文件不可能有相同的 URL。那么接下来，我们只需要求出这 1000 对小文件中相同的 URL 就好了。
接着遍历 ai( i∈[0,999] )，把 URL 存储到一个 HashSet 集合中。然后遍历 bi 中每个 URL，看在 HashSet 集合中是否存在，若存在，说明这就是共同的 URL，可以把这个 URL 保存到一个单独的文件中。

方法总结

1.分而治之，进行哈希取余；
2.对每个子文件进行 HashSet 统计。

基本思路是这样，如果分完还有很大的文件，那就文件再继续分，一直递归到文件符合设置的最大大小即可

*/
