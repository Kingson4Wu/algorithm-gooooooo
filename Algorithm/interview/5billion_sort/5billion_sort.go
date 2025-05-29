package _billion_sort

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// 第一步：分（分块 + 内部排序 + 写临时文件）
const (
	//memBufferSize  = 10_000_000 // 约占内存 100MB（每个 int64 是 8B）
	memBufferSize  = 1000
	inputFile      = "bigdata"
	tempFilePrefix = "bigdata.part."
)

func splitAndSortChunks() ([]string, error) {
	var tempFiles []string
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buffer := make([]int, 0, memBufferSize)
	chunkIdx := 0

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		buffer = append(buffer, num)

		if len(buffer) >= memBufferSize {
			tempName, err := flushSortedChunk(buffer, chunkIdx)
			if err != nil {
				return nil, err
			}
			tempFiles = append(tempFiles, tempName)
			buffer = buffer[:0]
			chunkIdx++
		}
	}

	// 最后一块
	if len(buffer) > 0 {
		tempName, err := flushSortedChunk(buffer, chunkIdx)
		if err != nil {
			return nil, err
		}
		tempFiles = append(tempFiles, tempName)
	}

	return tempFiles, nil
}

func flushSortedChunk(data []int, idx int) (string, error) {
	sort.Ints(data)
	filename := fmt.Sprintf("%s%d.sorted", tempFilePrefix, idx)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range data {
		_, err := fmt.Fprintf(writer, "%d\n", num)
		if err != nil {
			return "", err
		}
	}
	writer.Flush()
	return filename, nil
}

// 第二步：合（k 路归并）
func mergeSortedFiles(inputFiles []string, outputFile string) error {
	const bufferSize = 1000 // 缓冲区写入数量

	// 打开文件和 Scanner
	files := make([]*os.File, len(inputFiles))
	scanners := make([]*bufio.Scanner, len(inputFiles))
	currentValues := make([]*int, len(inputFiles)) // 每个文件当前行的值

	for i, fname := range inputFiles {
		f, err := os.Open(fname)
		if err != nil {
			return fmt.Errorf("打开文件失败 %s: %v", fname, err)
		}
		files[i] = f
		scanners[i] = bufio.NewScanner(f)
		if scanners[i].Scan() {
			v, _ := strconv.Atoi(scanners[i].Text())
			currentValues[i] = &v
		} else {
			currentValues[i] = nil
		}
	}

	// 打开输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	buffer := make([]int, 0, bufferSize)

	for {
		minIndex := -1
		var minVal int

		// 找最小值
		for i, val := range currentValues {
			if val == nil {
				continue
			}
			if minIndex == -1 || *val < minVal {
				minIndex = i
				minVal = *val
			}
		}

		// 所有都空了，退出循环
		if minIndex == -1 {
			break
		}

		// 添加到缓冲区
		buffer = append(buffer, minVal)
		if len(buffer) >= bufferSize {
			for _, v := range buffer {
				fmt.Fprintln(writer, v)
			}
			buffer = buffer[:0]
		}

		// 推进读取该文件下一行
		if scanners[minIndex].Scan() {
			v, _ := strconv.Atoi(scanners[minIndex].Text())
			currentValues[minIndex] = &v
		} else {
			currentValues[minIndex] = nil // 文件结束
		}
	}

	// 写剩余数据
	for _, v := range buffer {
		fmt.Fprintln(writer, v)
	}

	// 关闭文件
	for _, f := range files {
		f.Close()
	}
	return nil
}
